package gateways

import "context"

type Gateway interface {
	Run(ctx context.Context) error
	Close() error
}
