package grpc

import (
	"github.com/Alp4ka/classifier-aaS/internal/components/context"
)

type Config struct {
	Port           int
	ContextService context.Service
}
