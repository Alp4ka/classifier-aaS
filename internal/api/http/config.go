package http

import (
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Port      int
	RateLimit int
	DB        *sqlx.DB
}
