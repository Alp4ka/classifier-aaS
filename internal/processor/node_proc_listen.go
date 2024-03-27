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
			userInputRequired: true,
		}, nil
	}
	return &response{
		pipeOutput: req.userInput,
	}, nil
}

var _ nodeProc = (*listenNodeProc)(nil)
