package processor

import (
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/components/schema/entities"
)

type nodeFinish struct {
	*entities.NodeFinish
}

func newNodeFinish(n *entities.NodeFinish) node {
	return &nodeFinish{NodeFinish: n}
}

func (n *nodeFinish) Process(ctx context.Context, scope scope, req *nodeRequest) (*nodeResponse, error) {
	return &nodeResponse{
		Err:          nil,
		FutureAction: actionFinish,
		UserOutput:   nil,
	}, nil
}
