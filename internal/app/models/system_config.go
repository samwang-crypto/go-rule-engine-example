package models

import "golang.org/x/exp/slices"

type SystemConfig struct {
	Value              map[string]interface{}
	EntityRestrictions *EntityRestrictions
}

// this Load function is used to load the system config from the database / redis
func (s *SystemConfig) Load() error {
	//TODO: replace with acutal redis / database code
	// here just to simulate loading from the database
	s.Value["RecurringBuyEnabled"] = false
	s.Value["RecurringBuyPurchaseByCreditCardEnabled"] = true
	s.Value["RecurringBuyPurchaseByCreditCardInternalTesters"] = []string{}
	s.Value["RecurringBuyPurchaseByStableCoinEnabled"] = false
	s.Value["RecurringBuyPurchaseByStableCoinInternalTesters"] = []string{"sam.wang@crypto.com"}
	s.Value["RecurringBuyPurchaseByFiatWalletEnabled"] = false
	s.Value["RecurringBuyPurchaseByFiatWalletInternalTesters"] = []string{}

	s.EntityRestrictions.Value["singapore"] = []ForbiddenFeature{
		ForbiddenFeature{Feature: "sepa_deposit"},
		ForbiddenFeature{Feature: "au_bpay_deposit"},
		ForbiddenFeature{Feature: "au_npp_deposit"},
		ForbiddenFeature{Feature: "au_bpay_account_creation"},
		ForbiddenFeature{Feature: "sepa_account_creation"},
		ForbiddenFeature{Feature: "au_npp_account_creation"},
		ForbiddenFeature{Feature: "eur_fiat_to_crypto"},
		ForbiddenFeature{Feature: "eur_to_card_top_up"},
		ForbiddenFeature{Feature: "eur_crypto_to_fiat"},
		ForbiddenFeature{Feature: "aud_crypto_to_fiat"},
		ForbiddenFeature{Feature: "aud_fiat_to_crypto"},
		ForbiddenFeature{Feature: "aud_to_card_top_up"},
	}

	s.EntityRestrictions.Value["australia"] = []ForbiddenFeature{
		ForbiddenFeature{Feature: "fiat_withdrawal", RequiredPersonalInformation: "residential_address:submit"},
		ForbiddenFeature{Feature: "pay_your_friends", RequiredPersonalInformation: "residential_address:submit"},
		ForbiddenFeature{Feature: "crypto_withdrawal", RequiredPersonalInformation: "residential_address:submit"},
		ForbiddenFeature{Feature: "deposit_to_ex", RequiredPersonalInformation: "residential_address:submit"},
		ForbiddenFeature{Feature: "fiat_to_card_top_up", RequiredPersonalInformation: "residential_address:submit"},
		ForbiddenFeature{Feature: "crypto_to_card_top_up", RequiredPersonalInformation: "residential_address:submit"},
		ForbiddenFeature{Feature: "sepa_account_creation"},
		ForbiddenFeature{Feature: "sepa_deposit"},
		ForbiddenFeature{Feature: "sepa_withdrawal"},
		ForbiddenFeature{Feature: "eur_crypto_to_fiat"},
		ForbiddenFeature{Feature: "eur_fiat_to_crypto"},
		ForbiddenFeature{Feature: "eur_to_card_top_up"},
	}
	return nil
}

func (s *SystemConfig) BooleanEqual(key string, value bool) bool {
	if val, ok := s.Value[key].(bool); ok {
		return val == value
	}
	return false
}

func (s *SystemConfig) Include(key string, value string) bool {
	if val, ok := s.Value[key].([]string); ok {
		return slices.Contains(val, value)
	}
	return false
}
