package entities

import "fmt"

type NodeListen struct {
	BaseNode
	OutputVariable string `json:"-"`
}

var _ Node = (*NodeListen)(nil)

func (n *NodeListen) FitNode(node Node) error {
	const fn = "NodeListen.FitNode"

	err := n.BaseNode.FitNode(node)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	// Output variable.
	outputVariableFieldAny, ok := node.GetData()["outputVariable"]
	if !ok {
		return fmt.Errorf("%s: missing output variable value, %w", fn, ErrMissingField)
	}
	outputVariableString, ok := outputVariableFieldAny.(string)
	if !ok {
		return fmt.Errorf("%s: wrong output variable field type, %w", fn, ErrFieldWrongType)
	}

	n.OutputVariable = outputVariableString
	return nil
}
