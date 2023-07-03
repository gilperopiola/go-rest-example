package repository

import (
	"log"
	"strconv"

	config "github.com/gilperopiola/go-rest-example/pkg/config"
	models "github.com/gilperopiola/go-rest-example/pkg/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Databaser interface {
	Setup(config config.DatabaseConfig)
	Purge()
	Migrate()
	Close()

	LoadTestingData()
	GetTestingUsers() []*models.User
}

type Database struct {
	DB *gorm.DB
}

func (database *Database) Setup(config config.DatabaseConfig) {
	var err error
	if database.DB, err = gorm.Open(config.TYPE, config.GetConnectionString()); err != nil {
		log.Fatalf("error connecting to database: %v", err)
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

func (database *Database) LoadTestingData() {
	for i := 1; i <= 3; i++ {
		database.DB.Create(&models.User{
			Username: "testing username " + strconv.Itoa(i),
			Email:    "testing email " + strconv.Itoa(i),
			Password: "testing password " + strconv.Itoa(i),
		})
	}
}

func (database *Database) GetTestingUsers() []*models.User {
	var users []*models.User
	database.DB.Find(&users)
	return users
}
