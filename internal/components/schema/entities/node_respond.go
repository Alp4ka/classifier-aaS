package entities

import (
	"fmt"
)

type NodeRespond struct {
	BaseNode
	Response string `json:"-"`
}

var _ Node = (*NodeRespond)(nil)

func (n *NodeRespond) FitNode(node Node) error {
	const fn = "NodeRespond.FitNode"

	err := n.BaseNode.FitNode(node)
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
