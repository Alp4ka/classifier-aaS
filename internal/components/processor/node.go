package processor

import (
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/components/schema/entities"
)

type (
	action        = int8
	scopeVariable = string
	scope         map[scopeVariable]string
)

const (
	actionNone action = iota
	actionListen
	actionRespond
	actionFinish
	actionFall
	actionError
)

type systemConfig struct {
	ClassifierAPIKey string
}

type nodeRequest struct {
	SystemConfig *systemConfig
	UserInput    *string
}

type nodeResponse struct {
	Err          error
	FutureAction action
	UserOutput   *string
}

type node interface {
	entities.Node
	Process(ctx context.Context, scope scope, req *nodeRequest) (*nodeResponse, error)
}
