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
		dataType:        l.OutputType(),
		end:             true,
		requestRequired: false,
	}, nil
}

var _ nodeProc = (*finishNodeProc)(nil)
