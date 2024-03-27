package processor

import (
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/schema"
	"strings"
	"sync"
)

type respondNodeProc struct {
	*schema.NodeRespond
	cnt int

	once      sync.Once
	formatted string
}

func (l *respondNodeProc) process(ctx context.Context, req *request) (*response, error) {
	// TODO: validate input.
	const inputFlag = "{input}"

	l.once.Do(func() {
		var input string
		if req.pipeInput != nil {
			input = req.pipeInput.(string)
		}
		l.formatted = strings.ReplaceAll(l.Response, inputFlag, input)
	})

	if l.cnt == 0 {
		l.cnt++
		return &response{
			userOutput: &l.formatted,
		}, nil
	}
	return &response{
		pipeOutput: l.formatted,
	}, nil
}

var _ nodeProc = (*respondNodeProc)(nil)
