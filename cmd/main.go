package main

import (
	"log"
	"time"

	"github.com/monacohq/go-rule-engine-example/configs"
	"github.com/monacohq/go-rule-engine-example/internal/app/models"
	"github.com/monacohq/go-rule-engine-example/internal/app/rules"
	"github.com/monacohq/go-rule-engine-example/internal/app/services"
)

func main() {
	var err error
	time.Local, err = time.LoadLocation("Asia/Hong_Kong")
	if err != nil {
		log.Printf("error loading location: %v\n", err)
	}

	configs.Load()
	cfg := configs.GetCurrentConfig()

	// preload system configuration from the database / redis
	systemConfig := &models.SystemConfig{
		Value:              make(map[string]interface{}),
		EntityRestrictions: map[string][]models.ForbiddenFeature{},
	}
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
			UUID:               "123e4567-e89b-12d3-a456-426614174000",
			Email:              "sam.wang@crypto.com",
			EntityId:           "canada",
			ResidentialAddress: "sample_address",
			Config: &models.UserConfig{
				RecurringBuyEnabled:  true,
				HasCryptoFiatAccount: false,
			},
			KycDocument: &models.KycDocument{
				AppliedAt: "2020-05-27 15:00:00",
			},
		},
	}
	// we don't know what's inside DSL, using map with interface{} to store final results
	expectation := &models.Result{Value: map[string]interface{}{}}
	validateFeatures := []string{"recurring_buy", "entity_restriction"}

	start := time.Now()
	err = eng.Execute(fact, expectation, validateFeatures...)
	elapsed := time.Since(start)
	log.Printf("Execute took %s", elapsed)
	log.Printf("result: %s, err: %v\n", expectation.ToJson(), err)
}
