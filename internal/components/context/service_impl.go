package context

import (
	"context"
	"errors"
	"fmt"
	"time"

	contextrepository "github.com/Alp4ka/classifier-aaS/internal/components/context/repository"
	schemacomponent "github.com/Alp4ka/classifier-aaS/internal/components/schema"
	"github.com/Alp4ka/classifier-aaS/internal/storage"
	timepkg "github.com/Alp4ka/classifier-aaS/pkg/time"
	"github.com/google/uuid"
	"github.com/guregu/null/v5"
)

const DefaultSessionLifetime = time.Minute * 10

type Config struct {
	SchemaService schemacomponent.Service
	Repository    contextrepository.Repository
	OpenAIAPIKey  string
}

type serviceImpl struct {
	repo contextrepository.Repository

	schemaService schemacomponent.Service
}

func NewService(cfg Config) Service {
	return &serviceImpl{
		repo:          cfg.Repository,
		schemaService: cfg.SchemaService,
	}
}

func (s *serviceImpl) ReleaseSession(ctx context.Context, sessID uuid.UUID, state contextrepository.SessionState) error {
	const fn = "serviceImpl.ReleaseSession"

	err := s.repo.WithTransaction(
		ctx,
		func(innerCtx context.Context) error {
			sess, err := s.repo.GetSession(innerCtx,
				&contextrepository.GetSessionFilter{
					ID: uuid.NullUUID{UUID: sessID, Valid: true},
				})
			if err != nil {
				if errors.Is(err, storage.ErrEntityNotFound) {
					return ErrSessionDoesNotExist
				}
				return fmt.Errorf("unable to find session: %w", err)
			} else if sess.State != contextrepository.SessionStateActive { // Already closed.
				return nil
			}

			_, err = s.repo.UpdateSession(
				innerCtx,
				contextrepository.Session{
					ID:       sessID,
					State:    state,
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
		func(ctx context.Context) error {
			sessionRecord, err := s.repo.GetSession(ctx, &contextrepository.GetSessionFilter{
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

			tree, err := schemaVariantRecord.Description.MapAndValidate()
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
		func(ctx context.Context) error {
			recentSchema, err := s.schemaService.GetSchema(ctx, &schemacomponent.GetSchemaFilter{Latest: null.BoolFrom(true)})
			if err != nil {
				return err
			} else if recentSchema.ActualVariant == nil {
				return ErrNoActualVariant
			}

			tree, err := recentSchema.ActualVariant.Description.MapAndValidate()
			if err != nil {
				return err
			}
			startNode, err := tree.GetStart()
			if err != nil {
				return err
			}

			sessionRecord, err := s.repo.CreateSession(ctx, contextrepository.Session{
				ID:              params.SessionID,
				State:           contextrepository.SessionStateActive,
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

var _ Service = (*serviceImpl)(nil)
