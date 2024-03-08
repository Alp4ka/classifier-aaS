package contextcomponent

import (
	"context"
)

type Service interface {
	Run(ctx context.Context) error
	Close() error
	GetSession(ctx context.Context, params *OpenSessionParams) (*Session, error)
	ProcessRequest(ctx context.Context, event *Request) (*Response, error)
}
