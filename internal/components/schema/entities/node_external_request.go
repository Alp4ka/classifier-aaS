package entities

import (
	"encoding/json"
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
)

type MethodType = string

const (
	MethodGET    MethodType = "GET"
	MethodPOST   MethodType = "POST"
	MethodPUT    MethodType = "PUT"
	MethodPATCH  MethodType = "PATCH"
	MethodDELETE MethodType = "DELETE"
)

var _AvailableMethods = mapset.NewThreadUnsafeSet(MethodGET)

type NodeExternalRequest struct {
	BaseNode
	URL     string            `json:"-"`
	Method  MethodType        `json:"-"`
	Headers map[string]string `json:"-"`
	Body    []byte            `json:"-"`
}

var _ Node = (*NodeExternalRequest)(nil)

func (n *NodeExternalRequest) FitNode(node Node) error {
	const fn = "NodeExternalRequest.FitNode"

	err := n.BaseNode.FitNode(node)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	// URL.
	url, ok := node.GetData()["url"]
	if !ok {
		return fmt.Errorf("%s: missing url, %w", fn, ErrMissingField)
	}
	urlStr, ok := url.(string)
	if !ok {
		return fmt.Errorf("%s: wrong url field type, %w", fn, ErrFieldWrongType)
	}
	n.URL = urlStr

	// Method.
	method, ok := node.GetData()["method"]
	if !ok {
		return fmt.Errorf("%s: missing method, %w", fn, ErrMissingField)
	}
	methodStr, ok := method.(string)
	if !ok {
		return fmt.Errorf("%s: wrong method field type, %w", fn, ErrFieldWrongType)
	}
	if ok = _AvailableMethods.Contains(methodStr); !ok {
		return fmt.Errorf("%s: unavailable method %s, %w", fn, methodStr, ErrFieldWrongType)
	}
	n.Method = MethodType(methodStr)

	// TODO: Headers.
	n.Headers = make(map[string]string)

	// Body. TODO: depth check.
	body, ok := node.GetData()["requestBody"]
	if !ok {
		return fmt.Errorf("%s: missing body, %w", fn, ErrMissingField)
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}
	n.Body = bodyBytes

	return nil
}
