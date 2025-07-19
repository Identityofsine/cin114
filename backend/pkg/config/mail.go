package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// MailSettings contains the configuration for SMTP email sending
type MailSettings struct {
	ConfigFile
	SMTPServer   string `yaml:"smtp_server"`
	SMTPPort     int    `yaml:"smtp_port"`
	SMTPUsername string `yaml:"smtp_username"`
	SMTPPassword string `yaml:"smtp_password"`
	FromEmail    string `yaml:"from_email"`
	FromName     string `yaml:"from_name"`
}

var (
	mailCachedConfig *MailSettings
)

// validateMailSecrets checks if mail secrets are properly configured
func validateMailSecrets(config *MailSettings) error {
	var errors []string

	// Check if we're in development mode
	env := os.Getenv("GO_ENV")
	isDev := env == "development" || env == "dev"

	// Validate SMTP Server
	if config.SMTPServer == "" {
		if !isDev {
			errors = append(errors, "SMTP_SERVER is required in production")
		}
	}

	// Validate SMTP Port
	if config.SMTPPort == 0 {
		if !isDev {
			errors = append(errors, "SMTP_PORT is required in production")
		}
	} else if config.SMTPPort < 1 || config.SMTPPort > 65535 {
		errors = append(errors, "SMTP_PORT must be between 1 and 65535")
	}

	// Validate SMTP Username
	if config.SMTPUsername == "" || config.SMTPUsername == "your-email@gmail.com" {
		if !isDev {
			errors = append(errors, "SMTP_USERNAME is required in production")
		}
	}

	// Validate SMTP Password
	if config.SMTPPassword == "" || config.SMTPPassword == "your-app-password" {
		if !isDev {
			errors = append(errors, "SMTP_PASSWORD is required in production")
		}
	}

	// Validate From Email
	if config.FromEmail == "" {
		errors = append(errors, "SMTP_FROM_EMAIL is required")
	} else if !strings.Contains(config.FromEmail, "@") {
		errors = append(errors, "SMTP_FROM_EMAIL must be a valid email address")
	}

	// Validate From Name
	if config.FromName == "" {
		config.FromName = "Cin114 Tickets" // Default value
	}

	if len(errors) > 0 {
		return fmt.Errorf("Mail configuration errors: %s", strings.Join(errors, "; "))
	}

	return nil
}

func GetMailSettings() *MailSettings {
	if mailCachedConfig != nil {
		return mailCachedConfig
	}

	// Try to load from YAML config first
	if config, err := loadConfig[*MailSettings]("mail"); err == nil {
		// If YAML values are empty, fall back to environment variables
		if config.SMTPServer == "" {
			// Only load .env file if it exists (for local development)
			if _, err := os.Stat(".env"); err == nil {
				godotenv.Load(".env")
			}
			config.SMTPServer = os.Getenv("SMTP_SERVER")
		}

		if config.SMTPPort == 0 {
			if portStr := os.Getenv("SMTP_PORT"); portStr != "" {
				if port, err := strconv.Atoi(portStr); err == nil {
					config.SMTPPort = port
				}
			}
		}

		if config.SMTPUsername == "" {
			config.SMTPUsername = os.Getenv("SMTP_USERNAME")
		}

		if config.SMTPPassword == "" {
			config.SMTPPassword = os.Getenv("SMTP_PASSWORD")
		}

		if config.FromEmail == "" {
			config.FromEmail = os.Getenv("SMTP_FROM_EMAIL")
		}

		if config.FromName == "" {
			config.FromName = os.Getenv("SMTP_FROM_NAME")
		}

		// Debug: Print what we loaded (without exposing the password)
		fmt.Printf("Mail config loaded - SMTP Server: %s, Port: %d\n", config.SMTPServer, config.SMTPPort)
		fmt.Printf("Mail config loaded - Username: %s, From: %s <%s>\n", config.SMTPUsername, config.FromName, config.FromEmail)
		if len(config.SMTPPassword) > 0 {
			fmt.Printf("Mail config loaded - Password length: %d\n", len(config.SMTPPassword))
		} else {
			fmt.Printf("Mail config loaded - Password is empty\n")
		}

		// Validate the configuration
		if validationErr := validateMailSecrets(config); validationErr != nil {
			fmt.Printf("Warning: %v\n", validationErr)
		}

		mailCachedConfig = config
		return config
	} else {
		fmt.Printf("Error loading mail config: %v\n", err)
		// Fallback to environment variables only
		if _, err := os.Stat(".env"); err == nil {
			godotenv.Load(".env")
		}

		smtpPort := 587 // Default Gmail SMTP port
		if portStr := os.Getenv("SMTP_PORT"); portStr != "" {
			if port, err := strconv.Atoi(portStr); err == nil {
				smtpPort = port
			}
		}

		config := &MailSettings{
			SMTPServer:   getEnvOrDefault("SMTP_SERVER", "smtp.gmail.com"),
			SMTPPort:     smtpPort,
			SMTPUsername: os.Getenv("SMTP_USERNAME"),
			SMTPPassword: os.Getenv("SMTP_PASSWORD"),
			FromEmail:    os.Getenv("SMTP_FROM_EMAIL"),
			FromName:     getEnvOrDefault("SMTP_FROM_NAME", "Cin114 Tickets"),
		}

		// Debug: Print fallback config
		fmt.Printf("Using fallback environment variables for mail config\n")
		fmt.Printf("Mail config (fallback) - SMTP Server: %s, Port: %d\n", config.SMTPServer, config.SMTPPort)
		fmt.Printf("Mail config (fallback) - Username: %s, From: %s <%s>\n", config.SMTPUsername, config.FromName, config.FromEmail)

		// Validate the configuration
		if validationErr := validateMailSecrets(config); validationErr != nil {
			fmt.Printf("Warning: %v\n", validationErr)
		}

		mailCachedConfig = config
		return config
	}
}

// getEnvOrDefault returns environment variable value or default if not set
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
