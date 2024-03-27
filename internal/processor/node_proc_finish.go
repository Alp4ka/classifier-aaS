package processor

import (
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/schema"
)

type finishNodeProc struct {
	*schema.NodeFinish
}

func (l *finishNodeProc) process(ctx context.Context, req *request) (*response, error) {
	return &response{
		pipeEnd: true,
	}, nil
}

var _ nodeProc = (*finishNodeProc)(nil)
