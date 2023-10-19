package config

import "fmt"

type General struct {
	Debug   bool
	Port    string
	Timeout int
}

type Database struct {
	Type     string
	Username string
	Password string
	Hostname string
	Port     string
	Schema   string

	Purge   bool
	Debug   bool
	Destroy bool
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
	Secret              string
	SessionDurationDays int
	HashSalt            string
}

type Monitoring struct {
	Enabled bool
	AppName string
	Secret  string
}
