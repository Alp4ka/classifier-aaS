package grpc

import (
	"context"
	"fmt"

	contextcomponent "github.com/Alp4ka/classifier-aaS/internal/components/context"
	"github.com/Alp4ka/classifier-aaS/internal/telemetry"
	api "github.com/Alp4ka/classifier-api"
)

func (s *Server) AcquireSession(ctx context.Context, req *api.AcquireSessionRequest) (*api.AcquireSessionResponse, error) {
	const fn = "Server.AcquireSession"

	sess, err := s.contextService.CreateSession(ctx,
		&contextcomponent.CreateSessionParams{
			Agent:   req.GetAgent(),
			Gateway: req.GetGateway(),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to acquire session: %w", fn, err)
	}

	telemetry.T().IncrementSessionCount(sess.Model.Gateway)
	return &api.AcquireSessionResponse{SessionId: sess.Model.ID.String()}, nil
}
