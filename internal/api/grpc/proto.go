package grpc

import (
	"fmt"

	processorcomponent "github.com/Alp4ka/classifier-aaS/internal/components/processor"
	api "github.com/Alp4ka/classifier-api"
)

func actionToProtoAction(action processorcomponent.Action) (api.Action, error) {
	switch action {
	case processorcomponent.ActionFinish:
		return api.Action_ACTION_FINISH, nil
	case processorcomponent.ActionListen:
		return api.Action_ACTION_LISTEN, nil
	case processorcomponent.ActionRespond:
		return api.Action_ACTION_RESPOND, nil
	}

	return api.Action_ACTION_UNSPECIFIED, fmt.Errorf("unknown action: %v", action)
}
