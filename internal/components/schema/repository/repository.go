package repository

import (
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/storage"
	"github.com/jmoiron/sqlx"
)

type SchemaRepository interface {
	GetSchema(ctx context.Context, filter *GetSchemaFilter) (*Schema, error)
	CreateSchema(ctx context.Context, schema Schema) (*Schema, error)
	UpdateSchema(ctx context.Context, schema Schema) (*Schema, error)
}

type SchemaVariantRepository interface {
	GetSchemaVariant(ctx context.Context, filter *GetSchemaVariantFilter) (*SchemaVariant, error)
	CreateSchemaVariant(ctx context.Context, schema SchemaVariant) (*SchemaVariant, error)
}

type Repository interface {
	storage.Storage

	SchemaRepository
	SchemaVariantRepository
}

func NewRepository(db *sqlx.DB) Repository {
	return &repositoryImpl{
		Storage: storage.NewPostgresStorage(db),
	}
}
