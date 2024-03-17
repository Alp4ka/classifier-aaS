package telegramchat

import (
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/gateways"
)

type Gateway struct {
}

func (g Gateway) Run(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (g Gateway) Close() error {
	//TODO implement me
	panic("implement me")
}

var _ gateways.Gateway = (*Gateway)(nil)
