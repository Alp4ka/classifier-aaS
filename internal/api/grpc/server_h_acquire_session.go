package grpc

import (
	"context"
	"fmt"
	"github.com/Alp4ka/classifier-aaS/internal/contextcomponent"
	"github.com/Alp4ka/classifier-aaS/internal/telemetry"
	api "github.com/Alp4ka/classifier-api"
)

func (s *Server) AcquireSession(ctx context.Context, req *api.AcquireSessionRequest) (*api.AcquireSessionResponse, error) {
	sess, err := s.contextService.AcquireSession(ctx,
		&contextcomponent.AcquireSessionParams{
			Agent:   req.GetAgent(),
			Gateway: req.GetGateway(),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to acquire session: %w", err)
	}

	telemetry.T().IncrementSessionCount(sess.Model.Gateway)
	return &api.AcquireSessionResponse{SessionId: sess.Model.ID.String()}, nil
}
