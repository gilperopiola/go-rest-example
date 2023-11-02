package repository

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/config"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type database struct {
	db *gorm.DB
}

func NewDatabase(config *config.Config, logger *DBLogger) database {
	var database database
	database.setup(config, logger)
	return database
}

func (database *database) DB() *gorm.DB {
	return database.db
}

func (database *database) setup(config *config.Config, logger *DBLogger) {

	// Create connection. It's deferred closed in main.go.
	// Retry connection if it fails due to Docker's orchestration.
	if err := database.connectToDB(config, logger); err != nil {
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

func (database *database) connectToDB(config *config.Config, logger *DBLogger) error {
	var err error
	retries := 0
	dbConfig := config.Database

	// Retry connection if it fails due to Docker's orchestration
	for retries < dbConfig.MaxRetries {
		if database.db, err = gorm.Open(mysql.Open(dbConfig.GetConnectionString()), &gorm.Config{Logger: logger}); err == nil {
			break
		}

		retries++
		if retries >= dbConfig.MaxRetries {
			return fmt.Errorf("error connecting to database after %d retries: %v", dbConfig.MaxRetries, err)
		}

		fmt.Println("error connecting to database, retrying... ")
		time.Sleep(time.Duration(dbConfig.RetryDelay) * time.Second)
	}
	return nil
}

func (database *database) configure(config *config.Config) {
	mySQLDB, _ := database.db.DB()
	dbConfig := config.Database

	// Set connection pool limits
	mySQLDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
	mySQLDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
	mySQLDB.SetConnMaxLifetime(time.Hour)

	// Log queries if debug = true
	if dbConfig.Debug {
		database.db.Debug()
	}

	// Destroy or clean tables
	if dbConfig.Destroy {
		for _, model := range models.AllModels {
			database.db.Migrator().DropTable(model)
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
		admin := makeAdminModel("ferra.main@gmail.com", common.Hash(dbConfig.AdminPassword, config.JWT.HashSalt))
		if err := database.DB().Create(admin).Error; err != nil {
			fmt.Println(err.Error())
		}
	}
}

func makeAdminModel(email, password string) *models.User {
	return &models.User{
		Username: "admin",
		Email:    email,
		Password: password,
		IsAdmin:  true,
	}
}
