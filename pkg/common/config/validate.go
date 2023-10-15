package config

import (
	"log"
)

// Validate environment variables which need to be set
func (config *Config) validate() {
	var (
		necessaryVars = map[string]string{
			prefix + "PORT":            config.PORT,
			prefix + "DATABASE_TYPE":   config.DATABASE.TYPE,
			prefix + "DATABASE_SCHEMA": config.DATABASE.SCHEMA,
		}
		missingVars = []string{}
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
