// internal/config/config.go
package config

import (
    "log"

    "github.com/spf13/viper"
)

type Config struct {
    ResyAPIKey  string
    ResyAuthKey string
    Port        string
    Debug       bool
}

func Load() (*Config, error) {
    viper.SetConfigFile(".env")
    // Set up Viper to read environment variables
    viper.AutomaticEnv()

    
    if viper.GetBool("DEBUG") {
        log.Printf("📋 Configuration:")
        log.Printf("API Key: %s", viper.GetString("RESY_API_KEY"))
        log.Printf("Auth Token: %s", viper.GetString("RESY_AUTH_TOKEN"))
    }

    return &Config{
        ResyAPIKey:  viper.GetString("RESY_API_KEY"),
        ResyAuthKey: viper.GetString("RESY_AUTH_TOKEN"),
        Port:        viper.GetString("PORT"),
        Debug:       viper.GetBool("DEBUG"),
    }, nil
}