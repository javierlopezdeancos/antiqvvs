package environment

import (
	"os"
	"path"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v72"
)

// Types to different environments
var Types = map[string]string{
	"test":       "test",
	"production": "prod",
}

// Configuration configuration type from env file
type Configuration struct {
	EnvType   string
	StripeKey string
}

// Config configuration from env file
var Config Configuration

// LoadEnv load environment configuration and configure stripe from him
func LoadEnv() error {
	err := godotenv.Load(path.Join("./", ".env"))

	if err != nil {
		return err
	}

	Config.StripeKey = os.Getenv("STRIPE_SECRET_KEY")
	stripe.Key = Config.StripeKey

	if stripe.Key == "" {
		panic("STRIPE_SECRET_KEY must be in environment")
	}

	Config.EnvType = os.Getenv("TYPE")

	if Config.EnvType == "" {
		panic("TYPE must be in environment")
	}

	return nil
}
