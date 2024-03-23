package schema

type NodeExternalRequest struct {
	BaseNode
}

func (n *NodeExternalRequest) FromNode(node Node) error {
	return n.BaseNode.FromNode(node)
}

var _ Node = (*NodeExternalRequest)(nil)
