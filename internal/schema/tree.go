package schema

import "fmt"

type Tree map[NodeID]Node

func (t Tree) GetStart() (Node, error) {
	for _, node := range t {
		if node.GetType() == NodeTypeStart {
			return node, nil
		}
	}
	return nil, fmt.Errorf("start node not found")
}
