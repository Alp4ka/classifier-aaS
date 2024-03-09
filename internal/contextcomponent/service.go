package contextcomponent

import (
	"context"
)

type Service interface {
	AcquireSession(ctx context.Context, params *AcquireSessionParams) (*Session, error)
	Handle(ctx context.Context, session *Session, req *Request) (*Response, error)
}
