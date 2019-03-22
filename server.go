package main

import (
	"log"
	"os"
)

var config MyConfig
var database MyDatabase
var router MyRouter

func main() {
	config.Setup()
	database.Setup()
	defer database.Close()
	router.Setup()

	log.Println("server started")
	router.Run(":" + os.Getenv("PORT"))
}
