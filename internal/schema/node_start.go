package schema

type NodeStart struct {
	BaseNode
}

var _ Node = (*NodeStart)(nil)
