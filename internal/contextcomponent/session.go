package contextcomponent

import (
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/contextcomponent/repository"
)

type closeFnType = func(ctx context.Context) error

type Session struct {
	Model   *repository.Session
	closeFn closeFnType
}

func NewSession(model *repository.Session, closeFn closeFnType) *Session {
	return &Session{Model: model, closeFn: closeFn}
}

func (s *Session) Close(ctx context.Context) error {
	return s.closeFn(ctx)
}

var _ closeFnType = (*Session)(nil).Close
