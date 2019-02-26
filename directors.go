package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type DirectorActions interface {
	GetMovies() *[]Movie

	Validate() error
}

type Director struct {
	ID          uint      `json:"id" gorm:"auto_increment;unique;not null"`
	Name        string    `json:"name" gorm:"unique;not null"`
	Active      bool      `json:"active" gorm:"default: 1"`
	DateCreated time.Time `json:"date_created" gorm:"default: current_timestamp"`

	Movies *[]Movie `json:"movies,omitempty" gorm:"-"`
}

func CreateDirector(c *gin.Context) {
	var director Director
	c.BindJSON(&director)

	if err := director.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Create(&director).Error; err != nil {
		c.JSON(http.StatusBadRequest, beautifyDatabaseError(err))
		return
	}

	director.Movies = &[]Movie{}
	c.JSON(http.StatusOK, director)
}

func ReadDirector(c *gin.Context) {
	id := c.Param("id")
	var director Director
	if err := db.First(&director, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	director.Movies = director.GetMovies()
	c.JSON(http.StatusOK, director)
}

func ReadDirectors(c *gin.Context) {
	var directors []Director

	id, name, limit, offset, sortField, sortDir := getReadDirectorsParameters(c)

	tempDB := db

	if id > 0 {
		tempDB = tempDB.Where("id = ?", id)
	}

	tempDB = tempDB.Where("name LIKE ?", name)

	if sortField != "" && (sortDir == "ASC" || sortDir == "DESC") {
		tempDB = tempDB.Order(sortField + " " + sortDir)
	}

	tempDB.Limit(limit).Offset(offset).Find(&directors)

	for key := range directors {
		log.Printf("%v", directors[key])
		directors[key].Movies = directors[key].GetMovies()
	}

	c.JSON(http.StatusOK, directors)
}

func UpdateDirector(c *gin.Context) {
	id := c.Param("id")
	var director Director
	if err := db.First(&director, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.BindJSON(&director)
	if err := db.Save(&director).Error; err != nil {
		c.JSON(http.StatusBadRequest, beautifyDatabaseError(err))
		return
	}

	director.Movies = director.GetMovies()
	c.JSON(http.StatusOK, director)
}

//extra
func (director *Director) GetMovies() *[]Movie {
	var movies []Movie
	db.Where("director_id = ?", director.ID).Find(&movies)

	for key := range movies {
		movies[key].DirectorID = 0
	}

	return &movies
}

func (director *Director) Validate() error {
	if len(director.Name) == 0 {
		return errors.New("name required")
	}
	return nil
}

func getReadDirectorsParameters(c *gin.Context) (int, string, string, string, string, string) {
	id, _ := strconv.Atoi(c.Query("ID"))
	name := "%" + c.Query("Name") + "%"
	limit := c.Query("Limit")
	offset := c.Query("Offset")
	sortField := c.Query("SortField")
	sortDir := c.Query("SortDir")
	return id, name, limit, offset, sortField, sortDir
}
