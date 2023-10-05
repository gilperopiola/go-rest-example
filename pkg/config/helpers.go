package config

import (
	"log"
)

func (config *Config) GetPort() string {
	return config.PORT
}

func (config *Config) GetDebugMode() bool {
	return config.DEBUG
}

func (config *Config) GetDatabaseConfig() DatabaseConfig {
	return config.DATABASE
}

func (config *Config) GetJWTConfig() JWTConfig {
	return config.JWT
}

func (config *DatabaseConfig) GetConnectionString() string {
	username := config.USERNAME
	password := config.PASSWORD
	hostname := config.HOSTNAME
	port := config.PORT
	schema := config.SCHEMA
	params := "?charset=utf8&parseTime=True&loc=Local"

	return username + ":" + password + "@tcp(" + hostname + ":" + port + ")/" + schema + params
}

// Validate environment variables which need to be set
func (config *Config) validateEnvVars() {
	var (
		missingVars   = []string{}
		necessaryVars = map[string]string{
			prefix + "PORT":            config.PORT,
			prefix + "DATABASE_TYPE":   config.DATABASE.TYPE,
			prefix + "DATABASE_SCHEMA": config.DATABASE.SCHEMA,
		}
	)

	// Check if each necessary variable is set
	for name, value := range necessaryVars {
		if value == "" {
			missingVars = append(missingVars, name)
		}
	}

	// If there are any missing variables, log an error
	if len(missingVars) > 0 {
		log.Fatalf("error validating environment variables, not set: %v", missingVars)
	}
}
