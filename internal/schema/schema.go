package schema

import (
	"encoding/json"
	"fmt"
)

const MaxNodesCount = 100

type Description struct {
	nodes []*BaseNode
}

func (d *Description) UnmarshalJSON(data []byte) error {
	var nodes []*BaseNode
	err := json.Unmarshal(data, &nodes)
	if err != nil {
		return fmt.Errorf("failed to unmarshal description: %w", err)
	}

	(*d).nodes = nodes
	return nil
}

func (d *Description) MapAndValidate() (map[NodeID]Node, error) {
	// Validate.
	if len(d.nodes) > MaxNodesCount {
		return nil, fmt.Errorf("too many nodes, max is %d, actual is %d", MaxNodesCount, len(d.nodes))
	}

	// TODO: hasFinishPath  bool. If it has any path.

	// Build mapping.
	var (
		cntStartNodes  int
		cntFinishNodes int
	)
	mapping := make(map[NodeID]Node)
	for _, node := range d.nodes {
		if node == nil {
			return nil, fmt.Errorf("node is nil")
		}
		if node.Type == NodeTypeStart {
			cntStartNodes++
			if cntStartNodes > 1 {
				return nil, fmt.Errorf("too many start nodes")
			}
		}
		if node.Type == NodeTypeFinish {
			cntFinishNodes++
			if cntFinishNodes > 1 {
				return nil, fmt.Errorf("too many finish nodes")
			}
		}
		if _, ok := mapping[node.ID]; ok {
			return nil, fmt.Errorf("duplicate node ID %s", node.ID)
		}
		mapping[node.ID] = node
	}
	if cntStartNodes != 1 {
		return nil, fmt.Errorf("no start node")
	}
	if cntFinishNodes != 1 {
		return nil, fmt.Errorf("no finish node")
	}

	// Check for fake paths.
	for k, _ := range mapping {
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

		tmpNode, err := FromNode(mapping[k])
		if err != nil {
			return nil, fmt.Errorf("node %s %w", mapping[k].GetID(), err)
		}

		err = tmpNode.Validate()
		if err != nil {
			return nil, fmt.Errorf("node %s failed to validate %w", mapping[k].GetID(), err)
		}

		mapping[k] = tmpNode
	}

	return mapping, nil
}

//type Processor struct {
//	desc Description
//}
//
//func NewProcessor(desc Description) *Processor {
//	return &Processor{desc: desc}
//}
//
//func (p *Processor) Process(schema *Schema) *Schema {
//
//}
