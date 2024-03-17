package app

import "github.com/jmoiron/sqlx"

type Config struct {
	DB            *sqlx.DB
	OpenAIAPIKey  string
	HTTPPort      int
	HTTPRateLimit int
}
