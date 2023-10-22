package config

import "fmt"

type General struct {
	Debug   bool   `envconfig:"GO_REST_EXAMPLE_DEBUG"`
	Port    string `envconfig:"GO_REST_EXAMPLE_PORT"`
	Timeout int    `envconfig:"GO_REST_EXAMPLE_TIMEOUT_SECONDS"`
}

type Database struct {
	Type     string `envconfig:"GO_REST_EXAMPLE_DATABASE_TYPE"`
	Username string `envconfig:"GO_REST_EXAMPLE_DATABASE_USERNAME"`
	Password string `envconfig:"GO_REST_EXAMPLE_DATABASE_PASSWORD"`
	Hostname string `envconfig:"GO_REST_EXAMPLE_DATABASE_HOSTNAME"`
	Port     string `envconfig:"GO_REST_EXAMPLE_DATABASE_PORT"`
	Schema   string `envconfig:"GO_REST_EXAMPLE_DATABASE_SCHEMA"`

	Purge   bool `envconfig:"GO_REST_EXAMPLE_DATABASE_PURGE"`
	Debug   bool `envconfig:"GO_REST_EXAMPLE_DATABASE_DEBUG"`
	Destroy bool `envconfig:"GO_REST_EXAMPLE_DATABASE_DESTROY"`
}

func (config *Database) GetConnectionString() string {
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

type JWT struct {
	Secret              string `envconfig:"GO_REST_EXAMPLE_JWT_SECRET"`
	SessionDurationDays int    `envconfig:"GO_REST_EXAMPLE_JWT_SESSION_DURATION_DAYS"`
	HashSalt            string `envconfig:"GO_REST_EXAMPLE_JWT_HASH_SALT"`
}

type Monitoring struct {
	Enabled bool   `envconfig:"GO_REST_EXAMPLE_MONITORING_ENABLED"`
	AppName string `envconfig:"GO_REST_EXAMPLE_MONITORING_APP_NAME"`
	Secret  string `envconfig:"GO_REST_EXAMPLE_MONITORING_SECRET"`
}
