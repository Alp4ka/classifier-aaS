package schema

import "fmt"

type NodeRespond struct {
	baseNode
	Response string `json:"-"`
}

func (n *NodeRespond) fromNode(node Node) error {
	const fn = "NodeRespond.fromNode"

	err := n.baseNode.fromNode(node)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	response, ok := node.GetData()["response"]
	if !ok {
		return fmt.Errorf("%s: missing response, %w", fn, ErrMissingField)
	}

	responseStr, ok := response.(string)
	if !ok {
		return fmt.Errorf("%s: wrong response field type, %w", fn, ErrFieldWrongType)
	}

	n.Response = responseStr
	return nil
}

func (n *NodeRespond) InputType() NodeDataType {
	return DataTypeString
}

func (n *NodeRespond) OutputType() NodeDataType {
	return DataTypeString
}

var _ Node = (*NodeRespond)(nil)
