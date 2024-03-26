package grpc

import (
	"context"
	"fmt"
	contextrepository "github.com/Alp4ka/classifier-aaS/internal/contextcomponent/repository"
	"github.com/Alp4ka/classifier-aaS/pkg/api"
	"github.com/google/uuid"
)

func (s *Server) ReleaseSession(ctx context.Context, req *api.ReleaseSessionRequest) (*api.ReleaseSessionResponse, error) {
	sessionID, err := uuid.Parse(req.GetSessionId())
	if err != nil {
		return nil, fmt.Errorf("unable to parse session uuid: %w", err)
	}

	err = s.contextService.ReleaseSession(ctx, sessionID, contextrepository.SessionStateClosedGateway)
	if err != nil {
		return nil, fmt.Errorf("failed to release session: %w", err)
	}

	return &api.ReleaseSessionResponse{}, nil
}
