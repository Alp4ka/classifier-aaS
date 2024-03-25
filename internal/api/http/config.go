package http

import (
	"github.com/Alp4ka/classifier-aaS/internal/schemacomponent"
)

type Config struct {
	Port          int
	SchemaService schemacomponent.Service
	RateLimit     int
}
