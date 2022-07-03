package rules

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/monacohq/go-rule-engine-example/configs"
)

const (
	UserKey         = "User"
	SystemConfigKey = "SystemConfig"
	ResultKey       = "Result"
)

type IKnowledgeLibrary interface {
	GetLoadedFeatures() map[string]*configs.Feature
	GetLibrary() *ast.KnowledgeLibrary
	LoadRules() error
}

type knowledgeLibraryImpl struct {
	library  *ast.KnowledgeLibrary
	features []*configs.Feature
}

func (k *knowledgeLibraryImpl) GetLoadedFeatures() map[string]*configs.Feature {
	results := make(map[string]*configs.Feature)
	for _, f := range k.features {
		results[f.Name] = f
	}
	return results
}

func (k *knowledgeLibraryImpl) GetLibrary() *ast.KnowledgeLibrary {
	return k.library
}

func (k *knowledgeLibraryImpl) LoadRules() error {
	var errResult error
	k.library = ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(k.library)

	//TODO: Readjust the loaded features
	for _, f := range k.features {
		loc := fmt.Sprintf("./configs/dsl/%s", f.DSL)
		err := rb.BuildRuleFromResource(f.Name, f.Version, pkg.NewFileResource(loc))
		errResult = multierror.Append(errResult, err)
	}

	return errResult
}

func New(opts ...Option) IKnowledgeLibrary {
	impl := &knowledgeLibraryImpl{}
	for _, opt := range opts {
		opt(impl)
	}
	return impl
}
