package app

import "github.com/jmoiron/sqlx"

type Config struct {
	DB           *sqlx.DB
	OpenAIAPIKey string

	// HTTP.
	HTTPPort      int
	HTTPRateLimit int

	// GRPC.
	GRPCPort int
}
