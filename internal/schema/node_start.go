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
	return DataTypeString
}

func (n *NodeStart) OutputType() NodeDataType {
	return DataTypeString
}

var _ Node = (*NodeStart)(nil)
