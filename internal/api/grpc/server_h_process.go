package grpc

import (
	"context"
	"errors"
	"fmt"
	"github.com/Alp4ka/classifier-aaS/internal/contextcomponent"
	contextrepository "github.com/Alp4ka/classifier-aaS/internal/contextcomponent/repository"
	"github.com/Alp4ka/classifier-aaS/internal/processor"
	api "github.com/Alp4ka/classifier-api"
	"github.com/Alp4ka/mlogger"
	"github.com/Alp4ka/mlogger/field"
	"github.com/google/uuid"
	"github.com/guregu/null/v5"
	"github.com/hashicorp/go-set/v2"
	"golang.org/x/sync/errgroup"
	"io"
	"sync"
)

func (s *Server) Process2(src api.GWManagerService_ProcessServer) (err error) {
	ctx := src.Context()
	defer func() {
		if err != nil {
			mlogger.L(ctx).Error("Process failed", field.Error(err))
		}
	}()
	deferableF := func() {}
	defer deferableF()

	var (
		once    sync.Once
		session *contextcomponent.Session
		proc    *processor.Processor
		reqids  = set.New[string](1)
	)

	for {
		// TODO: Metrics.

		// Reading request.
		req, err := src.Recv()
		if err == io.EOF {
			err = s.contextService.ReleaseSession(ctx, session.Model.ID, contextrepository.SessionStateClosedAgent)
			if err != nil {
				return fmt.Errorf("failed to release session as agent: %w", err)
			}
			mlogger.L(ctx).Info("Client closed connection")
			return nil
		} else if err != nil {
			return fmt.Errorf("failed to receive client request: %w", err)
		}
		if reqids.Contains(req.GetRequestId()) {
			return errors.New("duplicate request id")
		}
		reqids.Insert(req.GetRequestId())

		// Configure ctx.
		ctx = field.WithContextFields(ctx, field.String("session_id", req.GetSessionId()), field.String("req_id", req.GetRequestId()), field.String("req", req.GetRequestData()))

		// Session.
		sessID, err := uuid.Parse(req.GetSessionId())
		if err != nil {
			return fmt.Errorf("unable to parse session uuid: %w", err)
		}
		once.Do(func() {
			session, err = s.contextService.GetSession(ctx, &contextcomponent.GetSessionParams{
				SessionID: uuid.NullUUID{UUID: sessID, Valid: true},
				Active:    null.BoolFrom(true),
			})
			if err != nil {
				return
			}
			proc, err = processor.NewProcessor(session.Tree)
			if err != nil {
				return
			}
			respChan, err := proc.Start(ctx)
			if err != nil {
				return
			}
			deferableF = proc.Close

			go func() {
				for {
					resp, ok := <-respChan
					if !ok {
						return
					}

					// Sending response.
					ret := &api.ProcessResponse{
						ResponseData:     &resp.Data,
						End:              resp.End,
						RequestRequired:  resp.RequestRequired,
						ResponseRequired: resp.ResponseRequired,
					}
					err = src.Send(ret)
					if err != nil {
						mlogger.L(ctx).Error("Send response failed", field.Error(err))
					}
					if ret.End {
						err = s.contextService.ReleaseSession(ctx, session.Model.ID, contextrepository.SessionStateFinished)
						if err != nil {
							mlogger.L(ctx).Error("Failed to release session as finish", field.Error(err))
							return
						}
					}
				}
			}()

			mlogger.L(ctx).Info("Session connected")
		})
		if err != nil {
			return fmt.Errorf("failed to get session: %w", err)
		}
		if session.Model.ID != sessID {
			return fmt.Errorf("session id mismatch")
		}
		if !session.Expired() {
			err = s.contextService.ReleaseSession(ctx, session.Model.ID, contextrepository.SessionStateClosedRotten)
			return errors.Join(fmt.Errorf("session expired"), err)
		}

		// Handle.
		mlogger.L(ctx).Info("Handle client request")
		proc.Handle(&processor.Request{Data: req.GetRequestData()})
	}
}

func (s *Server) Process(src api.GWManagerService_ProcessServer) (err error) {
	ctx := src.Context()
	defer func() {
		if err != nil {
			mlogger.L(ctx).Error("Process failed", field.Error(err))
		}
	}()

	// Read initial request.
	req, err := src.Recv()
	if err == io.EOF {
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to receive initial client request: %w", err)
	}

	// Configure ctx.
	ctx = field.WithContextFields(ctx,
		field.String("session_id", req.GetSessionId()),
		field.String("req_id", req.GetRequestId()),
	)
	s.StoreRequest(req.GetRequestId())

	// Init session.
	sessID, err := uuid.Parse(req.GetSessionId())
	if err != nil {
		return fmt.Errorf("session id wrong format: %w", err)
	}
	session, err := s.contextService.GetSession(ctx, &contextcomponent.GetSessionParams{
		SessionID: uuid.NullUUID{UUID: sessID, Valid: true},
		Active:    null.BoolFrom(true),
	})
	proc, err := processor.NewProcessor(session.Tree)
	if err != nil {
		return fmt.Errorf("failed to create processor: %w", err)
	}
	defer proc.Close()
	retChan, err := proc.Start(ctx)
	if err != nil {
		return fmt.Errorf("failed to start processor: %w", err)
	}

	errg, errctx := errgroup.WithContext(ctx)
	errg.Go(func() error {
		return s.handleClientRequest(errctx, session, proc, src)
	})
	errg.Go(func() error {
		return s.handleProcessorResponse(errctx, session, retChan, src)
	})
	err = errg.Wait()
	if errors.Is(err, io.EOF) {
		mlogger.L(ctx).Info("Finished")
		return nil
	}

	return err
}

func (s *Server) handleClientRequest(
	ctx context.Context,
	sess *contextcomponent.Session,
	proc *processor.Processor,
	src api.GWManagerService_ProcessServer,
) error {
	for {
		// Reading request.
		req, err := src.Recv()
		if err != nil {
			return err
		}

		// Deduplication.
		if s.IsDuplicate(req.GetRequestId()) {
			continue // TODO: mb not continue or return error.
		}
		s.StoreRequest(req.GetRequestId())

		// Session. // TODO: rwmu since we can release session both ways.
		if !sess.Expired() && sess.Operable() {
			return fmt.Errorf("session expired")
		}

		// Handle.
		mlogger.L(ctx).Info("Handle client request")
		proc.Handle(&processor.Request{Data: req.GetRequestData()})
	}
}

func (s *Server) handleProcessorResponse(
	ctx context.Context,
	sess *contextcomponent.Session,
	ch <-chan *processor.Response,
	src api.GWManagerService_ProcessServer,
) error {
	for {
		ret, ok := <-ch
		if !ok {
			return nil
		}

		// Sending response.
		err := src.Send(
			&api.ProcessResponse{
				ResponseData:     &ret.Data,
				End:              ret.End,
				RequestRequired:  ret.RequestRequired,
				ResponseRequired: ret.ResponseRequired,
			},
		)
		if err != nil {
			return err
		}
		if ret.End {
			return io.EOF
		}
	}
}
