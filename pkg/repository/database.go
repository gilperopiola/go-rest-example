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

const (
	maxRetries = 5
	retryDelay = 5 // In seconds
)

func (database *Database) Setup(config config.DatabaseConfig, logger logger.LoggerI) {

	// Create connection. It's deferred closed in main.go.
	// Retry connection if it fails due to Docker's orchestration.
	var err error
	retries := 0
	for retries < maxRetries {
		if database.DB, err = gorm.Open(config.Type, config.GetConnectionString()); err != nil {
			retries++
			if retries == maxRetries {
				logger.Fatalf("error connecting to database after %d retries: %v", maxRetries, err)
				os.Exit(1)
			}
			logger.Warn("error connecting to database, retrying... ")
			time.Sleep(retryDelay * time.Second)
			continue
		}
		break
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
