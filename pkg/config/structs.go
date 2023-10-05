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
