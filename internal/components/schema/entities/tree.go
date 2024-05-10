package entities

import (
	"fmt"
)

type Tree map[NodeID]Node

func (t Tree) GetStart() (Node, error) {
	const fn = "Tree.GetStart"

	var ret Node

	for _, node := range t {
		if node.IsStart() {
			if ret == nil {
				ret = node
			} else {
				return nil, fmt.Errorf("%s: too many start nodes", fn)
			}
		}
	}

	if ret == nil {
		return nil, fmt.Errorf("%s: start node not found", fn)
	}

	return ret, nil
}

func (t Tree) GetFinish() (Node, error) {
	const fn = "Tree.GetFinish"

	var ret Node

	for _, node := range t {
		if node.IsFinish() {
			if ret == nil {
				ret = node
			} else {
				return nil, fmt.Errorf("%s: too many finish nodes", fn)
			}
		}
	}

	if ret == nil {
		return nil, fmt.Errorf("%s: finish node not found", fn)
	}

	return ret, nil
}

func (t Tree) HasCyclesOrUnreachableFinish() error {
	const fn = "Tree.HasCyclesOrUnreachableFinish"

	visitedStack := make(map[NodeID]bool)
	startNode, err := t.GetStart()
	if err != nil {
		return fmt.Errorf("%s: can't get start node; %w", fn, err)
	}
	finishNode, err := t.GetFinish()
	if err != nil {
		return fmt.Errorf("%s: can't get finish node; %w", fn, err)
	}

	err = hasCycles(t, startNode, finishNode, visitedStack)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	if visitedFinish := visitedStack[finishNode.GetID()]; !visitedFinish {
		return fmt.Errorf("%s: can't reach finish node", fn)
	}

	return nil
}

func hasCycles(tree Tree, currentNode, finishNode Node, visitedStack map[NodeID]bool) error {
	if visitedStack[currentNode.GetID()] {
		return fmt.Errorf("cycle detected on node: %s", currentNode.GetID())
	}
	visitedStack[currentNode.GetID()] = true

	if currentNode == finishNode {
		return nil
	}

	nextID := currentNode.GetNextID()
	if !nextID.Valid {
		return nil
	}
	// Next node not found.
	nextNode, ok := tree[nextID.UUID]
	if !ok {
		return fmt.Errorf("node %s relates to node %s that does not exist", currentNode.GetID(), nextID.UUID)
	}
	// Points to itself.
	if nextNode.GetID() == currentNode.GetID() {
		return fmt.Errorf("node %s relates to itself", currentNode.GetID())
	}

	return hasCycles(tree, nextNode, finishNode, visitedStack)
}
