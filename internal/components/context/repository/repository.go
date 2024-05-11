package repository

import (
	"context"

	"github.com/Alp4ka/classifier-aaS/internal/storage"
	"github.com/jmoiron/sqlx"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, session Session) (*Session, error)
	GetSession(ctx context.Context, filter *GetSessionFilter) (*Session, error)
	UpdateSession(ctx context.Context, session Session) (*Session, error)
}

type Repository interface {
	storage.Storage

	SessionRepository
}

func NewRepository(db *sqlx.DB) Repository {
	return &repositoryImpl{
		Storage: storage.NewPostgresStorage(db),
	}
}
