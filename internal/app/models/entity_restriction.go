package models

var fiatFeatures = []string{"sepa_deposit", "au_npp_deposit", "au_bpay_deposit", "ca_interac_etransfer_deposit"}
var skipForAppliedAfter = []string{"card_purchase_crypto", "crypto_deposit", "pay_your_friends_receive", "exchange_to_crypto"}

// This is just temporary, it will be replaced by whitelist
type ForbiddenFeature struct {
	Feature                     string
	RequiredPersonalInformation string
}

type EntityRestrictions map[string][]ForbiddenFeature
