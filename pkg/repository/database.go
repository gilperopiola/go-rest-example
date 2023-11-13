package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type database struct {
	db *gorm.DB
}

func NewDatabase() *database {
	var database database

	// Create connection. It's deferred closed in main.go.
	// Retry connection if it fails due to Docker's orchestration.
	if err := database.connectToDB(); err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	// Set connection pool limits
	// Log queries if debug = true
	// Destroy or clean tables
	// AutoMigrate fields
	// Create admin
	database.configure()

	return &database
}

func (database *database) DB() *gorm.DB {
	return database.db
}

func (database *database) connectToDB() error {
	dbConfig := common.Cfg.Database
	retries := 0
	var err error

	// Retry connection if it fails due to Docker's orchestration
	for retries < dbConfig.MaxRetries {
		if database.db, err = gorm.Open(mysql.Open(dbConfig.GetConnectionString()), &gorm.Config{Logger: common.Logger}); err == nil {
			break
		}

		retries++
		if retries >= dbConfig.MaxRetries {
			common.Logger.Error(context.Background(), fmt.Sprintf("error connecting to database after %d retries: %v", dbConfig.MaxRetries, err), nil)
			return err
		}

		common.Logger.Info(context.Background(), "error connecting to database, retrying... ", nil)
		time.Sleep(time.Duration(dbConfig.RetryDelay) * time.Second)
	}
	return nil
}

func (database *database) configure() {
	dbConfig := common.Cfg.Database
	mySQLDB, err := database.db.DB()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

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
		admin := makeAdminModel("ferra.main@gmail.com", common.Hash(dbConfig.AdminPassword, common.Cfg.Auth.HashSalt))
		if err := database.db.Create(admin).Error; err != nil {
			fmt.Println(err.Error())
		}
	}

	// Just for formatting the logs :)
	if common.Cfg.LogInfo {
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
