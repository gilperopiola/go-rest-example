package config

// Config Defaults - TODO - Move to config file

var (
	defaultPort  = "8040"
	defaultDebug = true

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

	defaultMonitoringEnabled = false
	defaultMonitoringAppName = "go-rest-example"
	defaultMonitoringSecret  = ""
)
