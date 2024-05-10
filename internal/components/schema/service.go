package schema

import (
	"context"
	"github.com/google/uuid"
)

type Service interface {
	GetSchema(ctx context.Context, filter *GetSchemaFilter) (*SchemaReq, error)
	CreateSchema(ctx context.Context, params *CreateSchemaParams) (*SchemaReq, error)
	UpdateSchema(ctx context.Context, params *UpdateSchemaParams) (*SchemaReq, error)

	GetSchemaVariant(ctx context.Context, id uuid.UUID) (*VariantReq, error)
}
