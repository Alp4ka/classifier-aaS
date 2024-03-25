package schemacomponent

import (
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/schema"
	"github.com/google/uuid"
)

type Service interface {
	GetSchema(ctx context.Context, filter *GetSchemaFilter) (*schema.Schema, error)
	CreateSchema(ctx context.Context, params *CreateSchemaParams) (*schema.Schema, error)
	UpdateSchema(ctx context.Context, params *UpdateSchemaParams) (*schema.Schema, error)

	GetSchemaVariant(ctx context.Context, id uuid.UUID) (*schema.Variant, error)
}
