package processor

import (
	"context"
	"fmt"
	"github.com/Alp4ka/mlogger"
	"github.com/Alp4ka/mlogger/field"
)


type Request struct {

}

type Response struct {

}

func (p *Processor) Handle(ctx context.Context, req *Request) (*Response, error) {
	const fn = "Processor.Handle"

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		mlogger.L(ctx).Info(
			"Processing node",
			field.String("id", p.currentNode.GetID().String()),
			field.String("type", p.currentNode.GetType()),
		)

		ret, err := p.currentNode.Process(ctx, p.scope, )
		if err != nil {
			return nil, fmt.Errorf("%s: %w", fn, err)
		}

		// Return conditions.
		if ret. {
			return &nodeResponse{
				End: true,
			}, nil
		} else if !nRet.fall() {
			return &nodeResponse{
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
