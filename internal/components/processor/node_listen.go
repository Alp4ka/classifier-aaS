package processor

import (
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/components/schema/entities"
)

type nodeListen struct {
	*entities.NodeListen

	goForward bool
}

func newNodeListen(n *entities.NodeListen) node {
	return &nodeListen{NodeListen: n, goForward: false}
}

func (n *nodeListen) Process(_ context.Context, req *nodeRequest) (*nodeResponse, error) {
	if !n.goForward {
		n.goForward = true
		return &nodeResponse{
				Err:          nil,
				FutureAction: nodeActionListen,
			},
			nil
	}
	req.Scope[n.OutputVariable] = req.UserInput

	return &nodeResponse{
			Err:          nil,
			FutureAction: nodeActionNone,
		},
		nil
}
