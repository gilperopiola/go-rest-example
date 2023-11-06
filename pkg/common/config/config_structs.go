package config

import "fmt"

// If you add something here, remember to add it to the .env_example file

type General struct {
	AppName        string `envconfig:"GO_REST_EXAMPLE_APP_NAME"`
	Debug          bool   `envconfig:"GO_REST_EXAMPLE_DEBUG"`
	LogInfo        bool   `envconfig:"GO_REST_EXAMPLE_LOG_INFO"`
	Port           string `envconfig:"GO_REST_EXAMPLE_PORT"`
	Profiling      bool   `envconfig:"GO_REST_EXAMPLE_PROFILING"`
	Timeout        int    `envconfig:"GO_REST_EXAMPLE_TIMEOUT_SECONDS"`
	RateLimiterRPS int    `envconfig:"GO_REST_EXAMPLE_RATE_LIMITER_RPS"`
}

type Database struct {
	Type     string `envconfig:"GO_REST_EXAMPLE_DATABASE_TYPE"`
	Username string `envconfig:"GO_REST_EXAMPLE_DATABASE_USERNAME"`
	Password string `envconfig:"GO_REST_EXAMPLE_DATABASE_PASSWORD"`
	Hostname string `envconfig:"GO_REST_EXAMPLE_DATABASE_HOSTNAME"`
	Port     string `envconfig:"GO_REST_EXAMPLE_DATABASE_PORT"`
	Schema   string `envconfig:"GO_REST_EXAMPLE_DATABASE_SCHEMA"`

	Clean         bool   `envconfig:"GO_REST_EXAMPLE_DATABASE_CLEAN"`
	Destroy       bool   `envconfig:"GO_REST_EXAMPLE_DATABASE_DESTROY"`
	AdminInsert   bool   `envconfig:"GO_REST_EXAMPLE_DATABASE_ADMIN_INSERT"`
	AdminPassword string `envconfig:"GO_REST_EXAMPLE_DATABASE_ADMIN_PASSWORD"`

	MaxIdleConns int `envconfig:"GO_REST_EXAMPLE_DATABASE_MAX_IDLE_CONNS"`
	MaxOpenConns int `envconfig:"GO_REST_EXAMPLE_DATABASE_MAX_OPEN_CONNS"`
	MaxRetries   int `envconfig:"GO_REST_EXAMPLE_DATABASE_MAX_RETRIES"`
	RetryDelay   int `envconfig:"GO_REST_EXAMPLE_DATABASE_RETRY_DELAY_SECONDS"`
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

type Auth struct {
	JWTSecret           string `envconfig:"GO_REST_EXAMPLE_AUTH_JWT_SECRET"`
	SessionDurationDays int    `envconfig:"GO_REST_EXAMPLE_AUTH_SESSION_DURATION_DAYS"`
	HashSalt            string `envconfig:"GO_REST_EXAMPLE_AUTH_HASH_SALT"`
}

type Monitoring struct {
	NewRelicEnabled    bool   `envconfig:"GO_REST_EXAMPLE_MONITORING_NEW_RELIC_ENABLED"`
	NewRelicAppName    string `envconfig:"GO_REST_EXAMPLE_MONITORING_NEW_RELIC_APP_NAME"`
	NewRelicLicenseKey string `envconfig:"GO_REST_EXAMPLE_MONITORING_NEW_RELIC_LICENSE_KEY"`

	PrometheusEnabled bool   `envconfig:"GO_REST_EXAMPLE_MONITORING_PROMETHEUS_ENABLED"`
	PrometheusAppName string `envconfig:"GO_REST_EXAMPLE_MONITORING_PROMETHEUS_APP_NAME"`
}
