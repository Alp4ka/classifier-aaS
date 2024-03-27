package config

import (
	"fmt"
	"github.com/caarlos0/env"
)

type Config struct {
	// Postgres.
	PgDSN                string `env:"PG_DSN" envDefault:"postgres://db:db@localhost:228/classifier-aas?sslmode=disable"`
	PgMaxOpenConnections int    `env:"PG_MAX_OPEN_CONNS" envDefault:"50"`
	PgMaxIdleConnections int    `env:"PG_MAX_IDLE_CONNS" envDefault:"50"`

	// Interactions.
	OpenAIAPIKey string `env:"OPENAI_API_KEY"`

	// HTTP API.
	HTTPPort  int `env:"HTTP_PORT" envDefault:"1111"`
	RateLimit int `env:"HTTP_RATE_LIMIT" envDefault:"10"`

	// GRPC API.
	GRPCPort int `env:"GRPC_PORT" envDefault:"1113"`
}

func FromEnv() (*Config, error) {
	cfg := &Config{}

	err := env.Parse(cfg)
	if err != nil {
		return nil, fmt.Errorf("config.FromEnv: %w", err)
	}

	return cfg, nil
}
