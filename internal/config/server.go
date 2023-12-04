package config

import (
	"os"
)

// ConfigServer represents the configuration settings for the application.
type ConfigServer struct {
	ServerPort  string
	DatabaseURL string
	// Add other configuration parameters as needed
}

// LoadConfigServer loads the configuration settings from environment variables or a configuration file.
func LoadConfigServer() *ConfigServer {
	serverPort := getEnv("SERVER_PORT", "8080")
	databaseURL := getEnv("DATABASE_URL", "your_database_connection_string")

	return &ConfigServer{
		ServerPort:  serverPort,
		DatabaseURL: databaseURL,
		// Initialize other configuration parameters here
	}
}

// getEnv retrieves the value of an environment variable with a fallback.
func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}
