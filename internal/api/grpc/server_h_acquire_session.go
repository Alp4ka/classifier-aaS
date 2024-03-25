package grpc

import (
	"context"
	"fmt"
	"github.com/Alp4ka/classifier-aaS/internal/contextcomponent"
	"github.com/Alp4ka/classifier-aaS/pkg/api"
	"github.com/google/uuid"
)

func (s *Server) AcquireSession(ctx context.Context, req *api.AcquireSessionRequest) (*api.AcquireSessionResponse, error) {
	sessionID, err := uuid.Parse(req.GetSessionId())
	if err != nil {
		return nil, fmt.Errorf("unable to parse session uuid: %w", err)
	}

	sess, err := s.contextService.AcquireSession(ctx,
		&contextcomponent.AcquireSessionParams{
			SessionID: sessionID,
			Agent:     req.GetAgent(),
			Gateway:   req.GetGateway(),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to acquire session: %w", err)
	}

	return &api.AcquireSessionResponse{SessionId: sess.Model.ID.String()}, nil
}
