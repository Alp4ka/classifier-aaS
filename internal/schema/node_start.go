package schema

type NodeStart struct {
	BaseNode
}

func (n *NodeStart) FromNode(node Node) error {
	return n.BaseNode.FromNode(node)
}

var _ Node = (*NodeStart)(nil)
