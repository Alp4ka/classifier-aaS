package schema

type NodeClassify struct {
	BaseNode
}

func (n *NodeClassify) FromNode(node Node) error {
	return n.BaseNode.FromNode(node)
}

var _ Node = (*NodeClassify)(nil)
