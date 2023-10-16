package config

import "fmt"

type GeneralConfig struct {
	Debug   bool
	Port    string
	Timeout int
}

type DatabaseConfig struct {
	Type     string
	Username string
	Password string
	Hostname string
	Port     string
	Schema   string

	Purge bool
	Debug bool
}

type JWTConfig struct {
	Secret              string
	SessionDurationDays int
}

type MonitoringConfig struct {
	Enabled bool
	AppName string
	Secret  string
}

func (config *Config) General() GeneralConfig {
	return config.general
}

func (config *Config) Database() DatabaseConfig {
	return config.database
}

func (config *Config) JWT() JWTConfig {
	return config.jwt
}

func (config *Config) Monitoring() MonitoringConfig {
	return config.monitoring
}

func (config *DatabaseConfig) GetConnectionString() string {

	var (
		username = config.Username
		password = config.Password
		hostname = config.Hostname
		port     = config.Port
		schema   = config.Schema
		params   = "?charset=utf8&parseTime=True&loc=Local"
	)

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s%s",
		username, password, hostname, port, schema, params,
	)
}
