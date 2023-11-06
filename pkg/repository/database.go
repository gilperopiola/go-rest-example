package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/config"
	"github.com/gilperopiola/go-rest-example/pkg/common/middleware"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type database struct {
	db *gorm.DB
}

func NewDatabase(config *config.Config, logger *middleware.LoggerAdapter) *database {
	var database database
	database.setup(config, logger)
	return &database
}

func (database *database) DB() *gorm.DB {
	return database.db
}

func (database *database) setup(config *config.Config, logger *middleware.LoggerAdapter) {

	// Create connection. It's deferred closed in main.go.
	// Retry connection if it fails due to Docker's orchestration.
	if err := database.connectToDB(config, logger); err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	// Set connection pool limits
	// Log queries if debug = true
	// Destroy or clean tables
	// AutoMigrate fields
	// Create admin
	database.configure(config)
}

func (database *database) connectToDB(config *config.Config, logger *middleware.LoggerAdapter) error {
	dbConfig := config.Database
	retries := 0
	var err error

	// Retry connection if it fails due to Docker's orchestration
	for retries < dbConfig.MaxRetries {
		if database.db, err = gorm.Open(mysql.Open(dbConfig.GetConnectionString()), &gorm.Config{Logger: logger}); err == nil {
			break
		}

		retries++
		if retries >= dbConfig.MaxRetries {
			logger.Logger.Error(fmt.Sprintf("error connecting to database after %d retries: %v", dbConfig.MaxRetries, err), nil)
			return err
		}

		logger.Logger.Info("error connecting to database, retrying... ", map[string]interface{}{})
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

	// Destroy or clean tables
	if dbConfig.Destroy {
		for _, model := range models.AllModels {
			database.db.Migrator().DropTable(model)
		}
	} else if dbConfig.Clean {
		for _, model := range models.AllModels {
			database.db.Delete(model)
		}
	}

	// AutoMigrate fields
	database.db.AutoMigrate(models.AllModels...)

	// Insert admin user
	if dbConfig.AdminInsert {
		admin := makeAdminModel("ferra.main@gmail.com", common.Hash(dbConfig.AdminPassword, config.Auth.HashSalt))
		if err := database.db.Create(admin).Error; err != nil {
			fmt.Println(err.Error())
		}
	}

	// Just for formatting the logs :)
	if config.General.LogInfo {
		fmt.Println("")
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
