package main

import (
	"log"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DatabaseActions interface {
	Setup()
	Purge()
	Migrate()

	LoadTestingData()
	GetTestingUsers() []*User
	GetTestingMovies() []*Movie
	GetTestingDirectors() []*Director
	GetTestingActors() []*Actor

	BeautifyError(error) string
}

type MyDatabase struct {
	*gorm.DB
}

func (database *MyDatabase) Setup() {
	var err error
	database.DB, err = gorm.Open(config.DATABASE.TYPE, config.DATABASE.USERNAME+":"+config.DATABASE.PASSWORD+"@tcp("+config.DATABASE.HOSTNAME+":"+
		config.DATABASE.PORT+")/"+config.DATABASE.SCHEMA+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	if config.DATABASE.PURGE {
		database.Purge()
	}

	database.Migrate()
}

func (database *MyDatabase) Purge() {
	database.Delete(User{})
	database.Delete(Movie{})
	database.Delete(Director{})
	database.Delete(Actor{})
	database.Exec("DELETE FROM movie_actors")
}

func (database *MyDatabase) Migrate() {
	database.AutoMigrate(&User{})
	database.AutoMigrate(&Movie{})
	database.AutoMigrate(&Director{})
	database.AutoMigrate(&Actor{})
}

func (database *MyDatabase) LoadTestingData() {
	for i := 1; i <= 3; i++ {
		database.Create(&User{
			Username: "testing username " + strconv.Itoa(i),
			Email:    "testing email " + strconv.Itoa(i),
			Password: "testing password " + strconv.Itoa(i),
		})

		database.Create(&Movie{
			Name:   "testing name " + strconv.Itoa(i),
			Year:   i,
			Rating: float32(i),
			Director: &Director{
				Name: "testing name " + strconv.Itoa(i),
			},
			Actors: []*Actor{{Name: "testing name " + strconv.Itoa(i)}},
		})
	}
}

func (database *MyDatabase) GetTestingUsers() []*User {
	var users []*User
	database.Find(&users)
	return users
}

func (database *MyDatabase) GetTestingMovies() []*Movie {
	var movies []*Movie
	database.Find(&movies)
	return movies
}

func (database *MyDatabase) GetTestingDirectors() []*Director {
	var directors []*Director
	database.Find(&directors)
	return directors
}

func (database *MyDatabase) GetTestingActors() []*Actor {
	var actors []*Actor
	database.Find(&actors)
	return actors
}

func (database *MyDatabase) BeautifyError(err error) string {
	s := err.Error()

	if strings.Contains(s, "Duplicate entry") {
		duplicateField := strings.Split(s, "'")[3]
		return duplicateField + " already in use"
	}

	return s
}
