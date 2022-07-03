package models

// user's current configuration
type UserConfig struct {
	RecurringBuyEnabled bool `json:"recurring_buy_enabled"`
}

// stores the user's information as facts
type User struct {
	ID     string      `json:"id"`
	Email  string      `json:"email"`
	Config *UserConfig `json:"config"`
}
