package processor

import (
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/schema"
)

type externalRequestNodeProc struct {
	*schema.NodeExternalRequest
}

func (l *externalRequestNodeProc) process(ctx context.Context, req *request) (*response, error) {
	panic("implement me")
}

var _ nodeProc = (*externalRequestNodeProc)(nil)
