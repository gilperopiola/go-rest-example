package main

import (
	"errors"
	"fmt"
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
	GetJSONBody() string
}

type Movie struct {
	ID          uint      `json:"id" gorm:"auto_increment;unique;not null"`
	Name        string    `json:"name" gorm:"unique;not null"`
	Year        int       `json:"year"`
	Rating      float32   `json:"rating" gorm:"default: 0"`
	DirectorID  int       `json:"director_id,omitempty"`
	Active      bool      `json:"active" gorm:"default: 1"`
	DateCreated time.Time `json:"date_created" gorm:"default: current_timestamp"`

	Director *Director `json:"director,omitempty" database:"-"`
	Actors   []*Actor  `json:"actors" gorm:"many2many:movie_actors" database:"-"`
}

func CreateMovie(c *gin.Context) {
	var movie Movie
	c.BindJSON(&movie)

	if err := movie.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := database.Create(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, database.BeautifyError(err))
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
	if err := database.First(&movie, id).Error; err != nil {
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

	tempDatabase := database

	if id > 0 {
		tempDatabase.DB = tempDatabase.Where("id = ?", id)
	}

	tempDatabase.DB = tempDatabase.Where("name LIKE ?", name)

	if sortField != "" && (sortDir == "ASC" || sortDir == "DESC") {
		tempDatabase.DB = tempDatabase.Order(sortField + " " + sortDir)
	}

	tempDatabase.Limit(limit).Offset(offset).Find(&movies)

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

	if err := database.First(&movie, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	log.Printf("%v", movie)

	c.BindJSON(&movie)
	if err := database.Save(&movie).Error; err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, database.BeautifyError(err))
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
	database.Where("id = ?", movie.DirectorID).Find(&director)
	return &director
}

func (movie *Movie) GetActors() []*Actor {
	var actors []*Actor
	database.Model(&movie).Association("Actors").Find(&actors)
	return actors
}

func (movie *Movie) Validate() error {
	if len(movie.Name) == 0 || movie.DirectorID < 1 {
		return errors.New("name and director required")
	}

	return nil
}

func (movie *Movie) GetJSONBody() string {
	body := `{
		"name": "` + movie.Name + `",
		"year": ` + strconv.Itoa(movie.Year) + `,
		"rating": ` + fmt.Sprintf("%f", movie.Rating) + `,
		"active": ` + strconv.FormatBool(movie.Active) + `,
		"director_id": ` + strconv.Itoa(movie.DirectorID) + `
	}`

	return body
}

func getReadMoviesParameters(c *gin.Context) (id int, name string, limit, offset, sortField, sortDir string) {
	id, _ = strconv.Atoi(c.Query("ID"))
	name = "%" + c.Query("Name") + "%"
	return id, name, c.Query("Limit"), c.Query("Offset"), c.Query("SortField"), c.Query("SortDir")
}
