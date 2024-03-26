package schema

import "fmt"

type NodeListen struct {
	baseNode
}

func (n *NodeListen) fromNode(node Node) error {
	const fn = "NodeListen.fromNode"
	err := n.baseNode.fromNode(node)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}
	return nil
}

func (n *NodeListen) InputType() NodeDataType {
	return DataTypeAny
}

func (n *NodeListen) OutputType() NodeDataType {
	return DataTypeString
}

var _ Node = (*NodeListen)(nil)
