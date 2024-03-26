package processor

import (
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/schema"
)

type classifyNodeProc struct {
	*schema.NodeClassify
}

func (l *classifyNodeProc) process(ctx context.Context, req *request) (*response, error) {
	// TODO: validate input.
	return defaultProcess(ctx, req)
}

var _ nodeProc = (*classifyNodeProc)(nil)
