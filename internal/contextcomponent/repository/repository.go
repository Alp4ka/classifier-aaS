package repository

import (
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/storage"
	sqlpkg "github.com/Alp4ka/classifier-aaS/pkg/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, dbtx sqlpkg.DBTx, session Session) (*Session, error)
	GetSession(ctx context.Context, dbtx sqlpkg.DBTx, id uuid.UUID) (*Session, error)
	UpdateSession(ctx context.Context, tx sqlpkg.Tx, session Session) (*Session, error)
}

type EventRepository interface {
	CreateEvent(ctx context.Context, dbtx sqlpkg.DBTx, event Event) (*Event, error)
}

type Repository interface {
	storage.SQLStorage

	SessionRepository
	EventRepository
}

func NewRepository(db *sqlx.DB) Repository {
	return &repositoryImpl{
		SQLStorageImpl: &storage.SQLStorageImpl{
			Db: db,
		},
	}
}
