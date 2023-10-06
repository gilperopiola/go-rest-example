package repository

import (
	"os"
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/config"
	"github.com/gilperopiola/go-rest-example/pkg/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type DatabaseInterface interface {
	Setup(config config.DatabaseConfig, logger *logrus.Logger)
	Purge()
	Migrate()
	Close()
}

type Database struct {
	DB *gorm.DB
}

func NewDatabase(config config.DatabaseConfig, logger *logrus.Logger) Database {
	var database Database
	database.Setup(config, logger)
	return database
}

func (database *Database) Setup(config config.DatabaseConfig, logger *logrus.Logger) {

	// Create connection
	var err error
	if database.DB, err = gorm.Open(config.TYPE, config.GetConnectionString()); err != nil {
		logger.Fatalf("error connecting to database: %v", err)
		os.Exit(1)
	}

	// Set connection pool limits
	database.DB.DB().SetMaxIdleConns(10)
	database.DB.DB().SetMaxOpenConns(100)
	database.DB.DB().SetConnMaxLifetime(time.Hour)

	// Log queries
	if config.DEBUG {
		database.DB.LogMode(true)
	}

	// Clean database
	if config.PURGE {
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
