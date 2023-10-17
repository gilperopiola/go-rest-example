package config

import (
	"os"
	"strconv"
)

// Env vars prefix, likely your app's name => GO_REST_EXAMPLE_PORT, GO_REST_EXAMPLE_DATABASE_TYPE, etc.
const prefix = "GO_REST_EXAMPLE_"

type Config struct {
	general    GeneralConfig
	database   DatabaseConfig
	jwt        JWTConfig
	monitoring MonitoringConfig
}

func NewConfig() *Config {
	config := Config{}
	config.Setup()
	return &config
}

func NewTestConfig() *Config {
	config := Config{}
	config.Setup()
	config.database.Purge = true
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

func (config *Config) loadGeneralVars() {
	config.general.Port = getEnv(prefix+"PORT", defaultPort)
	config.general.Debug = getEnvBool(prefix+"DEBUG", defaultDebug)
	config.general.Timeout = getEnvInt(prefix+"TIMEOUT_SECONDS", defaultTimeoutSeconds)
}

func (config *Config) loadDatabaseVars() {
	config.database.Type = getEnv(prefix+"DATABASE_TYPE", defaultDatabaseType)
	config.database.Username = getEnv(prefix+"DATABASE_USERNAME", defaultDatabaseUsername)
	config.database.Password = getEnv(prefix+"DATABASE_PASSWORD", defaultDatabasePassword)
	config.database.Hostname = getEnv(prefix+"DATABASE_HOSTNAME", defaultDatabaseHostname)
	config.database.Port = getEnv(prefix+"DATABASE_PORT", defaultDatabasePort)
	config.database.Schema = getEnv(prefix+"DATABASE_SCHEMA", defaultDatabaseSchema)
	config.database.Purge = getEnvBool(prefix+"DATABASE_PURGE", defaultDatabasePurge)
	config.database.Debug = getEnvBool(prefix+"DATABASE_DEBUG", defaultDatabaseDebug)
	config.database.Destroy = getEnvBool(prefix+"DATABASE_DESTROY", defaultDatabaseDestroy)
}

func (config *Config) loadJWTVars() {
	config.jwt.Secret = getEnv(prefix+"JWT_SECRET", defaultJWTSecret)
	config.jwt.SessionDurationDays = getEnvInt(prefix+"JWT_SESSION_DURATION_DAYS", defaultJWTSessionDurationDays)
	config.jwt.HashSalt = getEnv(prefix+"JWT_HASH_SALT", defaultJWTHashSalt)
}

func (config *Config) loadMonitoringVars() {
	config.monitoring.Enabled = getEnvBool(prefix+"MONITORING_ENABLED", defaultMonitoringEnabled)
	config.monitoring.AppName = getEnv(prefix+"MONITORING_APP_NAME", defaultMonitoringAppName)
	config.monitoring.Secret = getEnv(prefix+"MONITORING_SECRET", defaultMonitoringSecret)
}

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
