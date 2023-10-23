package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	General    General
	Database   Database
	JWT        JWT
	Monitoring Monitoring
}

func New(envFilename string) *Config {
	config := Config{}
	config.setup(envFilename)
	return &config
}

func isLastThreeCmd() bool {
	dir, _ := os.Getwd()

	// Extract the current directory name from the path
	dirName := strings.Split(dir, string(os.PathSeparator))
	currentDir := dirName[len(dirName)-1]

	// Check if the last 3 letters are "cmd"
	return strings.HasSuffix(currentDir, "cmd")
}

func (config *Config) setup(envFilename string) {

	envFilePath := ".env"
	if isLastThreeCmd() {
		envFilePath = "../.env"
	}

	// Load .env file
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	// Parse the environment variables into the Config struct
	err = envconfig.Process("", config)
	if err != nil {
		log.Fatalf("error parsing environment variables: %v", err)
	}
}
