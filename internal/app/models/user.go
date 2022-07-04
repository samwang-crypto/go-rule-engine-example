package models

import "time"

// user's current configuration
type UserConfig struct {
	RecurringBuyEnabled    bool `json:"recurring_buy_enabled"`
	CryptoEnabled          bool `json:"crypto_enabled"`
	FiatEnabled            bool `json:"fiat_enabled"`
	CardApplicationEnabled bool `json:"card_application_enabled"`
	HasCryptoFiatAccount   bool `json:"has_crypto_fiat_account"`
}

// user's KYC document state
type KycDocument struct {
	IssuingCountry           string    `json:"issuing_country"`
	UsState                  string    `json:"us_state"`
	Type                     string    `json:"type"`
	Verification             string    `json:"verification"`
	TermsOfServiceAcceptedAt time.Time `json:"terms_of_service_accepted_at"`
	ApprovedAt               time.Time `json:"approved_at"`
}

// stores the user's information as facts
type User struct {
	UUID               string       `json:"uuid"`
	Email              string       `json:"email"`
	PhoneCountry       string       `json:"phone_country"`
	EntityId           string       `json:"entity_id"`
	ResidentialAddress string       `json:"residential_address"`
	KycDocument        *KycDocument `json:"kyc_document"`
	Config             *UserConfig  `json:"config"`
}
