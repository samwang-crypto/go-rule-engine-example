package models

import "golang.org/x/exp/slices"

type SystemConfig struct {
	Value map[string]interface{}
}

// this Load function is used to load the system config from the database / redis
func (s *SystemConfig) Load() error {
	//TODO: replace with acutal redis / database code
	// here just to simulate loading from the database
	s.Value["RecurringBuyEnabled"] = false
	s.Value["RecurringBuyPurchaseByCreditCardEnabled"] = true
	s.Value["RecurringBuyPurchaseByCreditCardInternalTesters"] = []string{}
	s.Value["RecurringBuyPurchaseByStableCoinEnabled"] = false
	s.Value["RecurringBuyPurchaseByStableCoinInternalTesters"] = []string{}
	s.Value["RecurringBuyPurchaseByFiatWalletEnabled"] = false
	s.Value["RecurringBuyPurchaseByFiatWalletInternalTesters"] = []string{}
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
