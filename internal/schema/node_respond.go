package schema

type NodeRespond struct {
	BaseNode
}

func (n *NodeRespond) FromNode(node Node) error {
	return n.BaseNode.FromNode(node)
}

var _ Node = (*NodeRespond)(nil)
