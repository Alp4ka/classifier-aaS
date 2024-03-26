package schema

import "fmt"

type NodeFinish struct {
	baseNode
}

func (n *NodeFinish) Validate() error {
	if err := n.baseNode.Validate(); err != nil {
		return err
	}
	if n.baseNode.NextID.Valid {
		return fmt.Errorf("node %s is not finished since it has a next node", n.baseNode.ID)
	}
	if n.baseNode.NextErrorID.Valid {
		return fmt.Errorf("node %s is not finished since it has error path", n.baseNode.ID)
	}

	return nil
}

func (n *NodeFinish) fromNode(node Node) error {
	const fn = "NodeFinish.fromNode"
	err := n.baseNode.fromNode(node)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}
	return nil
}

func (n *NodeFinish) InputType() NodeDataType {
	return DataTypeAny
}

func (n *NodeFinish) OutputType() NodeDataType {
	return DataTypeNone
}

var _ Node = (*NodeFinish)(nil)
