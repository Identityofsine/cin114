package config

import (
	"fmt"
	"os"
	"strings"
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

// validateStripeSecrets checks if Stripe secrets are properly configured
func validateStripeSecrets(config *StripeSettings) error {
	var errors []string

	// Check if we're in development mode
	env := os.Getenv("GO_ENV")
	isDev := env == "development" || env == "dev"

	// Validate Stripe Secret Key
	if config.StripeSecretKey == "" || config.StripeSecretKey == "sk_test_placeholder" {
		if !isDev {
			errors = append(errors, "STRIPE_SECRET_KEY is required in production")
		}
	} else if !strings.HasPrefix(config.StripeSecretKey, "sk_") {
		errors = append(errors, "STRIPE_SECRET_KEY must start with 'sk_'")
	}

	// Validate Webhook Secret (optional in dev)
	if config.WebhookSecret == "" || config.WebhookSecret == "whsec_placeholder" {
		if !isDev {
			errors = append(errors, "STRIPE_WEBHOOK_SECRET is required in production")
		}
	} else if !strings.HasPrefix(config.WebhookSecret, "whsec_") {
		errors = append(errors, "STRIPE_WEBHOOK_SECRET must start with 'whsec_'")
	}

	// Validate Redirect URL
	if config.StripeRedirectURL == "" {
		errors = append(errors, "STRIPE_REDIRECT_URL is required")
	}

	if len(errors) > 0 {
		return fmt.Errorf("Stripe configuration errors: %s", strings.Join(errors, "; "))
	}

	return nil
}

func GetStripeSettings() *StripeSettings {
	if stripeCachedConfig != nil {
		return stripeCachedConfig
	}
	if config, err := loadConfig[*StripeSettings]("stripe"); err == nil {
		config.StripeSecretKey = os.Getenv("STRIPE_SECRET_KEY")
		config.WebhookSecret = os.Getenv("STRIPE_WEBHOOK_SECRET")
		config.StripeRedirectURL = os.Getenv("STRIPE_REDIRECT_URL")

		// Validate the configuration
		if validationErr := validateStripeSecrets(config); validationErr != nil {
			fmt.Printf("Warning: %v\n", validationErr)
		}

		stripeCachedConfig = config
		return config
	} else {
		fmt.Printf("Error loading stripe config: %v\n", err)
		config := &StripeSettings{
			StripeSecretKey:   os.Getenv("STRIPE_SECRET_KEY"),
			WebhookSecret:     os.Getenv("STRIPE_WEBHOOK_SECRET"),
			StripeRedirectURL: os.Getenv("STRIPE_REDIRECT_URL"),
		}

		// Validate the configuration
		if validationErr := validateStripeSecrets(config); validationErr != nil {
			fmt.Printf("Warning: %v\n", validationErr)
		}

		return config
	}
}
