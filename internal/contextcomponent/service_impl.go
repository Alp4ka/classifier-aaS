package contextcomponent

import (
	"context"
	"errors"
	"fmt"
	"github.com/Alp4ka/classifier-aaS/internal/contextcomponent/repository"
	"github.com/Alp4ka/classifier-aaS/internal/schemacomponent"
	"github.com/Alp4ka/classifier-aaS/internal/storage"
	sqlpkg "github.com/Alp4ka/classifier-aaS/pkg/sql"
	timepkg "github.com/Alp4ka/classifier-aaS/pkg/time"
	"github.com/google/uuid"
	"github.com/guregu/null/v5"
	"time"
)

const DefaultSessionLifetime = time.Hour

type Config struct {
	SchemaService schemacomponent.Service
	Repository    repository.Repository
	OpenAIAPIKey  string
}

type serviceImpl struct {
	repo repository.Repository

	schemaService schemacomponent.Service
}

func NewService(cfg Config) Service {
	return &serviceImpl{
		repo:          cfg.Repository,
		schemaService: cfg.SchemaService,
	}
}

func (s *serviceImpl) ReleaseSession(ctx context.Context, sessionID uuid.UUID) error {
	const fn = "serviceImpl.ReleaseSession"

	err := s.repo.WithTransaction(
		ctx,
		func(innerCtx context.Context, tx sqlpkg.Tx) error {
			_, err := s.repo.UpdateSession(
				innerCtx,
				tx,
				repository.Session{
					ID:       sessionID,
					State:    repository.SessionStateClosedAgent,
					ClosedAt: null.TimeFrom(timepkg.Now()),
				},
			)
			return err
		},
	)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

type GetSessionParams struct {
	SessionID uuid.NullUUID

	Gateway null.String
	Agent   null.String
	Active  null.Bool
}

func (s *serviceImpl) GetSession(ctx context.Context, params *GetSessionParams) (*Session, error) {
	const fn = "serviceImpl.GetSession"

	var session *Session
	err := s.repo.WithTransaction(
		ctx,
		func(ctx context.Context, tx sqlpkg.Tx) error {
			sessionRecord, err := s.repo.GetSession(ctx, tx, &repository.GetSessionFilter{
				ID:      params.SessionID,
				Gateway: params.Gateway,
				Agent:   params.Agent,
				Active:  params.Active,
			})
			if err != nil {
				if errors.Is(err, storage.ErrEntityNotFound) {
					return ErrSessionDoesNotExist
				}
				return err
			}

			schemaVariantRecord, err := s.schemaService.GetSchemaVariant(ctx, sessionRecord.SchemaVariantID)
			if err != nil {
				if errors.Is(err, storage.ErrEntityNotFound) {
					return ErrNoActualVariant
				}
				return err
			}

			tree, err := schemaVariantRecord.Description.Map()
			if err != nil {
				return err
			}

			session = &Session{sessionRecord, tree}
			return nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return session, nil
}

type CreateSessionParams struct {
	SessionID uuid.UUID
	Agent     string
	Gateway   string
}

func (s *serviceImpl) CreateSession(ctx context.Context, params *CreateSessionParams) (*Session, error) {
	const fn = "serviceImpl.CreateSession"

	var session *Session
	err := s.repo.WithTransaction(
		ctx,
		func(ctx context.Context, tx sqlpkg.Tx) error {
			recentSchema, err := s.schemaService.GetSchema(ctx, &schemacomponent.GetSchemaFilter{Latest: null.BoolFrom(true)})
			if err != nil {
				return err
			} else if recentSchema.ActualVariant == nil {
				return ErrNoActualVariant
			}

			tree, err := recentSchema.ActualVariant.Description.Map()
			if err != nil {
				return err
			}
			startNode, err := tree.GetStart()
			if err != nil {
				return err
			}

			sessionRecord, err := s.repo.CreateSession(ctx, tx, repository.Session{
				ID:              params.SessionID,
				State:           repository.SessionStateActive,
				Agent:           params.Agent,
				Gateway:         params.Gateway,
				ValidUntil:      timepkg.Now().Add(DefaultSessionLifetime),
				ClosedAt:        null.TimeFromPtr(nil),
				SchemaVariantID: recentSchema.ActualVariant.ID,
				SchemaNodeID:    startNode.GetID(),
			})
			if err != nil {
				return err
			}

			session = &Session{sessionRecord, tree}
			return nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return session, nil
}

type AcquireSessionParams struct {
	Agent   string
	Gateway string
}

func (s *serviceImpl) AcquireSession(ctx context.Context, params *AcquireSessionParams) (*Session, error) {
	const fn = "serviceImpl.AcquireSession"

	session, err := s.GetSession(ctx, &GetSessionParams{
		Gateway: null.StringFrom(params.Gateway),
		Agent:   null.StringFrom(params.Agent),
		Active:  null.BoolFrom(true),
	})
	if err == nil {
		if !session.Model.ValidUntil.Before(timepkg.Now()) &&
			session.Model.State == repository.SessionStateActive &&
			!session.Model.ClosedAt.Valid {
			return session, nil
		}
	} else if !errors.Is(err, ErrSessionDoesNotExist) {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	session, err = s.CreateSession(ctx, &CreateSessionParams{
		SessionID: uuid.New(),
		Agent:     params.Agent,
		Gateway:   params.Gateway,
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return session, nil
}

type Request struct {
	Req  string                 `json:"req"`
	Meta map[string]interface{} `json:"meta,omitempty"`
}

type Response struct {
	Resp string                 `json:"resp"`
	Meta map[string]interface{} `json:"meta,omitempty"`
}

func (s *serviceImpl) Handle(ctx context.Context, session *Session, req *Request) (*Response, error) {
	panic("pizdec")
	// Process request corresponding to step.
	// Save request.

	// Update context.
	// Commit.
	return nil, nil
}

var _ Service = (*serviceImpl)(nil)
