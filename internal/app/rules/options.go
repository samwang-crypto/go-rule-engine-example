package rules

import "github.com/monacohq/go-rule-engine-example/configs"

type Option func(o *knowledgeLibraryImpl)

func WithFeatures(features []*configs.Feature) Option {
	return func(o *knowledgeLibraryImpl) {
		o.features = features
	}
}
