package repository

import (
	"github.com/Alp4ka/classifier-aaS/internal/storage"
)

type repositoryImpl struct {
	storage.Storage
}

var _ Repository = (*repositoryImpl)(nil)
