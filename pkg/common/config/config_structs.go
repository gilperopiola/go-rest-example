package config

/*------------------------------------------------------------------------
// If you add something here, remember to add it to the .env_example file
/------------------------------*/

/*--------------------
//      General
//------------------*/

type General struct {
	AppName             string `envconfig:"GO_REST_EXAMPLE_APP_NAME"`              // Name of your app
	Debug               bool   `envconfig:"GO_REST_EXAMPLE_DEBUG"`                 // Enable debug mode
	LogInfo             bool   `envconfig:"GO_REST_EXAMPLE_LOG_INFO"`              // Enable info logs
	Port                string `envconfig:"GO_REST_EXAMPLE_PORT"`                  // Port to run the app on
	Profiling           bool   `envconfig:"GO_REST_EXAMPLE_PROFILING"`             // Enable profiling
	TimeoutSeconds      int    `envconfig:"GO_REST_EXAMPLE_TIMEOUT_SECONDS"`       // Timeout for HTTP requests
	RateLimiterRPS      int    `envconfig:"GO_REST_EXAMPLE_RATE_LIMITER_RPS"`      // Requests per second
	JWTSecret           string `envconfig:"GO_REST_EXAMPLE_JWT_SECRET"`            // JWT secret
	SessionDurationDays int    `envconfig:"GO_REST_EXAMPLE_SESSION_DURATION_DAYS"` // Session duration
	HashSalt            string `envconfig:"GO_REST_EXAMPLE_HASH_SALT"`             // Hash salt
}

/*--------------------
//     Databases
//------------------*/

type Database struct {
	Type  string `envconfig:"GO_REST_EXAMPLE_DATABASE_TYPE"` // Database type: mysql or mongodb
	SQL   SQL
	Mongo Mongo

	Clean         bool   `envconfig:"GO_REST_EXAMPLE_DATABASE_CLEAN"`          // Erase all database records
	Destroy       bool   `envconfig:"GO_REST_EXAMPLE_DATABASE_DESTROY"`        // Destroy database
	AdminInsert   bool   `envconfig:"GO_REST_EXAMPLE_DATABASE_ADMIN_INSERT"`   // Insert admin user
	AdminPassword string `envconfig:"GO_REST_EXAMPLE_DATABASE_ADMIN_PASSWORD"` // Admin password
}

type SQL struct {
	Username string `envconfig:"GO_REST_EXAMPLE_DATABASE_SQL_USERNAME"` // MySQL Username
	Password string `envconfig:"GO_REST_EXAMPLE_DATABASE_SQL_PASSWORD"` // MySQL Password
	Hostname string `envconfig:"GO_REST_EXAMPLE_DATABASE_SQL_HOSTNAME"` // MySQL Hostname
	Port     string `envconfig:"GO_REST_EXAMPLE_DATABASE_SQL_PORT"`     // MySQL Port
	Schema   string `envconfig:"GO_REST_EXAMPLE_DATABASE_SQL_SCHEMA"`   // MySQL Schema

	MaxIdleConns int `envconfig:"GO_REST_EXAMPLE_DATABASE_SQL_MAX_IDLE_CONNS"`      // MySQL Max Idle Connections
	MaxOpenConns int `envconfig:"GO_REST_EXAMPLE_DATABASE_SQL_MAX_OPEN_CONNS"`      // MySQL Max Open Connections
	MaxRetries   int `envconfig:"GO_REST_EXAMPLE_DATABASE_SQL_MAX_RETRIES"`         // MySQL Max Connection Retries
	RetryDelay   int `envconfig:"GO_REST_EXAMPLE_DATABASE_SQL_RETRY_DELAY_SECONDS"` // MySQL Connection Retry Delay
}

type Mongo struct {
	ConnectionString string `envconfig:"GO_REST_EXAMPLE_DATABASE_MONGO_CONNECTION_STRING"` // MongoDB Connection String
	DBName           string `envconfig:"GO_REST_EXAMPLE_DATABASE_MONGO_NAME"`              // MongoDB Database Name
}

/*--------------------
//     Monitoring
//------------------*/

type Monitoring struct {
	NewRelicEnabled    bool   `envconfig:"GO_REST_EXAMPLE_MONITORING_NEW_RELIC_ENABLED"`     // Enable NewRelic
	NewRelicAppName    string `envconfig:"GO_REST_EXAMPLE_MONITORING_NEW_RELIC_APP_NAME"`    // NewRelic App Name
	NewRelicLicenseKey string `envconfig:"GO_REST_EXAMPLE_MONITORING_NEW_RELIC_LICENSE_KEY"` // NewRelic License Key

	PrometheusEnabled bool   `envconfig:"GO_REST_EXAMPLE_MONITORING_PROMETHEUS_ENABLED"`  // Enable Prometheus
	PrometheusAppName string `envconfig:"GO_REST_EXAMPLE_MONITORING_PROMETHEUS_APP_NAME"` // Prometheus App Name
}
