package schema

import (
	"fmt"
	"github.com/google/uuid"
)

type NodeType string

func (n NodeType) Hash() string {
	return string(n)
}

type NodeDataType string

type NodeID = uuid.UUID

const (
	DataTypeString NodeDataType = "string"
	DataTypeJSON   NodeDataType = "json"
)

const (
	NodeTypeStart           NodeType = "start"
	NodeTypeFinish          NodeType = "finish"
	NodeTypeListen          NodeType = "listen"
	NodeTypeRespond         NodeType = "respond"
	NodeTypeIfElse          NodeType = "if_else"
	NodeTypeClassify        NodeType = "classify"
	NodeExternalRESTRequest NodeType = "external_request"
)

// NOTE: Does not supposed to be changed in runtime.
var _nodeTypeToInputType = map[NodeType]NodeDataType{
	NodeTypeListen:          DataTypeString,
	NodeTypeRespond:         DataTypeString,
	NodeTypeIfElse:          DataTypeString,
	NodeTypeClassify:        DataTypeJSON,
	NodeExternalRESTRequest: DataTypeString,
}

// NOTE: Does not supposed to be changed in runtime.
var _nodeTypeToOutputType = map[NodeType]NodeDataType{
	NodeTypeListen:          DataTypeString,
	NodeTypeRespond:         DataTypeString,
	NodeTypeIfElse:          DataTypeString,
	NodeTypeClassify:        DataTypeString,
	NodeExternalRESTRequest: DataTypeString,
}

type Node interface {
	mustEmbedBaseNode()
	FromNode(Node) error
	Validate() error
	InputType() NodeDataType
	OutputType() NodeDataType

	GetID() NodeID
	GetType() NodeType
	GetNextErrorID() uuid.NullUUID
	GetNextID() uuid.NullUUID
	GetData() map[string]any
}

type BaseNode struct {
	ID          NodeID         `json:"id"`
	Type        NodeType       `json:"type"`
	NextID      uuid.NullUUID  `json:"nextID"`
	NextErrorID uuid.NullUUID  `json:"nextErrorID"`
	Data        map[string]any `json:"data"`
}

func (n *BaseNode) mustEmbedBaseNode() {}

func (n *BaseNode) FromNode(node Node) error {
	n.ID = node.GetID()
	n.Type = node.GetType()
	n.NextID = node.GetNextID()
	n.NextErrorID = node.GetNextErrorID()
	n.Data = node.GetData()
	return nil
}

func (n *BaseNode) Validate() error {
	return nil
}

func (n *BaseNode) InputType() NodeDataType {
	return _nodeTypeToInputType[n.Type]
}

func (n *BaseNode) OutputType() NodeDataType {
	return _nodeTypeToOutputType[n.Type]
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

var _ Node = (*BaseNode)(nil)

func FromNode(node Node) (Node, error) {
	var ret Node
	switch node.GetType() {
	case NodeTypeStart:
		ret = &NodeStart{}
	case NodeTypeFinish:
		ret = &NodeFinish{}
	default:
		return nil, fmt.Errorf("unknown node type %s", node.GetType())
	}

	return ret, ret.FromNode(node)
}
