package schema

import (
	"github.com/google/uuid"
	"time"
)

type Variant struct {
	ID          uuid.UUID   `json:"id,omitempty"`
	Description Description `json:"description"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

type Schema struct {
	ID            uuid.UUID `json:"id"`
	ActualVariant *Variant  `json:"actualVariant"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
