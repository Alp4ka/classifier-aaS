package contextcomponent

import (
	"github.com/Alp4ka/classifier-aaS/internal/contextcomponent/repository"
	"github.com/Alp4ka/classifier-aaS/internal/schema"
	timepkg "github.com/Alp4ka/classifier-aaS/pkg/time"
)

type Session struct {
	Model *repository.Session
	Tree  schema.Tree
}

func (s *Session) Active() bool {
	return s.Model.ValidUntil.After(timepkg.Now()) &&
		s.Model.State == repository.SessionStateActive &&
		!s.Model.ClosedAt.Valid
}
