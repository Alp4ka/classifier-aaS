package schema

type NodeListen struct {
	BaseNode
}

func (n *NodeListen) FromNode(node Node) error {
	return n.BaseNode.FromNode(node)
}

var _ Node = (*NodeListen)(nil)
