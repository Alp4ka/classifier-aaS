package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Alp4ka/classifier-aaS/internal/components/processor/repository"
	timepkg "github.com/Alp4ka/classifier-aaS/pkg/time"
	"github.com/Alp4ka/mlogger"
	"github.com/Alp4ka/mlogger/field"
)

type Action = string

const (
	ActionNone    Action = "none"
	ActionListen  Action = "listen"
	ActionRespond Action = "respond"
	ActionFinish  Action = "finish"
)

type Request struct {
	UserInput string `json:"userInput"`
}

type Response struct {
	Action     Action `json:"futureAction"`
	UserOutput string `json:"userOutput"`
}

func (p *Processor) Process(ctx context.Context, req *Request) (*Response, error) {
	const fn = "Processor.Process"

	nodeReq := &nodeRequest{
		SystemConfig: p.systemConfig,
		UserInput:    req.UserInput,
		Scope:        p.scope,
	}
	parsedReq, _ := json.Marshal(nodeReq)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		ctx = field.WithContextFields(
			ctx,
			field.String("sessionID", p.sessionID.String()),
			field.String("id", p.currentNode.GetID().String()),
			field.String("type", p.currentNode.GetType()),
		)

		mlogger.L(ctx).Info("Processing node!", field.JSONEscape("nodeRequest", parsedReq))
		nodeResp, err := p.currentNode.Process(ctx, nodeReq)
		if err != nil { // Critical error while processing node.
			mlogger.L(ctx).Error("Failed to process node!", field.Error(err))
			return nil, fmt.Errorf("%s: %w", fn, err)
		}

		parsedResp, _ := json.Marshal(nodeResp)
		_, err = p.repository.CreateEvent(
			ctx,
			repository.Event{
				SessionID:    p.sessionID,
				Req:          string(parsedReq),
				Resp:         string(parsedResp),
				SchemaNodeID: p.currentNode.GetID(),
				CreatedAt:    timepkg.Now(),
				UpdatedAt:    timepkg.Now(),
			},
		)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to save event; %w", fn, err)
		}

		mlogger.L(ctx).Info("Processed node!", field.JSONEscape("nodeResponse", parsedResp))

		var nextNode node
		switch nodeResp.FutureAction {
		case nodeActionFall:
			if p.currentNode.GetNextID().Valid {
				nextNode, err = p.tree.Get(p.currentNode.GetNextID().UUID)
				if err != nil {
					return nil, fmt.Errorf("%s: cannot find next node; %w", fn, err)
				}
				break
			}
			fallthrough
		case nodeActionError:
			if p.currentNode.GetNextErrorID().Valid {
				nextNode, err = p.tree.Get(p.currentNode.GetNextID().UUID)
				if err != nil {
					return nil, fmt.Errorf("%s: cannot find next error node; %w", fn, err)
				}
			} else {
				nextNode, err = p.tree.GetFinish()
				if err != nil {
					return nil, fmt.Errorf("%s: cannot get finish since NextErrorID is invalid; %w", fn, err)
				}
			}
		case nodeActionFinish:
			return &Response{
				Action: ActionFinish,
			}, nil
		case nodeActionListen:
			return &Response{
				Action: ActionListen,
			}, nil
		case nodeActionRespond:
			return &Response{
				Action:     ActionRespond,
				UserOutput: nodeResp.UserOutput,
			}, nil
		default:
			return nil, fmt.Errorf("%s: unknown action returned from node: %s", fn, nodeResp.FutureAction)
		}

		p.currentNode = nextNode
	}
}
