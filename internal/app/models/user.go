package models

// user's current configuration
type UserConfig struct {
	RecurringBuyEnabled bool
}

// stores the user's information as facts
type User struct {
	Email  string
	Config *UserConfig
}
