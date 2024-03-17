package app

import (
	"context"
	"errors"
	"github.com/Alp4ka/classifier-aaS/internal/api"
	"github.com/Alp4ka/classifier-aaS/internal/contextcomponent"
	contextrepository "github.com/Alp4ka/classifier-aaS/internal/contextcomponent/repository"
	"github.com/Alp4ka/classifier-aaS/internal/schemacomponent"
	schemarepository "github.com/Alp4ka/classifier-aaS/internal/schemacomponent/repository"
)

type App struct {
	cfg Config

	httpServer *api.HTTPServer
	//gateways       []gateways.Gateway
}

func New(cfg Config) *App {
	return &App{
		cfg: cfg,
		httpServer: api.NewHTTPServer(
			api.Config{
				RateLimit: cfg.HTTPRateLimit,
				Port:      cfg.HTTPPort,
				ContextService: contextcomponent.NewService(
					contextcomponent.Config{
						Repository: contextrepository.NewRepository(cfg.DB),
					},
				),
				SchemaService: schemacomponent.NewService(
					schemacomponent.Config{
						Repository: schemarepository.NewRepository(cfg.DB),
					},
				),
			},
		),
	}
}

func (a *App) Run(ctx context.Context) error {
	errCh := make(chan error, 1)

	go func(errCh chan<- error) {
		errCh <- a.httpServer.Run()
	}(errCh)

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errCh:
		return err
	}
}

func (a *App) Close() error {
	return errors.Join(a.cfg.DB.Close(), a.httpServer.Close())
}
