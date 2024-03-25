package contextcomponent

import (
	"context"
	"github.com/google/uuid"
)

type Service interface {
	GetSession(ctx context.Context, params *GetSessionParams) (*Session, error)
	CreateSession(ctx context.Context, params *CreateSessionParams) (*Session, error)
	AcquireSession(ctx context.Context, params *AcquireSessionParams) (*Session, error)
	ReleaseSession(ctx context.Context, sessionID uuid.UUID) error
	Handle(ctx context.Context, session *Session, req *Request) (*Response, error)
}
