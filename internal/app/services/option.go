package services

import "github.com/monacohq/go-rule-engine-example/internal/app/rules"

type Option func(o *ruleEngineImpl)

func WithKnowledgeLibrary(knowledgeLibrary rules.IKnowledgeLibrary) Option {
	return func(o *ruleEngineImpl) {
		o.knowledgeLibrary = knowledgeLibrary
	}
}
