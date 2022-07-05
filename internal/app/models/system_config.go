package models

import (
	"time"

	"golang.org/x/exp/slices"
)

type SystemConfig struct {
	Value              map[string]interface{}
	EntityRestrictions EntityRestrictions
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

	t, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-05-27 15:00:00", time.Local)
	if err == nil {
		s.Value["CaAppliedForKycApprovedAfter"] = t
	}

	s.EntityRestrictions["canada"] = []ForbiddenFeature{
		{Feature: "crypto_deposit", RequiredPersonalInformation: "ca_dcbank_application:approve"},
		{Feature: "credit_card_purchase", RequiredPersonalInformation: "ca_dcbank_application:approve"},
		{Feature: "pay_your_friends_receive", RequiredPersonalInformation: "ca_dcbank_application:approve"},
		{Feature: "exchange_to_crypto_transfer", RequiredPersonalInformation: "ca_dcbank_application:approve"},
		{Feature: "sepa_withdrawal"},
		{Feature: "sepa_deposit"},
		{Feature: "eur_fiat_to_crypto"},
		{Feature: "eur_crypto_to_fiat"},
		{Feature: "sepa_account_creation"},
		{Feature: "eur_to_card_top_up"},
		{Feature: "card_purchase_crypto"},
		{Feature: "exchange_to_crypto"},
	}
	s.EntityRestrictions["australia"] = []ForbiddenFeature{
		{Feature: "fiat_withdrawal", RequiredPersonalInformation: "residential_address:submit"},
		{Feature: "pay_your_friends", RequiredPersonalInformation: "residential_address:submit"},
		{Feature: "crypto_withdrawal", RequiredPersonalInformation: "residential_address:submit"},
		{Feature: "deposit_to_ex", RequiredPersonalInformation: "residential_address:submit"},
		{Feature: "fiat_to_card_top_up", RequiredPersonalInformation: "residential_address:submit"},
		{Feature: "crypto_to_card_top_up", RequiredPersonalInformation: "residential_address:submit"},
		{Feature: "sepa_account_creation"},
		{Feature: "sepa_deposit"},
		{Feature: "sepa_withdrawal"},
		{Feature: "eur_crypto_to_fiat"},
		{Feature: "eur_fiat_to_crypto"},
		{Feature: "eur_to_card_top_up"},
	}
	return nil
}

func (s *SystemConfig) TimeLessThan(key string, value string) bool {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
	if err == nil {
		if val, ok := s.Value[key].(time.Time); ok {
			return t.Before(val)
		}
	}
	return false
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
