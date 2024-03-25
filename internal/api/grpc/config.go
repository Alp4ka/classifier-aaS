package grpc

import (
	"github.com/Alp4ka/classifier-aaS/internal/contextcomponent"
)

type Config struct {
	Port           int
	ContextService contextcomponent.Service
}
