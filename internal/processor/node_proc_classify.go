package processor

import (
	"context"
	"github.com/Alp4ka/classifier"
	"github.com/Alp4ka/classifier-aaS/internal/schema"
	"github.com/Alp4ka/classifier/openai"
)

type classifyNodeProc struct {
	*schema.NodeClassify
}

func (l *classifyNodeProc) process(ctx context.Context, req *request) (*response, error) {
	cls, err := openai.NewClassifier(openai.Config{APIKey: "sk-FDSCEjSJClUbk4I0RMRVykBpvNl5C4YR"})
	if err != nil {
		return nil, err
	}

	classes := make([]classifier.Class, 0, len(l.Classes))
	for _, c := range l.Classes {
		classes = append(classes, classifier.ClassStruct{Name: c.Name, Description: c.Description})
	}
	res, err := cls.Classify(ctx, classifier.Params{Classes: classes, Input: req.pipeInput.(string), AdditionalContext: l.Context})
	if err != nil {
		return nil, err
	}

	best, _ := res.Best()
	return &response{pipeOutput: best.Class().Name}, nil
}

var _ nodeProc = (*classifyNodeProc)(nil)
