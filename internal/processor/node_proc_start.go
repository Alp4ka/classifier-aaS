package processor

import (
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/schema"
)

type startNodeProc struct {
	*schema.NodeStart
}

func (l *startNodeProc) process(ctx context.Context, req *request) (*response, error) {
	return &response{
		pipeOutput: nil,
	}, nil
}

var _ nodeProc = (*startNodeProc)(nil)
