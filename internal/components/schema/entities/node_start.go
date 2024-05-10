package entities

import "fmt"

type NodeStart struct {
	BaseNode
}

var _ Node = (*NodeStart)(nil)

func (n *NodeStart) FitNode(node Node) error {
	const fn = "NodeStart.FitNode"

	err := n.BaseNode.FitNode(node)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}
	return nil
}
