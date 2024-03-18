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
	"time"
)

const DefaultSessionLifetime time.Duration = time.Hour

type Config struct {
	Repository   repository.Repository
	OpenAIAPIKey string
}

type serviceImpl struct {
	repo repository.Repository

	schemaService schemacomponent.Service
}

func NewService(cfg Config) Service {
	return &serviceImpl{
		repo: cfg.Repository,
	}
}

func (s *serviceImpl) closeSession(ctx context.Context, sessionID uuid.UUID) error {
	const fn = "serviceImpl.closeSession"

	err := s.repo.WithTransaction(
		ctx,
		func(innerCtx context.Context, tx sqlpkg.Tx) error {
			_, err := s.repo.UpdateSession(
				innerCtx,
				tx,
				repository.Session{
					ID:    sessionID,
					State: repository.SessionStateClosedAgent,
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

type AcquireSessionParams struct {
	SessionID uuid.UUID
	Agent     string
	Gateway   string
}

func (s *serviceImpl) AcquireSession(ctx context.Context, params *AcquireSessionParams) (*Session, error) {
	const fn = "serviceImpl.AcquireSession"
	if params == nil {
		panic("how the fuck did you get here?")
	}

	var session *Session
	err := s.repo.WithTransaction(
		ctx,
		func(ctx context.Context, tx sqlpkg.Tx) error {
			sessionRecord, err := s.repo.GetSession(
				ctx,
				s.repo.DB(),
				&repository.GetSessionFilter{
					ID:          params.SessionID,
					CurrentTime: timepkg.TimeNow(),
				},
			)
			if err != nil {
				if !errors.Is(err, storage.ErrEntityNotFound) {
					return err
				}

				sessionRecord = &repository.Session{
					ID:         params.SessionID,
					State:      repository.SessionStateActive,
					Agent:      params.Agent,
					Gateway:    params.Gateway,
					ValidUntil: timepkg.TimeNow().Add(DefaultSessionLifetime),
				}
				_, err = s.repo.CreateSession(ctx, s.repo.DB(), *sessionRecord)
				if err != nil {
					return err
				}
			}

			session = NewSession(
				sessionRecord,
				func(ctx context.Context) error {
					return s.closeSession(ctx, sessionRecord.ID)
				},
			)
			return nil
		},
	)
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
	if req == nil {
		panic("how the fuck did you get here?")
	}
	panic("not implemented")

	// Load schema if not presented.
	// Create context if not presented.

	// Get step of schema.
	// Process request corresponding to step.
	// Save request.

	// Update context.
	// Commit.
	return nil, nil
}

var _ Service = (*serviceImpl)(nil)
