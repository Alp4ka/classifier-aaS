package processor

import (
	"fmt"
	contextcomponent "github.com/Alp4ka/classifier-aaS/internal/components/context"
	"github.com/Alp4ka/classifier-aaS/internal/components/processor/repository"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Session          *contextcomponent.Session
	ClassifierAPIKey string
	DB               *sqlx.DB
}

type Processor struct {
	tree         tree
	currentNode  node
	systemConfig *systemConfig
	repository   repository.Repository
	scope        scope
	sessionID    uuid.UUID
}

func NewProcessor(cfg Config) (*Processor, error) {
	const fn = "NewProcessor"

	t, err := newTree(cfg.Session.Tree)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	start, err := t.GetStart()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &Processor{
			tree:        t,
			currentNode: start,
			systemConfig: &systemConfig{
				ClassifierAPIKey: cfg.ClassifierAPIKey,
			},
			repository: repository.NewRepository(cfg.DB),
			scope:      make(map[scopeVariable]string),
			sessionID:  cfg.Session.ID(),
		},
		nil
}
