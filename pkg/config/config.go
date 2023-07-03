package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Configurer interface {
	Setup()
}

type Config struct {
	PORT     string
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
	SECRET                   string
	SESSION_DURATION_SECONDS int
}

/* ------------------- */

func (config *Config) Setup() {
	viper.SetConfigName("config") // configuration file name without the .json or .yaml extension
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("error unmarshaling config: %v", err)
	}

	os.Setenv("PORT", config.PORT)
}

func (config *DatabaseConfig) GetConnectionString() string {
	return config.USERNAME + ":" + config.PASSWORD + "@tcp(" + config.HOSTNAME + ":" +
		config.PORT + ")/" + config.SCHEMA + "?charset=utf8&parseTime=True&loc=Local"
}
