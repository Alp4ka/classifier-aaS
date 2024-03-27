package processor

import (
	"context"
	"fmt"
	"github.com/Alp4ka/classifier-aaS/internal/schema"
	"sync/atomic"
)

type Processor struct {
	ctx    context.Context
	cancel context.CancelFunc

	tree tree

	reqExpect atomic.Bool
	reqChan   chan *Request
	respChan  chan *Response
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
	Data             string
	End              bool
	RequestRequired  bool
	ResponseRequired bool
	Error            error
}

func (p *Processor) Start(ctx context.Context) (<-chan *Response, error) {
	start, err := p.tree.getStart()
	if err != nil {

		return nil, err
	}
	curNode := start

	// Process
	p.ctx, p.cancel = context.WithCancel(ctx)
	p.respChan = make(chan *Response)
	p.reqChan = make(chan *Request)

	go func() {
		var req *Request

		nresp, nerr := curNode.process(p.ctx, &request{})
		if nerr != nil {
			p.respChan <- &Response{
				Data:            "",
				End:             true,
				Error:           nerr,
				RequestRequired: false,
			}
			return
		}

		for {
			select {
			case <-p.ctx.Done():
				close(p.respChan)
				return
			case req = <-p.reqChan:
			}
		}
	}()

	return p.respChan
}

func (p *Processor) Handle(req *Request) {
	const fn = "Processor.Handle"

	// TODO: Save current node id every step.
	// TODO: Events journal.
	// TODO: If expecting message.
	p.reqChan <- req
}

func (p *Processor) Close() {
	close(p.reqChan)
	p.cancel()
}
