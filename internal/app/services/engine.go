package services

import (
	"github.com/monacohq/go-rule-engine-example/internal/app/models"
	"github.com/monacohq/go-rule-engine-example/internal/app/rules"

	"github.com/hashicorp/go-multierror"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/engine"
)

type IRuleEngine interface {
	computeContext(fact *models.Fact) (ast.IDataContext, error)
	Execute(fact *models.Fact, expectation interface{}, features ...string) error
}

type ruleEngineImpl struct {
	engine           *engine.GruleEngine
	knowledgeLibrary rules.IKnowledgeLibrary
}

func (r *ruleEngineImpl) computeContext(fact *models.Fact) (ast.IDataContext, error) {
	var resultError error
	dataCtx := ast.NewDataContext()

	err := dataCtx.Add("User", fact.User)
	err = dataCtx.Add("SystemConfig", fact.SystemConfig)
	if err != nil {
		resultError = multierror.Append(resultError, err)
	}

	return dataCtx, resultError
}

func (r *ruleEngineImpl) Execute(fact *models.Fact, expectation interface{}, features ...string) error {
	dataCtx, err := r.computeContext(fact)
	if err != nil {
		return err
	}

	for _, feature := range r.knowledgeLibrary.GetLoadedFeatures() {
		kb := r.knowledgeLibrary.GetLibrary().NewKnowledgeBaseInstance(feature.Name, feature.Version)
		err = dataCtx.Add("Result", expectation)
		err = r.engine.Execute(dataCtx, kb)
	}

	return err
}

func New(opts ...Option) IRuleEngine {
	impl := &ruleEngineImpl{
		engine: engine.NewGruleEngine(),
	}
	for _, opt := range opts {
		opt(impl)
	}
	return impl
}
