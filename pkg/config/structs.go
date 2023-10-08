package config

type Config struct {
	PORT  string
	DEBUG bool

	DATABASE   DatabaseConfig
	JWT        JWTConfig
	MONITORING MonitoringConfig
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

type MonitoringConfig struct {
	ENABLED bool
	SECRET  string
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

func (config *Config) GetMonitoringConfig() MonitoringConfig {
	return config.MONITORING
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
