package api

import (
	"github.com/Alp4ka/classifier-aaS/internal/contextcomponent"
	"github.com/Alp4ka/classifier-aaS/internal/schemacomponent"
)

type Config struct {
	Port           int
	ContextService contextcomponent.Service
	SchemaService  schemacomponent.Service
	RateLimit      int
}
