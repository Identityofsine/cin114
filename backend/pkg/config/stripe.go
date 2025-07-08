package config

import (
	"fmt"
	"os"
)

// StripeSettings contains the configuration for Stripe payment processing

type StripeSettings struct {
	ConfigFile
	StripeSecretKey   string
	StripeRedirectURL string
	WebhookSecret     string
}

var (
	stripeCachedConfig *StripeSettings
)

func GetStripeSettings() *StripeSettings {
	if stripeCachedConfig != nil {
		return stripeCachedConfig
	}
	if config, err := loadConfig[*StripeSettings]("stripe"); err == nil {
		config.StripeSecretKey = os.Getenv("STRIPE_API_KEY")
		config.WebhookSecret = os.Getenv("STRIPE_WEBHOOK_SECRET")
		config.StripeRedirectURL = os.Getenv("STRIPE_REDIRECT_URL")
		stripeCachedConfig = config
		return config
	} else {
		fmt.Printf("Error loading stripe config: %v\n", err)
		return &StripeSettings{
			StripeSecretKey:   os.Getenv("STRIPE_API_KEY"),
			WebhookSecret:     os.Getenv("STRIPE_WEBHOOK_SECRET"),
			StripeRedirectURL: os.Getenv("STRIPE_REDIRECT_URL"),
		}
	}
}
