package processor

import (
	"fmt"

	"github.com/Alp4ka/classifier-aaS/internal/components/schema/entities"
)

type tree map[entities.NodeID]node

func newTree(schemaTree entities.Tree) (tree, error) {
	const fn = "newTree"

	ret := make(map[entities.NodeID]node, len(schemaTree))
	for id, n := range schemaTree {
		var temp node

		switch n.GetType() {
		case entities.NodeTypeClassify:
			temp = newNodeClassify(n.(*entities.NodeClassify))
		case entities.NodeTypeExternalRequest:
			temp = newNodeExternalRequest(n.(*entities.NodeExternalRequest))
		case entities.NodeTypeFinish:
			temp = newNodeFinish(n.(*entities.NodeFinish))
		case entities.NodeTypeListen:
			temp = newNodeListen(n.(*entities.NodeListen))
		case entities.NodeTypeRespond:
			temp = newNodeRespond(n.(*entities.NodeRespond))
		case entities.NodeTypeStart:
			temp = newNodeStart(n.(*entities.NodeStart))
		default:
			return nil, fmt.Errorf("%s: unprocessable node type %s", fn, n.GetType())
		}

		ret[id] = temp
	}

	return ret, nil
}

func (t *tree) GetStart() (node, error) {
	const fn = "tree.GetStart"

	for _, n := range *t {
		if n.GetType() == entities.NodeTypeStart {
			return n, nil
		}
	}

	return nil, fmt.Errorf("%s: start node not found", fn)
}

func (t *tree) GetFinish() (node, error) {
	const fn = "tree.GetFinish"

	for _, n := range *t {
		if n.GetType() == entities.NodeTypeFinish {
			return n, nil
		}
	}

	return nil, fmt.Errorf("%s: finish node not found", fn)
}

func (t *tree) Get(id entities.NodeID) (node, error) {
	const fn = "tree.Get"

	n, ok := (*t)[id]
	if !ok {
		return nil, fmt.Errorf("%s: node %s not found", fn, n.GetID())
	}

	return n, nil
}
