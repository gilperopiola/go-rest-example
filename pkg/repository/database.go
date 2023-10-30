package repository

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/config"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type database struct {
	db *gorm.DB
}

func NewDatabase(config *config.Config) database {
	var database database
	database.setup(config)
	return database
}

func (database *database) DB() *gorm.DB {
	return database.db
}

func (database *database) Close() {
	database.db.Close()
}

func (database *database) setup(config *config.Config) {

	// Create connection. It's deferred closed in main.go.
	// Retry connection if it fails due to Docker's orchestration.
	if err := database.connectToDB(config); err != nil {
		log.Fatalf("error connecting to database: %v", err)
		os.Exit(1)
	}

	// Set connection pool limits
	// Log queries if debug = true
	// Destroy or clean tables
	// AutoMigrate fields
	// Create admin
	database.configure(config)
}

func (database *database) connectToDB(config *config.Config) error {
	var err error
	retries := 0
	dbConfig := config.Database

	// Retry connection if it fails due to Docker's orchestration
	for retries < dbConfig.MaxRetries {
		if database.db, err = gorm.Open(dbConfig.Type, dbConfig.GetConnectionString()); err != nil {
			retries++
			if retries == dbConfig.MaxRetries {
				return fmt.Errorf("error connecting to database after %d retries: %v", dbConfig.MaxRetries, err)
			}
			fmt.Printf("error connecting to database, retrying... ")
			time.Sleep(time.Duration(dbConfig.RetryDelay) * time.Second)
			continue
		}
		break
	}
	return nil
}

func (database *database) configure(config *config.Config) {
	dbConfig := config.Database

	// Set connection pool limits
	database.db.DB().SetMaxIdleConns(dbConfig.MaxIdleConns)
	database.db.DB().SetMaxOpenConns(dbConfig.MaxOpenConns)
	database.db.DB().SetConnMaxLifetime(time.Hour)

	// Log queries if debug = true
	database.db.LogMode(dbConfig.Debug)

	// Destroy or clean tables
	if dbConfig.Destroy {
		for _, model := range models.AllModels {
			database.db.DropTable(model)
		}
	} else if dbConfig.Purge {
		for _, model := range models.AllModels {
			database.db.Delete(model)
		}
	}

	// AutoMigrate fields
	database.db.AutoMigrate(models.AllModels...)

	// Insert admin user
	if dbConfig.AdminInsert {
		admin := &models.User{
			Username: "admin",
			Email:    "ferra.main@gmail.com",
			Password: common.Hash(dbConfig.AdminPassword, config.JWT.HashSalt),
			IsAdmin:  true,
		}
		if err := database.DB().Create(admin).Error; err != nil {
			fmt.Print(err.Error())
		}
	}
}
