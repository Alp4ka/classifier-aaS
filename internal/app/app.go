package app

import (
	"context"
	"errors"
	"github.com/Alp4ka/classifier-aaS/internal/api/grpc"
	"github.com/Alp4ka/classifier-aaS/internal/api/http"
)

type App struct {
	cfg Config

	httpServer *http.Server
	grpcServer *grpc.Server
}

func New(cfg Config) *App {
	return &App{
		cfg: cfg,
		httpServer: http.NewHTTPServer(
			http.Config{
				RateLimit: cfg.HTTPRateLimit,
				Port:      cfg.HTTPPort,
				DB:        cfg.DB,
			},
		).WithMetrics(),
		grpcServer: grpc.New(
			grpc.Config{
				Port:             cfg.GRPCPort,
				ClassifierAPIKey: cfg.OpenAIAPIKey,
				DB:               cfg.DB,
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
