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

func (n *nodeRespond) Process(ctx context.Context, scope scope, req *nodeRequest) (*nodeResponse, error) {
	if !n.goForward {
		n.goForward = true

		outputText := FormatString(n.Response, scope)
		return &nodeResponse{
				Err:          nil,
				FutureAction: actionRespond,
				UserOutput:   &outputText,
			},
			nil
	}

	return &nodeResponse{
			Err:          nil,
			FutureAction: actionFall,
			UserOutput:   nil,
		},
		nil
}
