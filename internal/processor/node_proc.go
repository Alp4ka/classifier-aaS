package processor

import (
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/schema"
)

type request struct {
	pipeInput any
	userInput string
}

type response struct {
	pipeOutput any
	pipeErr    error
	pipeEnd    bool

	userOutput        *string
	userInputRequired bool
}

func (r *response) fall() bool {
	return !r.pipeEnd && r.userOutput == nil && !r.userInputRequired
}

type nodeProc interface {
	schema.Node
	process(ctx context.Context, req *request) (*response, error)
}
