package main

import (
	"log"
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

	runMigrations()
}

func deleteAllRecords() {
	db.Delete(User{})
	db.Delete(Movie{})
	db.Delete(Director{})
}

func runMigrations() {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Movie{})
	db.AutoMigrate(&Director{})
}

func beautifyDatabaseError(err error) string {
	s := err.Error()

	if strings.Contains(s, "Duplicate entry") {
		duplicateField := strings.Split(s, "'")[3]
		return duplicateField + " already in use"
	}

	return s
}
