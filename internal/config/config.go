// internal/config/config.go
package config

import (
	"os"
)

type Config struct {
    ResyAPIKey  string
    ResyAuthKey string
    Port        string
    Debug       bool
}

func Load() (*Config, error) {
    return &Config{
        ResyAPIKey:  os.Getenv("RESY_API_KEY"),
        ResyAuthKey: os.Getenv("RESY_AUTH_KEY"),
        Port:       getEnvOrDefault("PORT", "8080"),
        Debug:       os.Getenv("DEBUG") == "true",
    }, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
