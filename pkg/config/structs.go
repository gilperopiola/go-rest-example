package config

type Config struct {
	PORT  string
	DEBUG bool

	DATABASE DatabaseConfig
	JWT      JWTConfig
}

type DatabaseConfig struct {
	TYPE     string
	USERNAME string
	PASSWORD string
	HOSTNAME string
	PORT     string
	SCHEMA   string

	PURGE bool
	DEBUG bool
}

type JWTConfig struct {
	SECRET                string
	SESSION_DURATION_DAYS int
}

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
