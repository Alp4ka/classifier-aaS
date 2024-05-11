package grpc

import (
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Port             int
	ClassifierAPIKey string
	DB               *sqlx.DB
}
