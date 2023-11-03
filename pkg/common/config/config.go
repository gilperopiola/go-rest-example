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
	Auth       Auth
	Monitoring Monitoring
}

func New() *Config {
	config := Config{}
	config.setup()
	return &config
}

func (config *Config) setup() {

	// We may be on the cmd folder or not. Hacky, I know.
	envFilePath := ".env"
	if currentFolderIsCMD() {
		envFilePath = "../.env"
	}

	// Load .env file into environment variables
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

func currentFolderIsCMD() bool {
	dir, _ := os.Getwd()

	// Extract the current directory name from the path
	dirName := strings.Split(dir, string(os.PathSeparator))
	currentDir := dirName[len(dirName)-1]

	// Check if the last 3 letters are "cmd"
	return strings.HasSuffix(currentDir, "cmd")
}
