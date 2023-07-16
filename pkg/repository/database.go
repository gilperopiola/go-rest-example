package repository

import (
	"log"

	"github.com/gilperopiola/go-rest-example/pkg/config"
	"github.com/gilperopiola/go-rest-example/pkg/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Database struct {
	DB *gorm.DB
}

type DatabaseIFace interface {
	Setup(config config.DatabaseConfig)
	Purge()
	Migrate()
	Close()
}

func NewDatabase(config config.DatabaseConfig) Database {
	var database Database
	database.Setup(config)
	return database
}

/* ------------------- */

func (database *Database) Setup(config config.DatabaseConfig) {
	var err error
	if database.DB, err = gorm.Open(config.TYPE, config.GetConnectionString()); err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	if config.DEBUG {
		database.DB.LogMode(true)
	}

	if config.PURGE {
		database.Purge()
	}

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
