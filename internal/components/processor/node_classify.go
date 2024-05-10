package processor

import (
	"context"
	"github.com/Alp4ka/classifier"
	"github.com/Alp4ka/classifier-aaS/internal/components/schema/entities"
	"github.com/Alp4ka/classifier/openai"
)

type nodeClassify struct {
	*entities.NodeClassify
}

func newNodeClassify(n *entities.NodeClassify) node {
	return &nodeClassify{NodeClassify: n}
}

func (n *nodeClassify) Process(ctx context.Context, scope scope, req *nodeRequest) (*nodeResponse, error) {
	cls, err := openai.NewClassifier(openai.Config{APIKey: req.SystemConfig.ClassifierAPIKey})
	if err != nil {
		return &nodeResponse{
				Err:          err,
				FutureAction: actionError,
				UserOutput:   nil,
			},
			nil
	}

	classes := make([]classifier.Class, 0, len(n.Classes))
	for _, c := range n.Classes {
		classes = append(classes, classifier.ClassStruct{Name: c.Name, Description: c.Description})
	}

	input := scope[n.InputVariable]
	res, err := cls.Classify(ctx, classifier.Params{Classes: classes, Input: input, AdditionalContext: n.Context})
	if err != nil {
		return &nodeResponse{
				Err:          err,
				FutureAction: actionError,
				UserOutput:   nil,
			},
			nil
	}

	best, _ := res.Best()
	scope[n.OutputVariable] = best.Class().Name
	return &nodeResponse{
			Err:          nil,
			FutureAction: actionFall,
			UserOutput:   nil,
		},
		nil
}
