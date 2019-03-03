package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ActorActions interface {
	GetMovies() []*Movie

	Validate() error
}

type Actor struct {
	ID          uint      `json:"id" gorm:"auto_increment;unique;not null"`
	Name        string    `json:"name" gorm:"not null"`
	Active      bool      `json:"active" gorm:"default: 1"`
	DateCreated time.Time `json:"date_created" gorm:"default: current_timestamp"`

	Movies []*Movie `json:"movies,omitempty" gorm:"many2many:movie_actors" db:"-"`
}

func CreateActor(c *gin.Context) {
	var actor Actor
	c.BindJSON(&actor)

	if err := actor.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Create(&actor).Error; err != nil {
		c.JSON(http.StatusBadRequest, beautifyDatabaseError(err))
		return
	}

	actor.Movies = []*Movie{}
	c.JSON(http.StatusOK, actor)
}

func ReadActor(c *gin.Context) {
	id := c.Param("id")
	var actor Actor
	if err := db.First(&actor, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	actor.Movies = actor.GetMovies()
	c.JSON(http.StatusOK, actor)
}

func ReadActors(c *gin.Context) {
	var actors []Actor

	id, name, limit, offset, sortField, sortDir := getReadActorsParameters(c)

	tempDB := db

	if id > 0 {
		tempDB = tempDB.Where("id = ?", id)
	}

	tempDB = tempDB.Where("name LIKE ?", name)

	if sortField != "" && (sortDir == "ASC" || sortDir == "DESC") {
		tempDB = tempDB.Order(sortField + " " + sortDir)
	}

	tempDB.Limit(limit).Offset(offset).Find(&actors)

	for key := range actors {
		actors[key].Movies = actors[key].GetMovies()
	}

	c.JSON(http.StatusOK, actors)
}

func UpdateActor(c *gin.Context) {
	id := c.Param("id")
	var actor Actor

	if err := db.First(&actor, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.BindJSON(&actor)
	if err := db.Save(&actor).Error; err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, beautifyDatabaseError(err))
		return
	}

	actor.Movies = actor.GetMovies()
	c.JSON(http.StatusOK, actor)
}

//extra
func (actor *Actor) GetMovies() []*Movie {
	var movies []*Movie
	db.Model(&actor).Association("Movies").Find(&movies)
	return movies
}

func (actor *Actor) Validate() error {
	if len(actor.Name) == 0 {
		return errors.New("name required")
	}
	return nil
}

func getReadActorsParameters(c *gin.Context) (int, string, string, string, string, string) {
	id, _ := strconv.Atoi(c.Query("ID"))
	name := "%" + c.Query("Name") + "%"
	limit := c.Query("Limit")
	offset := c.Query("Offset")
	sortField := c.Query("SortField")
	sortDir := c.Query("SortDir")
	return id, name, limit, offset, sortField, sortDir
}
