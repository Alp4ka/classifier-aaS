package schema

import (
	"github.com/Alp4ka/classifier-aaS/internal/components/schema/entities"
	"github.com/google/uuid"
	"time"
)

type VariantReq struct {
	ID          uuid.UUID            `json:"id,omitempty"`
	Description entities.Description `json:"description"`
	CreatedAt   time.Time            `json:"createdAt"`
	UpdatedAt   time.Time            `json:"updatedAt"`
}

type SchemaReq struct {
	ID            uuid.UUID   `json:"id"`
	ActualVariant *VariantReq `json:"actualVariant"`
	CreatedAt     time.Time   `json:"createdAt"`
	UpdatedAt     time.Time   `json:"updatedAt"`
}
