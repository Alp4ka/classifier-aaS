package entities

import (
	"fmt"
)

type NodeClassify struct {
	BaseNode
	InputVariable  string              `json:"-"`
	OutputVariable string              `json:"-"`
	Context        string              `json:"-"`
	Classes        []NodeClassifyClass `json:"-"`
}

var _ Node = (*NodeClassify)(nil)

func (n *NodeClassify) FitNode(node Node) error {
	const fn = "NodeClassify.FitNode"

	err := n.BaseNode.FitNode(node)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	// Context.
	contextFieldAny, ok := node.GetData()["context"]
	if !ok {
		return fmt.Errorf("%s: missing context, %w", fn, ErrMissingField)
	}
	contextString, ok := contextFieldAny.(string)
	if !ok {
		return fmt.Errorf("%s: wrong context field type, %w", fn, ErrFieldWrongType)
	}

	// Classes.
	classesField, ok := node.GetData()["classes"]
	if !ok {
		return fmt.Errorf("%s: missing classes, %w", fn, ErrMissingField)
	}
	classesAny, ok := classesField.([]any)
	if !ok {
		return fmt.Errorf("%s: wrong classes field type, %w", fn, ErrFieldWrongType)
	}
	classes := make(NodeClassifyClasses, 0, len(classesAny))
	err = classes.FromSliceAny(classesAny)
	if err != nil {
		return fmt.Errorf("%s: cannot parse classes, %w", fn, err)
	}

	// Input variable.
	inputVariableFieldAny, ok := node.GetData()["inputVariable"]
	if !ok {
		return fmt.Errorf("%s: missing input variable value, %w", fn, ErrMissingField)
	}
	inputVariableString, ok := inputVariableFieldAny.(string)
	if !ok {
		return fmt.Errorf("%s: wrong input variable field type, %w", fn, ErrFieldWrongType)
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

	n.Context = contextString
	n.Classes = classes
	n.InputVariable = inputVariableString
	n.OutputVariable = outputVariableString
	return nil
}

type NodeClassifyClasses []NodeClassifyClass

func (n *NodeClassifyClasses) FromSliceAny(ms []any) error {
	const fn = "NodeClassifyClass.fromMaps"

	classes := make([]NodeClassifyClass, 0, len(ms))
	for _, m := range ms {
		converted, ok := m.(map[string]any)
		if !ok {
			return fmt.Errorf("%s: unable to map object", fn)
		}

		var c = NodeClassifyClass{}
		err := c.FromMapAny(converted)
		if err != nil {
			return fmt.Errorf("%s: %w", fn, err)
		}
		classes = append(classes, c)
	}

	*n = classes
	return nil
}

type NodeClassifyClass struct {
	Name        string
	Description string
}

func (n *NodeClassifyClass) FromMapAny(m map[string]any) error {
	const fn = "NodeClassifyClass.FromMapAny"

	// Name.
	name, ok := m["name"]
	if !ok {
		return fmt.Errorf("%s: missing name, %w", fn, ErrMissingField)
	}
	nameStr, ok := name.(string)
	if !ok {
		return fmt.Errorf("%s: wrong name field type, %w", fn, ErrFieldWrongType)
	}

	// Description.
	description, ok := m["value"]
	if !ok {
		return fmt.Errorf("%s: missing description, %w", fn, ErrMissingField)
	}
	descriptionStr, ok := description.(string)
	if !ok {
		return fmt.Errorf("%s: wrong description field type, %w", fn, ErrFieldWrongType)
	}

	n.Name = nameStr
	n.Description = descriptionStr
	return nil
}
