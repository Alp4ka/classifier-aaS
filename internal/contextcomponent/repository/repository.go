package repository

import (
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/storage"
	sqlpkg "github.com/Alp4ka/classifier-aaS/pkg/sql"
	"github.com/jmoiron/sqlx"
)

type SessionRepository interface {
	GetSession(ctx context.Context, dbtx sqlpkg.DBTx, filter *GetSessionFilter) (*Session, error)
	CreateSession(ctx context.Context, dbtx sqlpkg.DBTx, session Session) (*Session, error)
	UpdateSession(ctx context.Context, dbtx sqlpkg.Tx, session Session) (*Session, error)
}

type ContextRepository interface {
	GetContext(ctx context.Context, dbtx sqlpkg.DBTx, filter *GetContextFilter) (*Context, error)
	CreateContext(ctx context.Context, dbtx sqlpkg.DBTx, context Context) (*Context, error)
	UpdateContext(ctx context.Context, dbtx sqlpkg.Tx, context Context) (*Context, error)
}

type EventRepository interface {
	CreateEvent(ctx context.Context, dbtx sqlpkg.DBTx, event Event) error
}

type Repository interface {
	storage.SQLStorage

	SessionRepository
	ContextRepository
	EventRepository
}

func NewRepository(db *sqlx.DB) Repository {
	return &repositoryImpl{
		SQLStorageImpl: &storage.SQLStorageImpl{
			Db: db,
		},
	}
}
