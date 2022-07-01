package models

type SystemConfig struct {
	RecurringBuyEnabled                             bool
	RecurringBuyPurchaseByCreditCardEnabled         bool
	RecurringBuyPurchaseByCreditCardInternalTesters []string
	RecurringBuyPurchaseByStableCoinEnabled         bool
	RecurringBuyPurchaseByStableCoinInternalTesters []string
	RecurringBuyPurchaseByFiatWalletEnabled         bool
	RecurringBuyPurchaseByFiatWalletInternalTesters []string
}
