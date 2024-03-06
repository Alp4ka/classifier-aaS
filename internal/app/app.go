package app

import "context"

type App struct {
	cfg *Config
}

func New(cfg *Config) *App {
	return &App{
		cfg: cfg,
	}
}

func (a *App) Run(ctx context.Context) error {
	return nil
}
