package grpc

import (
	"context"
	"fmt"
	"github.com/Alp4ka/classifier-aaS/internal/contextcomponent"
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

	return &api.AcquireSessionResponse{SessionId: sess.Model.ID.String()}, nil
}
