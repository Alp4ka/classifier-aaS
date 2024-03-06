package telegramchat

import (
	"github.com/Alp4ka/classifier-aaS/internal/gateways"
)

type Gateway struct {
}

var _ gateways.Gateway = (*Gateway)(nil)
