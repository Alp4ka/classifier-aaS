package contextcomponent

import (
	"context"
	"github.com/google/uuid"
)

type Config struct {
}

type serviceImpl struct {
	cfg *Config

	doneCh chan struct{}
}

func NewService(cfg *Config) (Service, error) {
	return &serviceImpl{
		cfg:    cfg,
		doneCh: make(chan struct{}, 1),
	}, nil
}

func (s *serviceImpl) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	select {
	case <-ctx.Done():
		return nil
	case <-s.doneCh:
		return nil
	}
}

func (s *serviceImpl) Close() error {
	close(s.doneCh)

	return nil
}

func (s *serviceImpl) closeSession(ctx context.Context, sessionID uuid.UUID) error {
	return nil
}

type OpenSessionParams struct {
}

func (s *serviceImpl) GetSession(ctx context.Context, params *OpenSessionParams) (*Session, error) {
	return nil, nil
}

type Request struct {
	Req  string                 `json:"req"`
	Meta map[string]interface{} `json:"meta"`
}

type Response struct {
	Resp string                 `json:"resp"`
	Meta map[string]interface{} `json:"meta"`
}

func (s *serviceImpl) ProcessRequest(ctx context.Context, params *Request) (*Response, error) {
	return nil, nil
}
