package processor

import (
	"context"
	"fmt"
	"github.com/Alp4ka/classifier-aaS/internal/schema"
)

type Processor struct {
	ctx     context.Context
	tree    tree
	curNode nodeProc

	respChan chan *Response
}

func NewProcessor(schemaTree schema.Tree) (*Processor, error) {
	const fn = "NewProcessor"

	t := make(tree)
	err := t.fromSchemaTree(schemaTree)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &Processor{tree: t}, nil
}

type Request struct {
	Data string
}

type Response struct {
	Data            string
	End             bool
	RequestRequired bool
}

func (p *Processor) Start(ctx context.Context) <-chan *Response {
	p.ctx = ctx
	p.respChan = make(chan *Response)

	return p.respChan
}

func (p *Processor) Handle(req *Request) error {
	const fn = "Processor.Handle"

	// TODO: Save current node id every step.
	// TODO: Events journal.
	if p.curNode == nil {
		start, err := p.tree.getStart()
		if err != nil {
			return nil, fmt.Errorf("%s: %w", fn, err)
		}
		p.curNode = start
	}

	return &Response{Data: "zhopa123"}, nil
}
