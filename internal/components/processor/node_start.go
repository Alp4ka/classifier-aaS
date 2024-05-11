package processor

import (
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/components/schema/entities"
)

type nodeStart struct {
	*entities.NodeStart
}

func newNodeStart(n *entities.NodeStart) node {
	return &nodeStart{NodeStart: n}
}

func (n *nodeStart) Process(_ context.Context, _ *nodeRequest) (*nodeResponse, error) {
	return &nodeResponse{
		Err:          nil,
		FutureAction: nodeActionFall,
	}, nil
}
