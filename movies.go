package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type MovieActions interface {
	GetDirector() *Director
	GetActors() []*Actor

	Validate() error
}

type Movie struct {
	ID          uint      `json:"id" gorm:"auto_increment;unique;not null"`
	Name        string    `json:"name" gorm:"not null"`
	Year        int       `json:"year"`
	Rating      float32   `json:"rating" gorm:"default: 0"`
	DirectorID  int       `json:"director_id,omitempty"`
	Active      bool      `json:"active" gorm:"default: 1"`
	DateCreated time.Time `json:"date_created" gorm:"default: current_timestamp"`

	Director *Director `json:"director,omitempty" db:"-"`
	Actors   []*Actor  `json:"actors" gorm:"many2many:movie_actors" db:"-"`
}

func CreateMovie(c *gin.Context) {
	var movie Movie
	c.BindJSON(&movie)

	if err := movie.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Create(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, beautifyDatabaseError(err))
		return
	}

	movie.Director = movie.GetDirector()
	movie.DirectorID = 0
	movie.Actors = movie.GetActors()
	c.JSON(http.StatusOK, movie)
}

func ReadMovie(c *gin.Context) {
	id := c.Param("id")
	var movie Movie
	if err := db.First(&movie, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	movie.Director = movie.GetDirector()
	movie.DirectorID = 0
	movie.Actors = movie.GetActors()
	c.JSON(http.StatusOK, movie)
}

func ReadMovies(c *gin.Context) {
	var movies []Movie

	id, name, limit, offset, sortField, sortDir := getReadMoviesParameters(c)

	tempDB := db

	if id > 0 {
		tempDB = tempDB.Where("id = ?", id)
	}

	tempDB = tempDB.Where("name LIKE ?", name)

	if sortField != "" && (sortDir == "ASC" || sortDir == "DESC") {
		tempDB = tempDB.Order(sortField + " " + sortDir)
	}

	tempDB.Limit(limit).Offset(offset).Find(&movies)

	for key := range movies {
		movies[key].Director = movies[key].GetDirector()
		movies[key].DirectorID = 0
		movies[key].Actors = movies[key].GetActors()
	}

	c.JSON(http.StatusOK, movies)
}

func UpdateMovie(c *gin.Context) {
	id := c.Param("id")
	var movie Movie

	if err := db.First(&movie, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.BindJSON(&movie)
	if err := db.Save(&movie).Error; err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, beautifyDatabaseError(err))
		return
	}

	movie.Director = movie.GetDirector()
	movie.DirectorID = 0
	movie.Actors = movie.GetActors()
	c.JSON(http.StatusOK, movie)
}

//extra
func (movie *Movie) GetDirector() *Director {
	var director Director
	db.Where("id = ?", movie.DirectorID).Find(&director)
	return &director
}

func (movie *Movie) GetActors() []*Actor {
	var actors []*Actor
	db.Model(&movie).Association("Actors").Find(&actors)
	return actors
}

func (movie *Movie) Validate() error {
	if len(movie.Name) == 0 || movie.DirectorID < 1 {
		return errors.New("name and director required")
	}

	return nil
}

func getReadMoviesParameters(c *gin.Context) (int, string, string, string, string, string) {
	id, _ := strconv.Atoi(c.Query("ID"))
	name := "%" + c.Query("Name") + "%"
	limit := c.Query("Limit")
	offset := c.Query("Offset")
	sortField := c.Query("SortField")
	sortDir := c.Query("SortDir")
	return id, name, limit, offset, sortField, sortDir
}
