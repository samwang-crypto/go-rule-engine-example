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
	Execute(fact *models.Fact, expectation *models.Result, features ...string) error
}

type ruleEngineImpl struct {
	engine           *engine.GruleEngine
	systemConfigs    *models.SystemConfig
	knowledgeLibrary rules.IKnowledgeLibrary
}

func (r *ruleEngineImpl) computeContext(fact *models.Fact) (ast.IDataContext, error) {
	var errResult error
	dataCtx := ast.NewDataContext()

	err := dataCtx.Add(rules.UserKey, fact.User)
	err = dataCtx.Add(rules.SystemConfigKey, r.systemConfigs)
	if err != nil {
		errResult = multierror.Append(errResult, err)
	}

	return dataCtx, errResult
}

func (r *ruleEngineImpl) Execute(fact *models.Fact, expectation *models.Result, features ...string) error {
	var errResult error
	dataCtx, err := r.computeContext(fact)
	if err != nil {
		errResult = multierror.Append(errResult, err)
		return err
	}

	loadedFeatures := r.knowledgeLibrary.GetLoadedFeatures()
	for _, featureName := range features {
		if cnf, ok := loadedFeatures[featureName]; ok {
			kb := r.knowledgeLibrary.GetLibrary().NewKnowledgeBaseInstance(cnf.Name, cnf.Version)
			err = dataCtx.Add(rules.ResultKey, expectation)
			err = r.engine.Execute(dataCtx, kb)
			if err != nil {
				errResult = multierror.Append(errResult, err)
			}
		}
	}

	return errResult
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
