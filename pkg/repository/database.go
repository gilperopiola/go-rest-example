package repository

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/common/config"
	"github.com/gilperopiola/go-rest-example/pkg/common/middleware"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type database struct {
	db DB
}

// Gorm is quite difficult to mock, not completed
type DB interface {
	Create(value interface{}) *gorm.DB
	Preload(column string, conditions ...interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	Model(value interface{}) *gorm.DB
	Update(attrs ...interface{}) *gorm.DB
	Delete(value interface{}, where ...interface{}) *gorm.DB
	Offset(offset interface{}) *gorm.DB
	Limit(limit interface{}) *gorm.DB
	Close() error
	LogMode(enable bool) *gorm.DB
	DB() *sql.DB
	AutoMigrate(values ...interface{}) *gorm.DB
	DropTable(values ...interface{}) *gorm.DB
}

func NewDatabase(config config.Database, logger middleware.LoggerI) database {
	var database database
	database.setup(config, logger)
	return database
}

const (
	maxRetries = 5
	retryDelay = 5 // In seconds
)

func (database *database) DB() DB {
	return database.db
}

func (database *database) setup(config config.Database, logger middleware.LoggerI) {

	// Create connection. It's deferred closed in main.go.
	// Retry connection if it fails due to Docker's orchestration.
	if err := database.connectToDB(config, logger); err != nil {
		logger.Fatalf("error connecting to database: %v", err)
		os.Exit(1)
	}

	// Set connection pool limits
	// Log queries if debug = true
	// Destroy or clean tables
	// AutoMigrate fields
	database.configure(config)
}

func (database *database) Close() {
	database.db.Close()
}

func (database *database) connectToDB(config config.Database, logger middleware.LoggerI) error {
	var err error
	retries := 0

	// Retry connection if it fails due to Docker's orchestration
	for retries < maxRetries {
		if database.db, err = gorm.Open(config.Type, config.GetConnectionString()); err != nil {
			retries++
			if retries == maxRetries {
				return fmt.Errorf("error connecting to database after %d retries: %v", maxRetries, err)
			}
			logger.Warn("error connecting to database, retrying... ")
			time.Sleep(retryDelay * time.Second)
			continue
		}
		break
	}
	return nil
}

func (database *database) configure(config config.Database) {

	// Set connection pool limits
	database.db.DB().SetMaxIdleConns(10)
	database.db.DB().SetMaxOpenConns(100)
	database.db.DB().SetConnMaxLifetime(time.Hour)

	// Log queries if debug = true
	if config.Debug {
		database.db.LogMode(true)
	}

	// Destroy or clean tables
	if config.Destroy {
		database.db.DropTable(&models.User{})
		database.db.DropTable(&models.UserDetail{})
		database.db.DropTable(&models.UserPost{})
	} else if config.Purge {
		database.db.Delete(models.User{})
		database.db.Delete(models.UserDetail{})
		database.db.Delete(models.UserPost{})
	}

	// AutoMigrate fields
	database.db.AutoMigrate(models.AllModels...)
}
