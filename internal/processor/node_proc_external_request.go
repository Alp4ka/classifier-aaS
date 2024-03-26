package processor

import (
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/schema"
)

type externalRequestNodeProc struct {
	*schema.NodeExternalRequest
}

func (l *externalRequestNodeProc) process(ctx context.Context, req *request) (*response, error) {
	// TODO: validate input.
	return defaultProcess(ctx, req)
}

var _ nodeProc = (*externalRequestNodeProc)(nil)
