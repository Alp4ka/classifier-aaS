package processor

import (
	"bytes"
	"context"
	"github.com/Alp4ka/classifier-aaS/internal/schema"
	"net/http"
)

type externalRequestNodeProc struct {
	*schema.NodeExternalRequest
}

func (l *externalRequestNodeProc) process(ctx context.Context, req *request) (*response, error) {
	url := l.URL
	method := l.Method
	body := l.Body

	httpReq, _ := http.NewRequest(string(method), url, bytes.NewReader(body))
	for k, v := range l.Headers {
		httpReq.Header.Add(k, v)
	}

	res, _ := http.DefaultClient.Do(httpReq)
	defer res.Body.Close()

	return &response{}, nil
}

var _ nodeProc = (*externalRequestNodeProc)(nil)
