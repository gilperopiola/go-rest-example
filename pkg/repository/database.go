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

	// Clean tables
	if config.Purge {
		database.DB.Delete(models.User{})
	}

	// Destroy tables
	if config.Destroy {
		database.DB.DropTable(&models.User{})
	}

	// Run migrations
	database.DB.AutoMigrate(&models.User{})
}

func (database *Database) Close() {
	database.DB.Close()
}
