package grpc

import (
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/Alp4ka/classifier-aaS/internal/telemetry"
	timepkg "github.com/Alp4ka/classifier-aaS/pkg/time"

	contextcomponent "github.com/Alp4ka/classifier-aaS/internal/components/context"
	processorcomponent "github.com/Alp4ka/classifier-aaS/internal/components/processor"
	api "github.com/Alp4ka/classifier-api"
	"github.com/Alp4ka/mlogger"
	"github.com/Alp4ka/mlogger/field"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
	"github.com/guregu/null/v5"
)

type ReqStorage struct {
	set mapset.Set[string]
}

func NewReqStorage() *ReqStorage {
	return &ReqStorage{set: mapset.NewThreadUnsafeSet[string]()}
}

func (r *ReqStorage) StoreRequest(req *api.ProcessRequest) error {
	if r.set.Contains(req.GetRequestId()) {
		return fmt.Errorf("duplicate request")
	}
	r.set.Add(req.GetRequestId())

	return nil
}

type environment struct {
	Session    *contextcomponent.Session
	Processor  *processorcomponent.Processor
	ReqStorage *ReqStorage
}

func (s *Server) Process(src api.GWManagerService_ProcessServer) (err error) {
	const fn = "Server.Process"

	ctx := src.Context()

	var (
		once                                = sync.Once{}
		awaitUserAction                     = true
		req             *api.ProcessRequest = nil
		env             *environment        = nil
	)

	// Metrics.
	{
		timeStart := timepkg.Now()
		defer func() {
			if env != nil && env.Session != nil {
				telemetry.T().ObserveProcessDuration(env.Session.Model.Gateway, timepkg.Now().Sub(timeStart))
			}
		}()
	}
	for {
		if awaitUserAction {
			// Reading request.
			req, err = src.Recv()
			if err == io.EOF {
				return nil
			} else if err != nil {
				return fmt.Errorf("failed to receive client request: %w", err)
			}

			// Prepare environment.
			once.Do(
				func() {
					var sessionID uuid.UUID
					sessionID, err = getSessionID(ctx)
					if err != nil {
						err = fmt.Errorf("%s: failed to get session id: %w", fn, err)
					}

					env, err = s.prepareEnvironment(ctx, sessionID)
					if err != nil {
						err = fmt.Errorf("%s: failed to prepare environment: %w", fn, err)
					}
				},
			)
			if err != nil {
				return err
			}

			err = env.ReqStorage.StoreRequest(req)
			if err != nil {
				mlogger.L(ctx).Warn(
					"Request duplication!",
					field.String("requestID", req.GetRequestId()),
					field.Bool("skip", true),
				)
				continue
			}

			// Processing.
			if !env.Session.Operable() {
				return fmt.Errorf("session expired")
			}
		}

		var resp *processorcomponent.Response
		resp, err = env.Processor.Process(ctx, &processorcomponent.Request{UserInput: req.GetRequestData()})
		if err != nil {
			return fmt.Errorf("%s: failed to process request: %w", fn, err)
		}

		var protoAction api.Action
		protoAction, err = actionToProtoAction(resp.Action)
		if err != nil {
			return fmt.Errorf("%s: failed to convert action to proto: %w", fn, err)
		}
		err = src.Send(
			&api.ProcessResponse{
				ResponseData: resp.UserOutput,
				Action:       protoAction,
			},
		)

		// Interactions.
		awaitUserAction = resp.Action != processorcomponent.ActionRespond
	}
}

func (s *Server) prepareEnvironment(ctx context.Context, sessionID uuid.UUID) (*environment, error) {
	const fn = "Server.prepareEnvironment"

	// Session.
	session, err := s.contextService.GetSession(ctx, &contextcomponent.GetSessionParams{
		SessionID: uuid.NullUUID{UUID: sessionID, Valid: true},
		Active:    null.BoolFrom(true),
	})
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get session: %w", fn, err)
	}

	// Processor initialization.
	processor, err := processorcomponent.NewProcessor(
		processorcomponent.Config{
			Session:          session,
			ClassifierAPIKey: s.cfg.ClassifierAPIKey,
			DB:               s.cfg.DB,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to create processor: %w", fn, err)
	}

	return &environment{
			Session:    session,
			Processor:  processor,
			ReqStorage: NewReqStorage(),
		},
		nil
}
