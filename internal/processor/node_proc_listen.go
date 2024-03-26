package processor

import (
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/schema"
)

type listenNodeProc struct {
	*schema.NodeListen
	cnt int
}

func (l *listenNodeProc) process(ctx context.Context, req *request) (*response, error) {
	// TODO: validate input.
	if l.cnt == 0 {
		l.cnt++
		return &response{
			end:             false,
			requestRequired: true,
		}, nil
	}
	return &response{
		data:     req.data,
		dataType: l.OutputType(),
		end:      false,
	}, nil
}

var _ nodeProc = (*listenNodeProc)(nil)
