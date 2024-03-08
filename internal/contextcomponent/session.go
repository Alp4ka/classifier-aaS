package contextcomponent

import (
	"github.com/google/uuid"
)

type closeFnType = func() error

type Session struct {
	ID      uuid.UUID
	closeFn closeFnType
}

func NewSession(id uuid.UUID, closeFn closeFnType) *Session {
	return &Session{ID: id, closeFn: closeFn}
}

func (s *Session) Close() error {
	return s.closeFn()
}

var _ closeFnType = (*Session)(nil).Close
