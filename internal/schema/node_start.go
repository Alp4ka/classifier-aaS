package schema

import "fmt"

type NodeStart struct {
	baseNode
}

func (n *NodeStart) fromNode(node Node) error {
	const fn = "NodeStart.fromNode"
	err := n.baseNode.fromNode(node)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}
	return nil
}

func (n *NodeStart) InputType() NodeDataType {
	return DataTypeNone
}

func (n *NodeStart) OutputType() NodeDataType {
	return DataTypeNone
}

var _ Node = (*NodeStart)(nil)
