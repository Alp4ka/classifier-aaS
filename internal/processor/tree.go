package processor

import (
	"fmt"
	"github.com/Alp4ka/classifier-aaS/internal/schema"
)

type tree map[schema.NodeID]nodeProc

func (t *tree) fromSchemaTree(schemaTree schema.Tree) error {
	*t = make(map[schema.NodeID]nodeProc)

	const fn = "tree.fromSchemaTree"
	for id, node := range schemaTree {
		var n nodeProc

		switch node.GetType() {
		case schema.NodeTypeClassify:
			n = &classifyNodeProc{NodeClassify: node.(*schema.NodeClassify)}
		case schema.NodeTypeExternalRequest:
			n = &externalRequestNodeProc{NodeExternalRequest: node.(*schema.NodeExternalRequest)}
		case schema.NodeTypeFinish:
			n = &finishNodeProc{NodeFinish: node.(*schema.NodeFinish)}
		case schema.NodeTypeListen:
			n = &listenNodeProc{NodeListen: node.(*schema.NodeListen)}
		case schema.NodeTypeRespond:
			n = &respondNodeProc{NodeRespond: node.(*schema.NodeRespond)}
		case schema.NodeTypeStart:
			n = &startNodeProc{NodeStart: node.(*schema.NodeStart)}
		default:
			return fmt.Errorf("%s: unprocessable node type %s", fn, node.GetType())
		}

		(*t)[id] = n
	}

	return nil
}

func (t *tree) getStart() (nodeProc, error) {
	for _, n := range *t {
		if n.GetType() == schema.NodeTypeStart {
			return n, nil
		}
	}
	return nil, fmt.Errorf("start node not found")
}

func (t *tree) getFinish() (nodeProc, error) {
	for _, n := range *t {
		if n.GetType() == schema.NodeTypeFinish {
			return n, nil
		}
	}
	return nil, fmt.Errorf("finish node not found")
}

func (t *tree) get(id schema.NodeID) (nodeProc, error) {
	n, ok := (*t)[id]
	if !ok {
		return nil, fmt.Errorf("node %s not found", n.GetID())
	}

	return n, nil
}
