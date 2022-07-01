package main

import (
	"fmt"

	"github.com/monacohq/go-rule-engine-example/configs"
	"github.com/monacohq/go-rule-engine-example/internal/app/models"
	"github.com/monacohq/go-rule-engine-example/internal/app/rules"
	"github.com/monacohq/go-rule-engine-example/internal/app/services"
)

type Result struct {
	PaymentTypeEnabled bool
	Enabled            bool
}

func main() {
	configs.Load()
	cfg := configs.GetCurrentConfig()

	// initialize knowledge base from dsl files
	lib := rules.New(rules.WithFeatures(cfg.Features))
	lib.LoadRules()

	// start engine
	eng := services.New(services.WithKnowledgeLibrary(lib))

	// --------------------------
	// establish user level facts
	user := &models.User{
		Email: "sam.wang@crypto.com",
		Config: &models.UserConfig{
			RecurringBuyEnabled: true,
		},
	}
	// establish system level facts
	sysConfig := &models.SystemConfig{
		RecurringBuyEnabled:                             true,
		RecurringBuyPurchaseByCreditCardEnabled:         true,
		RecurringBuyPurchaseByCreditCardInternalTesters: []string{"sam.wang@crypto.com"},
		RecurringBuyPurchaseByStableCoinEnabled:         true,
		RecurringBuyPurchaseByStableCoinInternalTesters: []string{""},
		RecurringBuyPurchaseByFiatWalletEnabled:         true,
		RecurringBuyPurchaseByFiatWalletInternalTesters: []string{""},
	}

	fact := &models.Fact{
		User:         user,
		SystemConfig: sysConfig,
	}
	expectation := &Result{}

	err := eng.Execute(fact, expectation)

	fmt.Printf("error: %v\n", err)
	fmt.Printf("result: %+v\n", expectation)
}
