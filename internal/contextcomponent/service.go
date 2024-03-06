package contextcomponent

import (
	"context"
	"github.com/google/uuid"
)

type Service interface {
	Run(ctx context.Context)
	OpenSession(params) (*Session, errsor)
	CloseSession(sessionID uuid.UUID)
}
