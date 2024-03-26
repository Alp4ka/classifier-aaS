package schema

import (
	"fmt"
)

type nodeClassifyClasses []NodeClassifyClass

func (n *nodeClassifyClasses) fromAny(ms []any) error {
	const fn = "NodeClassifyClass.fromMaps"

	classes := make([]NodeClassifyClass, 0, len(ms))
	for _, m := range ms {
		converted, ok := m.(map[string]any)
		if !ok {
			return fmt.Errorf("%s: unable to map object", fn)
		}

		var c = NodeClassifyClass{}
		err := c.fromMap(converted)
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

func (n *NodeClassifyClass) fromMap(m map[string]any) error {
	const fn = "NodeClassifyClass.fromMap"

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

type NodeClassify struct {
	baseNode
	Context string              `json:"-"`
	Classes []NodeClassifyClass `json:"-"`
}

func (n *NodeClassify) fromNode(node Node) error {
	const fn = "NodeClassify.fromNode"

	err := n.baseNode.fromNode(node)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	// Context.
	ctx, ok := node.GetData()["context"]
	if !ok {
		return fmt.Errorf("%s: missing context, %w", fn, ErrMissingField)
	}
	ctxStr, ok := ctx.(string)
	if !ok {
		return fmt.Errorf("%s: wrong context field type, %w", fn, ErrFieldWrongType)
	}
	n.Context = ctxStr

	// Classes.
	classesField, ok := node.GetData()["classes"]
	if !ok {
		return fmt.Errorf("%s: missing classes, %w", fn, ErrMissingField)
	}
	classesAny, ok := classesField.([]any)
	if !ok {
		return fmt.Errorf("%s: wrong classes field type, %w", fn, ErrFieldWrongType)
	}
	classes := make(nodeClassifyClasses, 0, len(classesAny))
	err = classes.fromAny(classesAny)
	if err != nil {
		return fmt.Errorf("%s: cannot parse classes, %w", fn, err)
	}
	n.Classes = classes

	return nil
}

func (n *NodeClassify) InputType() NodeDataType {
	return DataTypeString
}

func (n *NodeClassify) OutputType() NodeDataType {
	return DataTypeString
}

var _ Node = (*NodeClassify)(nil)
