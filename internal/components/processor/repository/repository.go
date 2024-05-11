package repository

import (
	"context"

	"github.com/Alp4ka/classifier-aaS/internal/storage"
	"github.com/jmoiron/sqlx"
)

type EventRepository interface {
	CreateEvent(ctx context.Context, event Event) (*Event, error)
}

type Repository interface {
	storage.Storage

	EventRepository
}

func NewRepository(db *sqlx.DB) Repository {
	return &repositoryImpl{
		Storage: storage.NewPostgresStorage(db),
	}
}
