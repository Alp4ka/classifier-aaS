package grpc

import (
	"context"
	"fmt"
	contextcomponent "github.com/Alp4ka/classifier-aaS/internal/components/context"
	processorcomponent "github.com/Alp4ka/classifier-aaS/internal/components/processor"
	"github.com/Alp4ka/classifier-aaS/internal/telemetry"
	timepkg "github.com/Alp4ka/classifier-aaS/pkg/time"
	api "github.com/Alp4ka/classifier-api"
	"github.com/Alp4ka/mlogger"
	"github.com/Alp4ka/mlogger/field"
	"github.com/google/uuid"
	"github.com/guregu/null/v5"
	"io"
)

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
	ctx = field.WithContextFields(ctx, field.String("session_id", req.GetSessionId()))

	// Init session.
	sessID, err := uuid.Parse(req.GetSessionId())
	if err != nil {
		return fmt.Errorf("session id wrong format: %w", err)
	}
	session, err := s.contextService.GetSession(ctx, &contextcomponent.GetSessionParams{
		SessionID: uuid.NullUUID{UUID: sessID, Valid: true},
		Active:    null.BoolFrom(true),
	})
	if err != nil {
		return fmt.Errorf("failed to get session: %w", err)
	}
	proc, err := processorcomponent.NewProcessor(session.Tree)
	if err != nil {
		return fmt.Errorf("failed to create processor: %w", err)
	}
	err = s.process(ctx, session, proc, src, req)
	if err != nil {
		return fmt.Errorf("failed to process: %w", err)
	}

	return nil
}

func (s *Server) process(
	ctx context.Context,
	sess *contextcomponent.Session,
	proc *processorcomponent.Processor,
	src api.GWManagerService_ProcessServer,
	initReq *api.ProcessRequest,
) error {
	const fn = "Server.process"

	// Metrics.
	{
		timeStart := timepkg.Now()
		defer func() {
			telemetry.T().ObserveProcessDuration(sess.Model.Gateway, timepkg.Now().Sub(timeStart))
		}()
	}

	var (
		req            *api.ProcessRequest
		firstIteration = true
	)

	for {
		if firstIteration {
			req = initReq
			firstIteration = false
		} else {
			// Reading request.
			var err error
			req, err = src.Recv()
			if err == io.EOF {
				return nil
			} else if err != nil {
				return fmt.Errorf("%s: failed to receive client request: %w", fn, err)
			}
		}

		ctx = field.WithContextFields(
			ctx,
			field.String("request_id", req.GetRequestId()),
			field.String("data", req.GetRequestData()),
		)
		// Deduplication.
		if s.IsDuplicate(req) {
			mlogger.L(ctx).Info("Duplicate request")
			continue // TODO: mb not continue or return error.
		}
		s.StoreRequest(req)

		// Session.
		if !sess.Operable() {
			return fmt.Errorf("session expired")
		}

		// Handle. TODO: pizdec
	handle:
		resp, err := proc.Handle(ctx, &processorcomponent.nodeRequest{Data: req.GetRequestData()})
		if err != nil {
			return fmt.Errorf("%s: failed to handle client request: %w", fn, err)
		}

		// Sending response.
		err = src.Send(&api.ProcessResponse{
			ResponseData:    resp.Output,
			End:             resp.End,
			RequestRequired: resp.InputRequired && resp.Output == nil, // TODO: Tozhe pizdec. Peredalat pod enum.
		})
		if err != nil {
			return fmt.Errorf("%s: failed to send response: %w", fn, err)
		}
		if resp.Output != nil {
			goto handle
		}
		if resp.End {
			return nil
		}
	}
}
