package processor

import (
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/components/schema/entities"
)

type nodeRespond struct {
	*entities.NodeRespond

	goForward bool
}

func newNodeRespond(n *entities.NodeRespond) node {
	return &nodeRespond{NodeRespond: n, goForward: false}
}

func (n *nodeRespond) Process(_ context.Context, req *nodeRequest) (*nodeResponse, error) {
	if !n.goForward {
		n.goForward = true

		outputText := FormatString(n.Response, req.Scope)
		return &nodeResponse{
				Err:          nil,
				FutureAction: nodeActionRespond,
				UserOutput:   outputText,
			},
			nil
	}

	return &nodeResponse{
			Err:          nil,
			FutureAction: nodeActionFall,
		},
		nil
}
