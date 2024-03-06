package repository

import (
	"github.com/Alp4ka/classifier-aaS/internal/storage"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	storage.SQLStorage
}

func NewRepository(db *sqlx.DB) Repository {
	return &repositoryImpl{
		SQLStorageImpl: &storage.SQLStorageImpl{
			DB: db,
		},
	}
}
