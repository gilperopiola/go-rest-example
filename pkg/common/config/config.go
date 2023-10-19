package config

import (
	"os"
	"strconv"
)

// Env vars prefix, likely your app's name => GO_REST_EXAMPLE_PORT, GO_REST_EXAMPLE_DATABASE_TYPE, etc.
const prefix = "GO_REST_EXAMPLE_"

type Config struct {
	General    General
	Database   Database
	JWT        JWT
	Monitoring Monitoring
}

func NewConfig() *Config {
	config := Config{}
	config.Setup()
	return &config
}

func (config *Config) Setup() {

	// Load variables
	config.loadGeneralVars()
	config.loadDatabaseVars()
	config.loadJWTVars()
	config.loadMonitoringVars()

	// Validate required variables, if not present, panic
	config.validate()
}

//-----------------------
//      LOAD VARS
//-----------------------

func (config *Config) loadGeneralVars() {
	config.General.Port = getEnv(prefix+"PORT", defaultPort)
	config.General.Debug = getEnvBool(prefix+"DEBUG", defaultDebug)
	config.General.Timeout = getEnvInt(prefix+"TIMEOUT_SECONDS", defaultTimeoutSeconds)
}

func (config *Config) loadDatabaseVars() {
	config.Database.Type = getEnv(prefix+"DATABASE_TYPE", defaultDatabaseType)
	config.Database.Username = getEnv(prefix+"DATABASE_USERNAME", defaultDatabaseUsername)
	config.Database.Password = getEnv(prefix+"DATABASE_PASSWORD", defaultDatabasePassword)
	config.Database.Hostname = getEnv(prefix+"DATABASE_HOSTNAME", defaultDatabaseHostname)
	config.Database.Port = getEnv(prefix+"DATABASE_PORT", defaultDatabasePort)
	config.Database.Schema = getEnv(prefix+"DATABASE_SCHEMA", defaultDatabaseSchema)
	config.Database.Purge = getEnvBool(prefix+"DATABASE_PURGE", defaultDatabasePurge)
	config.Database.Debug = getEnvBool(prefix+"DATABASE_DEBUG", defaultDatabaseDebug)
	config.Database.Destroy = getEnvBool(prefix+"DATABASE_DESTROY", defaultDatabaseDestroy)
}

func (config *Config) loadJWTVars() {
	config.JWT.Secret = getEnv(prefix+"JWT_SECRET", defaultJWTSecret)
	config.JWT.SessionDurationDays = getEnvInt(prefix+"JWT_SESSION_DURATION_DAYS", defaultJWTSessionDurationDays)
	config.JWT.HashSalt = getEnv(prefix+"JWT_HASH_SALT", defaultJWTHashSalt)
}

func (config *Config) loadMonitoringVars() {
	config.Monitoring.Enabled = getEnvBool(prefix+"MONITORING_ENABLED", defaultMonitoringEnabled)
	config.Monitoring.AppName = getEnv(prefix+"MONITORING_APP_NAME", defaultMonitoringAppName)
	config.Monitoring.Secret = getEnv(prefix+"MONITORING_SECRET", defaultMonitoringSecret)
}

//-----------------------
//       HELPERS
//-----------------------

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value == "true" || value == "1"
}

func getEnvInt(key string, fallback int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return intValue
}
