package app

import (
	"context"
	"errors"
	"github.com/Alp4ka/classifier-aaS/internal/api/grpc"
	"github.com/Alp4ka/classifier-aaS/internal/api/http"
	contextcomponent "github.com/Alp4ka/classifier-aaS/internal/components/context"
	contextrepository "github.com/Alp4ka/classifier-aaS/internal/components/context/repository"
	schemacomponent "github.com/Alp4ka/classifier-aaS/internal/components/schema"
	schemarepository "github.com/Alp4ka/classifier-aaS/internal/components/schema/repository"
)

type App struct {
	cfg Config

	httpServer *http.Server
	grpcServer *grpc.Server
}

func New(cfg Config) *App {
	schemaService := schemacomponent.NewService(
		schemacomponent.Config{
			Repository: schemarepository.NewRepository(cfg.DB),
		},
	)

	contextService := contextcomponent.NewService(
		contextcomponent.Config{
			SchemaService: schemaService,
			Repository:    contextrepository.NewRepository(cfg.DB),
		},
	)

	return &App{
		cfg: cfg,
		httpServer: http.NewHTTPServer(
			http.Config{
				RateLimit:     cfg.HTTPRateLimit,
				Port:          cfg.HTTPPort,
				SchemaService: schemaService,
			},
		).WithMetrics(),
		grpcServer: grpc.New(
			grpc.Config{
				ContextService: contextService,
				Port:           cfg.GRPCPort,
			},
		),
	}
}

func (a *App) Run(ctx context.Context) error {
	errCh := make(chan error, 2)

	go func(errCh chan<- error) {
		errCh <- a.httpServer.Run()
	}(errCh)

	go func(errCh chan<- error) {
		errCh <- a.grpcServer.Run()
	}(errCh)

	select {
	case <-ctx.Done():
		return nil
	case err := <-errCh:
		return err
	}
}

func (a *App) Close() (err error) {
	return errors.Join(a.httpServer.Close(), a.grpcServer.Close(), a.cfg.DB.Close())
}
