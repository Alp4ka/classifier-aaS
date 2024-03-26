package grpc

import (
	"errors"
	"fmt"
	"github.com/Alp4ka/classifier-aaS/internal/contextcomponent"
	contextrepository "github.com/Alp4ka/classifier-aaS/internal/contextcomponent/repository"
	"github.com/Alp4ka/classifier-aaS/internal/processor"
	"github.com/Alp4ka/classifier-aaS/pkg/api"
	"github.com/Alp4ka/mlogger"
	"github.com/Alp4ka/mlogger/field"
	"github.com/google/uuid"
	"github.com/guregu/null/v5"
	"io"
	"sync"
)

func (s *Server) Process(src api.GWManagerService_ProcessServer) (err error) {
	ctx := src.Context()
	defer func() {
		if err != nil {
			mlogger.L(ctx).Error("Process failed", field.Error(err))
		}
	}()

	var (
		once    sync.Once
		session *contextcomponent.Session
		proc    *processor.Processor
	)
	for {
		// TODO: Metrics.

		// Reading request.
		req, err := src.Recv()
		if err == io.EOF {
			mlogger.L(ctx).Info("Client closed connection")
			break
		} else if err != nil {
			return fmt.Errorf("failed to receive client request: %w", err)
		}

		// Configure ctx.
		ctx = field.WithContextFields(ctx, field.String("session_id", req.GetSessionId()), field.String("req", req.GetRequestData()))

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
			proc = processor.NewProcessor(session.Tree)
			mlogger.L(ctx).Info("Session connected")
		})
		if err != nil {
			return fmt.Errorf("failed to get session: %w", err)
		}
		if session.Model.ID != sessID {
			return fmt.Errorf("session id mismatch")
		}
		if !session.Active() {
			err = s.contextService.ReleaseSession(ctx, session.Model.ID, contextrepository.SessionStateClosedRotten)
			return errors.Join(fmt.Errorf("session expired"), err)
		}

		// Handle.
		mlogger.L(ctx).Info("Handle client request")
		resp, err := proc.Handle(ctx, &processor.Request{Data: req.GetRequestData()})
		if err != nil {
			return fmt.Errorf("failed to handle: %w", err)
		}

		// Sending response.
		ret := &api.ProcessResponse{
			ResponseData: &resp.Data,
			End:          resp.End,
		}
		err = src.Send(ret)
		if err != nil {
			return fmt.Errorf("failed to send message: %w", err)
		}
	}

	return nil
}
