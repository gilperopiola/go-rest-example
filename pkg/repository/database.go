package repository

import (
	"os"
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/common/config"
	"github.com/gilperopiola/go-rest-example/pkg/common/logger"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DatabaseInterface interface {
	Setup(config config.DatabaseConfig, logger logger.LoggerI)
	Purge()
	Migrate()
	Close()
}

type Database struct {
	DB *gorm.DB
}

func NewDatabase(config config.DatabaseConfig, logger logger.LoggerI) Database {
	var database Database
	database.Setup(config, logger)
	return database
}

func (database *Database) Setup(config config.DatabaseConfig, logger logger.LoggerI) {

	// Create connection. It's deferred closed in main.go
	var err error
	if database.DB, err = gorm.Open(config.Type, config.GetConnectionString()); err != nil {
		logger.Fatalf("error connecting to database: %v", err)
		os.Exit(1)
	}

	// Set connection pool limits
	database.DB.DB().SetMaxIdleConns(10)
	database.DB.DB().SetMaxOpenConns(100)
	database.DB.DB().SetConnMaxLifetime(time.Hour)

	// Log queries
	if config.Debug {
		database.DB.LogMode(true)
	}

	// Clean database
	if config.Purge {
		database.Purge()
	}

	// Run migrations
	database.Migrate()
}

func (database *Database) Purge() {
	database.DB.Delete(models.User{})
}

func (database *Database) Migrate() {
	database.DB.AutoMigrate(&models.User{})
}

func (database *Database) Close() {
	database.DB.Close()
}
