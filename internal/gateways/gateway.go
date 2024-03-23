package gateways

import (
	"context"
	"github.com/google/uuid"
)

type GatewayType = string

type Response struct {
	EventID   int64
	SessionID uuid.UUID
	Response  string
	Error     error
}

type Gateway interface {
	Run(ctx context.Context) error
	Close() error
	Type() GatewayType
	Accept(response Response) bool
}
