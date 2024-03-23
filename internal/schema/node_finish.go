package schema

import "fmt"

type NodeFinish struct {
	BaseNode
}

func (n *NodeFinish) Validate() error {
	if err := n.BaseNode.Validate(); err != nil {
		return err
	}
	if n.BaseNode.NextID.Valid {
		return fmt.Errorf("node %s is not finished since it has a next node", n.BaseNode.ID)
	}
	if n.BaseNode.NextErrorID.Valid {
		return fmt.Errorf("node %s is not finished since it has error path", n.BaseNode.ID)
	}

	return nil
}

func (n *NodeFinish) FromNode(node Node) error {
	return n.BaseNode.FromNode(node)
}

var _ Node = (*NodeFinish)(nil)
