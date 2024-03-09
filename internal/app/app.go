package app

import (
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/contextcomponent"
	"github.com/Alp4ka/classifier-aaS/internal/gateways"
	"github.com/Alp4ka/classifier-aaS/internal/schemacomponent"
)

type App struct {
	cfg *Config

	contextService contextcomponent.Service
	schemaService  schemacomponent.Service
	gateways       []gateways.Gateway
}

func New(cfg *Config) *App {
	panic("configure services")
	return &App{
		cfg: cfg,
	}
}

func (a *App) Run(ctx context.Context) error {
	panic("implement me")
	return nil
}
