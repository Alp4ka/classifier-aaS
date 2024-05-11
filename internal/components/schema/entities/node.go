package entities

import (
	"fmt"

	"github.com/google/uuid"
)

type (
	NodeType = string
	NodeID   = uuid.UUID
)

const (
	NodeTypeStart           NodeType = "start"
	NodeTypeFinish          NodeType = "finish"
	NodeTypeListen          NodeType = "listen"
	NodeTypeRespond         NodeType = "respond"
	NodeTypeClassify        NodeType = "classify"
	NodeTypeExternalRequest NodeType = "external_request"
)

func GetEnrichedNode(node Node) (Node, error) {
	const fn = "GetEnrichedNode"

	var ret Node
	switch node.GetType() {
	case NodeTypeStart:
		ret = new(NodeStart)
	case NodeTypeFinish:
		ret = new(NodeFinish)
	case NodeTypeListen:
		ret = new(NodeListen)
	case NodeTypeClassify:
		ret = new(NodeClassify)
	case NodeTypeRespond:
		ret = new(NodeRespond)
	case NodeTypeExternalRequest:
		ret = new(NodeExternalRequest)
	default:
		return nil, fmt.Errorf("%s: unknown node type %s", fn, node.GetType())
	}

	err := ret.FitNode(node)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to fit node; %w", fn, err)
	}

	return ret, nil
}

type Node interface {
	GetID() NodeID
	GetType() NodeType
	GetNextErrorID() uuid.NullUUID
	GetNextID() uuid.NullUUID
	GetData() map[string]any
	GetGridData() map[string]any
	Validate() error
	FitNode(Node) error
	IsStart() bool
	IsFinish() bool
}

type BaseNode struct {
	ID          NodeID         `json:"id"`
	Type        NodeType       `json:"type"`
	NextID      uuid.NullUUID  `json:"nextID"`
	NextErrorID uuid.NullUUID  `json:"nextErrorID"`
	Data        map[string]any `json:"data"`
	GridData    map[string]any `json:"gridData"`
}

var _ Node = (*BaseNode)(nil)

func (n *BaseNode) Validate() error {
	return nil
}

func (n *BaseNode) FitNode(node Node) error {
	n.ID = node.GetID()
	n.Type = node.GetType()
	n.NextID = node.GetNextID()
	n.NextErrorID = node.GetNextErrorID()
	n.Data = node.GetData()
	n.GridData = node.GetGridData()
	return nil
}

func (n *BaseNode) GetID() NodeID {
	return n.ID
}

func (n *BaseNode) GetType() NodeType {
	return n.Type
}

func (n *BaseNode) GetNextErrorID() uuid.NullUUID {
	return n.NextErrorID
}

func (n *BaseNode) GetNextID() uuid.NullUUID {
	return n.NextID
}

func (n *BaseNode) GetData() map[string]any {
	return n.Data
}

func (n *BaseNode) GetGridData() map[string]any {
	return n.GridData
}

func (n *BaseNode) IsStart() bool {
	return n.GetType() == NodeTypeStart
}

func (n *BaseNode) IsFinish() bool {
	return n.GetType() == NodeTypeFinish
}
