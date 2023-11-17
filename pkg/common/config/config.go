package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	General
	Database   Database
	Monitoring Monitoring
}

func New() *Config {
	config := Config{}
	config.setup()
	return &config
}

func (config *Config) setup() {
	// We may be on the cmd folder or not. Hacky, I know
	envFilePath := ".env"
	if currentFolderIsCMD() {
		envFilePath = "../.env"
	}
	// Load .env file into environment variables
	if err := godotenv.Load(envFilePath); err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}
	// Parse the environment variables into the Config struct
	if err := envconfig.Process("", config); err != nil {
		log.Fatalf("error parsing environment variables: %v", err)
	}
}

func (sqlConfig *SQL) GetMySQLConnectionString() string {
	var (
		username = sqlConfig.Username
		password = sqlConfig.Password
		hostname = sqlConfig.Hostname
		port     = sqlConfig.Port
		schema   = sqlConfig.Schema
		params   = "?charset=utf8&parseTime=True&loc=Local"
	)
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s", username, password, hostname, port, schema, params)
}

func currentFolderIsCMD() bool {
	dir, _ := os.Getwd()
	dirName := strings.Split(dir, string(os.PathSeparator))
	currentDir := dirName[len(dirName)-1]
	return strings.HasSuffix(currentDir, "cmd")
}
