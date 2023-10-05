package config

import (
	"github.com/gilperopiola/go-rest-example/pkg/utils"
)

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

var prefix = "GO_REST_EXAMPLE_"

func (config *Config) Setup() {

	// General configuration with defaults
	config.PORT = utils.GetEnv(prefix+"PORT", defaultPort)
	config.DEBUG = utils.GetEnvBool(prefix+"DEBUG", defaultDebug)

	// Database configuration with defaults
	config.DATABASE.TYPE = utils.GetEnv(prefix+"DATABASE_TYPE", defaultDatabaseType)
	config.DATABASE.USERNAME = utils.GetEnv(prefix+"DATABASE_USERNAME", defaultDatabaseUsername)
	config.DATABASE.PASSWORD = utils.GetEnv(prefix+"DATABASE_PASSWORD", defaultDatabasePassword)
	config.DATABASE.HOSTNAME = utils.GetEnv(prefix+"DATABASE_HOSTNAME", defaultDatabaseHostname)
	config.DATABASE.PORT = utils.GetEnv(prefix+"DATABASE_PORT", defaultDatabasePort)
	config.DATABASE.SCHEMA = utils.GetEnv(prefix+"DATABASE_SCHEMA", defaultDatabaseSchema)
	config.DATABASE.PURGE = utils.GetEnvBool(prefix+"DATABASE_PURGE", defaultDatabasePurge)
	config.DATABASE.DEBUG = utils.GetEnvBool(prefix+"DATABASE_DEBUG", defaultDatabaseDebug)

	// JWT configuration with defaults
	config.JWT.SECRET = utils.GetEnv(prefix+"JWT_SECRET", defaultJWTSecret)
	config.JWT.SESSION_DURATION_DAYS = utils.GetEnvInt(prefix+"JWT_SESSION_DURATION_DAYS", defaultJWTSessionDurationDays)

	config.validateEnvVars()
}

var (
	defaultPort  = "8040"
	defaultDebug = false

	defaultDatabaseType     = "mysql"
	defaultDatabaseUsername = "root"
	defaultDatabasePassword = ""
	defaultDatabaseHostname = "127.0.0.1"
	defaultDatabasePort     = "3306"
	defaultDatabaseSchema   = "go-rest-example-db"
	defaultDatabasePurge    = false
	defaultDatabaseDebug    = false

	defaultJWTSecret              = "a0#3ndl2"
	defaultJWTSessionDurationDays = 14
)
