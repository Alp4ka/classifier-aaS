package schema

type NodeRespond struct {
	BaseNode
}

func (n *NodeRespond) FromNode(node Node) error {
	//const fn = "NodeRespond.FromNode"
	//response, ok := node.GetData()["response"]
	//if !ok {
	//	return fmt.Errorf("%s: missing response, %w", fn, ErrMissingField)
	//}
	return n.BaseNode.FromNode(node)
}

var _ Node = (*NodeRespond)(nil)
