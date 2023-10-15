package config

import (
	"os"
	"strconv"
)

// Env vars prefix => GO_REST_EXAMPLE_PORT, GO_REST_EXAMPLE_DATABASE_TYPE, etc.
const prefix = "GO_REST_EXAMPLE_"

type ConfigI interface {
	Setup()

	GetDebugMode() bool
	GetPort() string
	GetDatabaseConfig() DatabaseConfig
	GetJWTConfig() JWTConfig
	GetMonitoringConfig() MonitoringConfig
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

func (config *Config) loadGeneralVars() {
	config.PORT = getEnv(prefix+"PORT", defaultPort)
	config.DEBUG = getEnvBool(prefix+"DEBUG", defaultDebug)
}

func (config *Config) loadDatabaseVars() {
	config.DATABASE.TYPE = getEnv(prefix+"DATABASE_TYPE", defaultDatabaseType)
	config.DATABASE.USERNAME = getEnv(prefix+"DATABASE_USERNAME", defaultDatabaseUsername)
	config.DATABASE.PASSWORD = getEnv(prefix+"DATABASE_PASSWORD", defaultDatabasePassword)
	config.DATABASE.HOSTNAME = getEnv(prefix+"DATABASE_HOSTNAME", defaultDatabaseHostname)
	config.DATABASE.PORT = getEnv(prefix+"DATABASE_PORT", defaultDatabasePort)
	config.DATABASE.SCHEMA = getEnv(prefix+"DATABASE_SCHEMA", defaultDatabaseSchema)
	config.DATABASE.PURGE = getEnvBool(prefix+"DATABASE_PURGE", defaultDatabasePurge)
	config.DATABASE.DEBUG = getEnvBool(prefix+"DATABASE_DEBUG", defaultDatabaseDebug)
}

func (config *Config) loadJWTVars() {
	config.JWT.SECRET = getEnv(prefix+"JWT_SECRET", defaultJWTSecret)
	config.JWT.SESSION_DURATION_DAYS = getEnvInt(prefix+"JWT_SESSION_DURATION_DAYS", defaultJWTSessionDurationDays)
}

func (config *Config) loadMonitoringVars() {
	config.MONITORING.ENABLED = getEnvBool(prefix+"MONITORING_ENABLED", defaultMonitoringEnabled)
	config.MONITORING.APP_NAME = getEnv(prefix+"MONITORING_APP_NAME", defaultMonitoringAppName)
	config.MONITORING.SECRET = getEnv(prefix+"MONITORING_SECRET", defaultMonitoringSecret)
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
