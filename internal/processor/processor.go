package processor

import (
	"context"
	"fmt"
	"github.com/Alp4ka/classifier-aaS/internal/schema"
)

type Processor struct {
	tree    schema.Tree
	curNode schema.Node
}

func NewProcessor(tree schema.Tree) *Processor {
	return &Processor{tree: tree}
}

type Request struct {
	Data string
}

type Response struct {
	Data string
	End  bool
}

func (p *Processor) Handle(ctx context.Context, req *Request) (*Response, error) {
	const fn = "Processor.Handle"

	// TODO: Save current node id every step.
	if p.curNode == nil {
		start, err := p.tree.GetStart()
		if err != nil {
			return nil, fmt.Errorf("%s: %w", fn, err)
		}
		p.curNode = start
		fmt.Println(p.curNode)
	}

	return &Response{Data: "zhopa123"}, nil
}
