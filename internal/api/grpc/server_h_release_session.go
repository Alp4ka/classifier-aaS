package grpc

import (
	"context"
	"fmt"

	contextrepository "github.com/Alp4ka/classifier-aaS/internal/components/context/repository"
	api "github.com/Alp4ka/classifier-api"
	"github.com/google/uuid"
)

func (s *Server) ReleaseSession(ctx context.Context, req *api.ReleaseSessionRequest) (*api.ReleaseSessionResponse, error) {
	const fn = "Server.ReleaseSession"

	sessID, err := uuid.Parse(req.GetSessionId())
	if err != nil {
		return nil, fmt.Errorf("%s: unable to parse session uuid: %w", fn, err)
	}

	err = s.contextService.ReleaseSession(ctx, sessID, contextrepository.SessionStateClosedGateway)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to release session: %w", fn, err)
	}

	return &api.ReleaseSessionResponse{}, nil
}
