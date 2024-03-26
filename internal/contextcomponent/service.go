package contextcomponent

import (
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/contextcomponent/repository"
	"github.com/google/uuid"
)

type Service interface {
	GetSession(ctx context.Context, params *GetSessionParams) (*Session, error)
	CreateSession(ctx context.Context, params *CreateSessionParams) (*Session, error)
	AcquireSession(ctx context.Context, params *AcquireSessionParams) (*Session, error)
	ReleaseSession(ctx context.Context, sessionID uuid.UUID, state repository.SessionState) error
}
