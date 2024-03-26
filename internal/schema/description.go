package schema

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

const MaxNodesCount = 12

type Tree map[NodeID]Node

func (t Tree) GetStart() (Node, error) {
	for _, node := range t {
		if node.GetType() == NodeTypeStart {
			return node, nil
		}
	}
	return nil, fmt.Errorf("start node not found")
}

type Description struct {
	nodes []*baseNode
}

func (d *Description) UnmarshalJSON(data []byte) error {
	var nodes []*baseNode
	err := json.Unmarshal(data, &nodes)
	if err != nil {
		return err
	}

	(*d).nodes = nodes
	return nil
}

func (d *Description) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.nodes)
}

func (d *Description) Scan(value interface{}) error {
	valueBytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("can't scan %T as Description", value)
	}

	return d.UnmarshalJSON(valueBytes)
}

func (d Description) Value() (driver.Value, error) {
	return d.MarshalJSON()
}

func (d *Description) MapAndValidate() (Tree, error) {
	// Validate.
	if len(d.nodes) > MaxNodesCount {
		return nil, fmt.Errorf("too many nodes, max is %d, actual is %d", MaxNodesCount, len(d.nodes))
	}

	// TODO: hasFinishPath  bool. If it has any path.
	// TODO: graph hasCycles?

	// Build mapping.
	var (
		cntStartNodes  int
		cntFinishNodes int
	)
	mapping := make(Tree)
	for _, node := range d.nodes {
		if node == nil {
			return nil, fmt.Errorf("node is nil")
		}
		if _, ok := mapping[node.ID]; ok {
			return nil, fmt.Errorf("duplicate node ID %s", node.ID)
		}
		mapping[node.ID] = node
	}

	// Check for fake paths.
	for k := range mapping {
		// Check if there are to many start/finish nodes.
		if mapping[k].GetType() == NodeTypeStart {
			// TODO start cannot be a next node id for any.
			cntStartNodes++
			if cntStartNodes > 1 {
				return nil, fmt.Errorf("too many start nodes")
			}
		}
		if mapping[k].GetType() == NodeTypeFinish {
			// TODO start cannot have next node.
			cntFinishNodes++
			if cntFinishNodes > 1 {
				return nil, fmt.Errorf("too many finish nodes")
			}
		}

		if mapping[k].GetNextErrorID().Valid {
			if mapping[k].GetNextErrorID().UUID == mapping[k].GetID() {
				return nil, fmt.Errorf("node %s has itself as next error node", mapping[k].GetID())
			}
			if _, ok := mapping[mapping[k].GetNextErrorID().UUID]; !ok {
				return nil, fmt.Errorf("node %s has next error node %s which is not defined", mapping[k].GetID(), mapping[k].GetNextErrorID().UUID)
			}
		}

		if mapping[k].GetNextID().Valid {
			if mapping[k].GetNextID().UUID == mapping[k].GetID() {
				return nil, fmt.Errorf("node %s has itself as next error node", mapping[k].GetID())
			}
			if _, ok := mapping[mapping[k].GetNextID().UUID]; !ok {
				return nil, fmt.Errorf("node %s has next node %s which is not defined", mapping[k].GetID(), mapping[k].GetNextID().UUID)
			}
		}

		// Fit node.
		tmpNode, err := FromNode(mapping[k])
		if err != nil {
			return nil, fmt.Errorf("node %s %w", mapping[k].GetID(), err)
		}

		// Validate.
		err = tmpNode.Validate()
		if err != nil {
			return nil, fmt.Errorf("node %s failed to validate %w", mapping[k].GetID(), err)
		}

		mapping[k] = tmpNode
	}

	for _, cur := range mapping {
		if cur.GetNextID().Valid {
			if cur.OutputType() != mapping[cur.GetNextID().UUID].InputType() {
				return nil, fmt.Errorf("node %s has next node %s which has incompatible input/output types", cur.GetID(), cur.GetNextID().UUID)
			}
		}

		if cur.GetNextErrorID().Valid {
			if cur.OutputType() != mapping[cur.GetNextErrorID().UUID].InputType() {
				return nil, fmt.Errorf("node %s has next node %s which has incompatible input/output types", cur.GetID(), cur.GetNextID().UUID)
			}
		}
	}
	// Also a problem if there are to many start/finish nodes.
	if cntStartNodes != 1 {
		return nil, fmt.Errorf("no start node")
	}
	if cntFinishNodes != 1 {
		return nil, fmt.Errorf("no finish node")
	}

	return mapping, nil
}

func (d *Description) Map() (Tree, error) {
	mapping := make(Tree)
	for _, node := range d.nodes {
		tmpNode, err := FromNode(node)
		if err != nil {
			return nil, fmt.Errorf("node %s %w", node.GetID(), err)
		}

		mapping[node.ID] = tmpNode
	}

	return mapping, nil
}
