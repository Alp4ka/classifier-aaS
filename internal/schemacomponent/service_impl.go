package schemacomponent

import (
	"github.com/Alp4ka/classifier-aaS/internal/schemacomponent/repository"
)

type Config struct {
	Repository repository.Repository
}

type serviceImpl struct {
	repo repository.Repository
}

func NewService(cfg Config) Service {
	return &serviceImpl{
		repo: cfg.Repository,
	}
}

var _ Service = (*serviceImpl)(nil)
