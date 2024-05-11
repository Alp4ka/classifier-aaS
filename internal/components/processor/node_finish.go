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

func (n *nodeFinish) Process(_ context.Context, _ *nodeRequest) (*nodeResponse, error) {
	return &nodeResponse{
		Err:          nil,
		FutureAction: nodeActionFinish,
	}, nil
}
