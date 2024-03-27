package processor

import (
	"context"
	"fmt"
	"github.com/Alp4ka/classifier-aaS/internal/schema"
	"github.com/Alp4ka/mlogger"
	"github.com/Alp4ka/mlogger/field"
)

type Processor struct {
	tree    tree
	curNode nodeProc
	nReq    *request
}

func NewProcessor(schemaTree schema.Tree) (*Processor, error) {
	const fn = "NewProcessor"

	t := make(tree)
	err := t.fromSchemaTree(schemaTree)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	start, err := t.getStart()
	if err != nil {

		return nil, err
	}

	return &Processor{tree: t, curNode: start, nReq: new(request)}, nil
}

type Request struct {
	Data string
}

type Response struct {
	Output        *string
	InputRequired bool
	End           bool
}

func (p *Processor) Handle(ctx context.Context, req *Request) (*Response, error) {
	const fn = "Processor.Handle"

	p.nReq.userInput = req.Data
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		mlogger.L(ctx).Info("Processing node",
			field.String("id", p.curNode.GetID().String()),
			field.String("type", string(p.curNode.GetType())),
		)
		nRet, err := p.curNode.process(ctx, p.nReq)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", fn, err)
		}

		// Return conditions.
		if nRet.pipeEnd {
			return &Response{
				End: true,
			}, nil
		} else if !nRet.fall() {
			return &Response{
				Output:        nRet.userOutput,
				InputRequired: nRet.userInputRequired,
				End:           false,
			}, nil
		}

		p.nReq.pipeInput = nRet.pipeOutput
		// Success path.
		if nRet.pipeErr == nil {
			if p.curNode.GetNextID().Valid { // goto finish.
				nextNode, err := p.tree.get(p.curNode.GetNextID().UUID)
				if err != nil {
					return nil, fmt.Errorf("%s: %w", fn, err)
				}
				p.curNode = nextNode
			} else { // goto finish.
				finishNode, err := p.tree.getFinish()
				if err != nil {
					return nil, fmt.Errorf("%s: %w", fn, err)
				}
				p.curNode = finishNode
			}
		} else { // Error path.
			if p.curNode.GetNextErrorID().Valid {
				nextNode, err := p.tree.get(p.curNode.GetNextErrorID().UUID)
				if err != nil {
					return nil, fmt.Errorf("%s: %w", fn, err)
				}
				p.curNode = nextNode
			} else { // goto finish.
				finishNode, err := p.tree.getFinish()
				if err != nil {
					return nil, fmt.Errorf("%s: %w", fn, err)
				}
				p.curNode = finishNode
			}
		}
	}
}
