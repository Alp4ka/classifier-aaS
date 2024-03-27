package contextcomponent

import (
	"github.com/Alp4ka/classifier-aaS/internal/contextcomponent/repository"
	"github.com/Alp4ka/classifier-aaS/internal/schema"
	timepkg "github.com/Alp4ka/classifier-aaS/pkg/time"
	"github.com/google/uuid"
)

type Session struct {
	Model *repository.Session
	Tree  schema.Tree
}

func (s *Session) Active() bool {
	return s.Model.State == repository.SessionStateActive
}

func (s *Session) Operable() bool {
	return s.Active() && !s.Model.ClosedAt.Valid && !s.Expired()
}

func (s *Session) Expired() bool {
	return s.Model.ValidUntil.Before(timepkg.Now())
}

func (s *Session) ID() uuid.UUID {
	return s.Model.ID
}
