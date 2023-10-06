package config

import "github.com/gilperopiola/go-rest-example/pkg/utils"

// Env vars prefix => GO_REST_EXAMPLE_PORT, GO_REST_EXAMPLE_DATABASE_TYPE, etc.
const prefix = "GO_REST_EXAMPLE_"

type ConfigInterface interface {
	Setup()

	GetDebugMode() bool
	GetPort() string
	GetDatabaseConfig() DatabaseConfig
	GetJWTConfig() JWTConfig
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

	// Validate required variables, if not present, panic
	config.validate()
}

func (config *Config) loadGeneralVars() {
	config.PORT = utils.GetEnv(prefix+"PORT", defaultPort)
	config.DEBUG = utils.GetEnvBool(prefix+"DEBUG", defaultDebug)
}

func (config *Config) loadDatabaseVars() {
	config.DATABASE.TYPE = utils.GetEnv(prefix+"DATABASE_TYPE", defaultDatabaseType)
	config.DATABASE.USERNAME = utils.GetEnv(prefix+"DATABASE_USERNAME", defaultDatabaseUsername)
	config.DATABASE.PASSWORD = utils.GetEnv(prefix+"DATABASE_PASSWORD", defaultDatabasePassword)
	config.DATABASE.HOSTNAME = utils.GetEnv(prefix+"DATABASE_HOSTNAME", defaultDatabaseHostname)
	config.DATABASE.PORT = utils.GetEnv(prefix+"DATABASE_PORT", defaultDatabasePort)
	config.DATABASE.SCHEMA = utils.GetEnv(prefix+"DATABASE_SCHEMA", defaultDatabaseSchema)
	config.DATABASE.PURGE = utils.GetEnvBool(prefix+"DATABASE_PURGE", defaultDatabasePurge)
	config.DATABASE.DEBUG = utils.GetEnvBool(prefix+"DATABASE_DEBUG", defaultDatabaseDebug)
}

func (config *Config) loadJWTVars() {
	config.JWT.SECRET = utils.GetEnv(prefix+"JWT_SECRET", defaultJWTSecret)
	config.JWT.SESSION_DURATION_DAYS = utils.GetEnvInt(prefix+"JWT_SESSION_DURATION_DAYS", defaultJWTSessionDurationDays)
}
