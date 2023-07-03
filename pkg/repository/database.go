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
	Setup(config config.Config)
	Purge()
	Migrate()

	LoadTestingData()
	GetTestingUsers() []*models.User

	BeautifyError(error) string
}

type Database struct {
	*gorm.DB
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
	database.Delete(models.User{})
}

func (database *Database) Migrate() {
	database.AutoMigrate(&models.User{})
}

func (database *Database) LoadTestingData() {
	for i := 1; i <= 3; i++ {
		database.Create(&models.User{
			Username: "testing username " + strconv.Itoa(i),
			Email:    "testing email " + strconv.Itoa(i),
			Password: "testing password " + strconv.Itoa(i),
		})
	}
}

func (database *Database) GetTestingUsers() []*models.User {
	var users []*models.User
	database.Find(&users)
	return users
}
