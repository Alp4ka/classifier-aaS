package repository

import (
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/storage"
	sqlpkg "github.com/Alp4ka/classifier-aaS/pkg/sql"
	"github.com/jmoiron/sqlx"
)

type SchemaRepository interface {
	GetSchema(ctx context.Context, dbtx sqlpkg.DBTx, filter *GetSchemaFilter) (*Schema, error)
	CreateSchema(ctx context.Context, dbtx sqlpkg.DBTx, schema Schema) error
	UpdateSchema(ctx context.Context, dbtx sqlpkg.Tx, schema Schema) error
}

type SchemaVariantRepository interface {
	GetSchemaVariant(ctx context.Context, dbtx sqlpkg.DBTx, filter *GetSchemaVariantFilter) (*SchemaVariant, error)
	CreateSchemaVariant(ctx context.Context, dbtx sqlpkg.DBTx, schema SchemaVariant) error
	UpdateSchemaVariant(ctx context.Context, dbtx sqlpkg.Tx, schema SchemaVariant) error
}

type Repository interface {
	storage.SQLStorage

	SchemaRepository
	SchemaVariantRepository
}

func NewRepository(db *sqlx.DB) Repository {
	return &repositoryImpl{
		SQLStorageImpl: &storage.SQLStorageImpl{
			Db: db,
		},
	}
}
