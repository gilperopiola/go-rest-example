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

	// Destroy or clean tables
	if config.Destroy {
		database.DB.DropTable(&models.User{})
		database.DB.DropTable(&models.UserDetail{})
		database.DB.DropTable(&models.UserPost{})
	} else if config.Purge {
		database.DB.Delete(models.User{})
		database.DB.Delete(models.UserDetail{})
		database.DB.Delete(models.UserPost{})
	}

	// Run migrations
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.UserDetail{})
	database.DB.AutoMigrate(&models.UserPost{})
}

func (database *Database) Close() {
	database.DB.Close()
}
