package repository

import (
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/storage"
	sqlpkg "github.com/Alp4ka/classifier-aaS/pkg/sql"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	storage.SQLStorage
	GetSession(ctx context.Context, dbtx sqlpkg.DBTX, filter *GetSessionFilter) (*Session, error)
	CreateSession(ctx context.Context, dbtx sqlpkg.DBTX, session Session) error
}

func NewRepository(db *sqlx.DB) Repository {
	return &repositoryImpl{
		SQLStorageImpl: &storage.SQLStorageImpl{
			DB: db,
		},
	}
}
