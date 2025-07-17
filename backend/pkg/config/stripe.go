package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// StripeSettings contains the configuration for Stripe payment processing

type StripeSettings struct {
	ConfigFile
	StripeSecretKey   string `yaml:"stripe_secret_key"`
	StripeRedirectURL string `yaml:"stripe_redirect_url"`
	WebhookSecret     string `yaml:"webhook_secret"`
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

	// Try to load from YAML config first
	if config, err := loadConfig[*StripeSettings]("stripe"); err == nil {
		// If YAML values are empty, fall back to environment variables
		if config.StripeSecretKey == "" {
			// Only load .env file if it exists (for local development)
			if _, err := os.Stat(".env"); err == nil {
				godotenv.Load(".env")
			}
			config.StripeSecretKey = os.Getenv("STRIPE_SECRET_KEY")
		}

		if config.WebhookSecret == "" {
			config.WebhookSecret = os.Getenv("STRIPE_WEBHOOK_SECRET")
		}

		if config.StripeRedirectURL == "" {
			config.StripeRedirectURL = os.Getenv("STRIPE_REDIRECT_URL")
		}

		// Debug: Print what we loaded (without exposing the full key)
		secretKeyLen := len(config.StripeSecretKey)
		if secretKeyLen > 0 {
			fmt.Printf("Stripe config loaded - Secret key length: %d, starts with: %s\n",
				secretKeyLen, config.StripeSecretKey[:7])
		} else {
			fmt.Printf("Stripe config loaded - Secret key is empty\n")
		}

		fmt.Printf("Stripe webhook secret length: %d\n", len(config.WebhookSecret))
		fmt.Printf("Stripe redirect URL: %s\n", config.StripeRedirectURL)

		// Validate the configuration
		if validationErr := validateStripeSecrets(config); validationErr != nil {
			fmt.Printf("Warning: %v\n", validationErr)
		}

		stripeCachedConfig = config
		return config
	} else {
		fmt.Printf("Error loading stripe config: %v\n", err)
		// Fallback to environment variables only
		if _, err := os.Stat(".env"); err == nil {
			godotenv.Load(".env")
		}

		config := &StripeSettings{
			StripeSecretKey:   os.Getenv("STRIPE_SECRET_KEY"),
			WebhookSecret:     os.Getenv("STRIPE_WEBHOOK_SECRET"),
			StripeRedirectURL: os.Getenv("STRIPE_REDIRECT_URL"),
		}

		// Debug: Print fallback config
		fmt.Printf("Using fallback environment variables for Stripe config\n")
		secretKeyLen := len(config.StripeSecretKey)
		if secretKeyLen > 0 {
			fmt.Printf("Stripe config (fallback) - Secret key length: %d, starts with: %s\n",
				secretKeyLen, config.StripeSecretKey[:7])
		} else {
			fmt.Printf("Stripe config (fallback) - Secret key is empty\n")
		}

		// Validate the configuration
		if validationErr := validateStripeSecrets(config); validationErr != nil {
			fmt.Printf("Warning: %v\n", validationErr)
		}

		return config
	}
}
