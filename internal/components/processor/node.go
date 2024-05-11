package processor

import (
	"context"

	"github.com/Alp4ka/classifier-aaS/internal/components/schema/entities"
)

type (
	nodeAction    = string
	scopeVariable = string
	scope         map[scopeVariable]string
)

const (
	nodeActionListen  nodeAction = "listen"
	nodeActionRespond nodeAction = "respond"
	nodeActionFinish  nodeAction = "finish"
	nodeActionFall    nodeAction = "fall"
	nodeActionError   nodeAction = "error"
)

type systemConfig struct {
	ClassifierAPIKey string
}

type nodeRequest struct {
	SystemConfig *systemConfig
	UserInput    string
	Scope        scope
}

type nodeResponse struct {
	Err          error
	FutureAction nodeAction
	UserOutput   string
}

type node interface {
	entities.Node
	Process(ctx context.Context, req *nodeRequest) (*nodeResponse, error)
}
