package main

import (
	"log"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func setupDatabase() {
	var err error
	db, err = gorm.Open(config.DATABASE.TYPE, config.DATABASE.USERNAME+":"+config.DATABASE.PASSWORD+"@tcp("+config.DATABASE.HOSTNAME+":"+
		config.DATABASE.PORT+")/"+config.DATABASE.SCHEMA+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	if config.DATABASE.PURGE {
		purgeDatabase()
	}

	migrateDatabase()
}

func purgeDatabase() {
	db.Delete(User{})
	db.Delete(Movie{})
	db.Delete(Director{})
	db.Delete(Actor{})
	db.Exec("DELETE FROM movie_actors")
}

func migrateDatabase() {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Movie{})
	db.AutoMigrate(&Director{})
	db.AutoMigrate(&Actor{})
}

func loadTestingData() {
	for i := 1; i <= 3; i++ {
		db.Create(&User{
			Username: "username " + strconv.Itoa(i),
			Email:    "email " + strconv.Itoa(i),
			Password: "password " + strconv.Itoa(i),
		})

		db.Create(&Movie{
			Name:   "name " + strconv.Itoa(i),
			Year:   i,
			Rating: float32(i),
			Director: &Director{
				Name: "name " + strconv.Itoa(i),
			},
			Actors: []*Actor{{Name: "name " + strconv.Itoa(i) + "a"}, {Name: "name " + strconv.Itoa(i) + "b"}},
		})
	}
}

func beautifyDatabaseError(err error) string {
	s := err.Error()

	if strings.Contains(s, "Duplicate entry") {
		duplicateField := strings.Split(s, "'")[3]
		return duplicateField + " already in use"
	}

	return s
}
