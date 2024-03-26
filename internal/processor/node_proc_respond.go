package processor

import (
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/schema"
)

type respondNodeProc struct {
	*schema.NodeRespond
}

func (l *respondNodeProc) process(ctx context.Context, req *request) (*response, error) {
	// TODO: validate input.
	return &response{
		data:     l.Response,
		dataType: l.OutputType(),
		responseRequired:
	}, nil
}

var _ nodeProc = (*respondNodeProc)(nil)
