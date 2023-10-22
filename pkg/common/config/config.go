package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	General    General
	Database   Database
	JWT        JWT
	Monitoring Monitoring
}

func New() *Config {
	config := Config{}
	config.setup()
	return &config
}

func (config *Config) setup() {

	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	// Parse the environment variables into the Config struct
	err = envconfig.Process("", config)
	if err != nil {
		log.Fatalf("error parsing environment variables: %v", err)
	}
}
