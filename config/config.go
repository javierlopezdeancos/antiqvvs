package config

import (
	"os"
)

// Config configuration structure to our stripe integration
type Config struct {
	StripePublishableKey string   `json:"stripePublishableKey"`
	StripeCountry        string   `json:"stripeCountry"`
	Country              string   `json:"country"`
	Currency             string   `json:"currency"`
	PaymentMethods       []string `json:"paymentMethods"`
}

// Default get default values to stripe integration
func Default() Config {
	stripeCountry := os.Getenv("STRIPE_ACCOUNT_COUNTRY")

	if stripeCountry == "" {
		stripeCountry = "US"
	}

	return Config{
		StripePublishableKey: os.Getenv("STRIPE_PUBLISHABLE_KEY"),
		StripeCountry:        stripeCountry,
		Country:              os.Getenv("STRIPE_ACCOUNT_COUNTRY"),
		Currency:             os.Getenv("STRIPE_ACCOUNT_CURRENCY"),
	}
}
