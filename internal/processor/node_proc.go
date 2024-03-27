package processor

import (
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/schema"
)

type request struct {
	data any
}

type response struct {
	pipeOutput       any
	resp             string
	err              error
	end              bool
	requestRequired  bool
	responseRequired bool
}

type procFunc func(ctx context.Context, req *request) (*response, error)

type nodeProc interface {
	schema.Node
	process(ctx context.Context, req *request) (*response, error)
}

var _ procFunc = (nodeProc)(nil).process
