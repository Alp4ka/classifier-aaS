package processor

import (
	"bytes"
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/components/schema/entities"
	"net/http"
	"time"
)

type nodeExternalRequest struct {
	*entities.NodeExternalRequest
}

func newNodeExternalRequest(n *entities.NodeExternalRequest) node {
	return &nodeExternalRequest{NodeExternalRequest: n}
}

func (n *nodeExternalRequest) Process(ctx context.Context, _ *nodeRequest) (*nodeResponse, error) {
	const _defaultTimeout = time.Second * 5

	// TODO: Add variables support.
	var (
		url    = n.URL
		method = n.Method
		body   = n.Body
	)

	ctx, cancel := context.WithTimeout(ctx, _defaultTimeout)
	defer cancel()

	httpReq, _ := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(body))
	for k, v := range n.Headers {
		httpReq.Header.Add(k, v)
	}

	res, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return &nodeResponse{
				Err:          err,
				FutureAction: nodeActionError,
			},
			nil
	}
	defer func() { _ = res.Body.Close() }()

	return &nodeResponse{
			Err:          nil,
			FutureAction: nodeActionFall,
		},
		nil
}
