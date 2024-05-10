package processor

import (
	"context"
	"fmt"
	"github.com/Alp4ka/classifier-aaS/internal/components/schema/entities"
)

type nodeListen struct {
	*entities.NodeListen

	goForward bool
}

func newNodeListen(n *entities.NodeListen) node {
	return &nodeListen{NodeListen: n, goForward: false}
}

func (n *nodeListen) Process(ctx context.Context, scope scope, req *nodeRequest) (*nodeResponse, error) {
	if !n.goForward {
		n.goForward = true
		return &nodeResponse{
				Err:          nil,
				FutureAction: actionListen,
				UserOutput:   nil,
			},
			nil
	}

	if req.UserInput == nil {
		return &nodeResponse{
				Err:          fmt.Errorf("user input is empty"),
				FutureAction: actionError,
				UserOutput:   nil,
			},
			nil
	}
	scope[n.OutputVariable] = *req.UserInput

	return &nodeResponse{
			Err:          nil,
			FutureAction: actionNone,
			UserOutput:   nil,
		},
		nil
}
