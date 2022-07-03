package main

import (
	"log"

	"github.com/monacohq/go-rule-engine-example/configs"
	"github.com/monacohq/go-rule-engine-example/internal/app/models"
	"github.com/monacohq/go-rule-engine-example/internal/app/rules"
	"github.com/monacohq/go-rule-engine-example/internal/app/services"
)

func main() {
	configs.Load()
	cfg := configs.GetCurrentConfig()

	// preload system configuration from the database / redis
	systemConfig := &models.SystemConfig{Value: make(map[string]interface{})}
	systemConfig.Load()

	// initialize knowledge base from dsl files
	lib := rules.New(rules.WithFeatures(cfg.Features))
	lib.LoadRules()

	// start engine for validation
	eng := services.New(
		services.WithKnowledgeLibrary(lib),
		services.WithSystemConfigs(systemConfig),
	)

	// establish user level facts, usually only bring in UUID, but if sensitive data is needed,
	// it should be passed in from request owner.
	fact := &models.Fact{
		User: &models.User{
			Email: "sam.wang@crypto.com",
			Config: &models.UserConfig{
				RecurringBuyEnabled: true,
			},
		},
	}
	// we don't know what's inside DSL, using map with interface{} to store final results
	expectation := &models.Result{Value: map[string]interface{}{}}
	validateFeatures := []string{"recurring_buy"}
	err := eng.Execute(fact, expectation, validateFeatures...)

	log.Printf("error: %v\n", err)
	log.Printf("result: %s\n", expectation.ToJson())
}
