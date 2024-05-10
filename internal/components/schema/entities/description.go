package entities

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

const (
	_MaxNodesCount = 12
	_MinNodesCount = 1
)

type Description struct {
	Nodes []Node
}

func (d *Description) UnmarshalJSON(data []byte) error {
	const fn = "Description.UnmarshalJSON"

	var baseNodes []*BaseNode
	err := json.Unmarshal(data, &baseNodes)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	nodes := make([]Node, len(baseNodes))
	for i, baseNode := range baseNodes {
		node, err := GetEnrichedNode(baseNode)
		if err != nil {
			return fmt.Errorf("%s: %w", fn, err)
		}
		nodes[i] = node
	}

	d.Nodes = nodes
	return nil
}

func (d *Description) MarshalJSON() ([]byte, error) {
	const fn = "Description.MarshalJSON"

	ret, err := json.Marshal(d.Nodes)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return ret, nil
}

func (d *Description) Scan(value interface{}) error {
	valueBytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("can't scan %T as []byte", value)
	}

	return d.UnmarshalJSON(valueBytes)
}

func (d *Description) Value() (driver.Value, error) {
	return d.MarshalJSON()
}

func (d *Description) MapAndValidate() (Tree, error) {
	const fn = "Description.MapAndValidate"

	if len(d.Nodes) > _MaxNodesCount {
		return nil, fmt.Errorf("%s: too many nodes, max is %d, actual is %d", fn, _MaxNodesCount, len(d.Nodes))
	} else if len(d.Nodes) < _MinNodesCount {
		return nil, fmt.Errorf("%s: too few nodes, min is %d, actual is %d", fn, _MinNodesCount, len(d.Nodes))
	}

	// Build Tree.
	tree := make(Tree, len(d.Nodes))
	for i, node := range d.Nodes {
		if node == nil {
			return nil, fmt.Errorf("%s, node %d is nil", fn, i)
		}
		if _, ok := tree[node.GetID()]; ok {
			return nil, fmt.Errorf("%s: duplicate node ID %s", fn, node.GetID())
		}

		// Fit node.
		enrichedNode, err := GetEnrichedNode(node)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to enrich node %s; %w", fn, node.GetID(), err)
		}

		// Validate.
		err = enrichedNode.Validate()
		if err != nil {
			return nil, fmt.Errorf("%s: node %s is invalid; %w", fn, enrichedNode.GetID(), err)
		}

		tree[node.GetID()] = enrichedNode
	}

	// Cycles and finish validation.
	err := tree.HasCyclesOrUnreachableFinish()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return tree, nil
}
