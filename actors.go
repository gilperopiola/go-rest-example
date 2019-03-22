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
	GetJSONBody() string
}

type Actor struct {
	ID          uint      `json:"id" gorm:"auto_increment;unique;not null"`
	Name        string    `json:"name" gorm:"not null"`
	Active      bool      `json:"active" gorm:"default: 1"`
	DateCreated time.Time `json:"date_created" gorm:"default: current_timestamp"`

	Movies []*Movie `json:"movies,omitempty" gorm:"many2many:movie_actors" database:"-"`
}

func CreateActor(c *gin.Context) {
	var actor Actor
	c.BindJSON(&actor)

	if err := actor.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := database.Create(&actor).Error; err != nil {
		c.JSON(http.StatusBadRequest, database.BeautifyError(err))
		return
	}

	actor.Movies = actor.GetMovies()
	c.JSON(http.StatusOK, actor)
}

func ReadActor(c *gin.Context) {
	id := c.Param("id")
	var actor Actor
	if err := database.First(&actor, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	actor.Movies = actor.GetMovies()
	c.JSON(http.StatusOK, actor)
}

func ReadActors(c *gin.Context) {
	var actors []Actor

	id, name, limit, offset, sortField, sortDir := getReadActorsParameters(c)

	tempDatabase := database

	if id > 0 {
		tempDatabase.DB = tempDatabase.Where("id = ?", id)
	}

	tempDatabase.DB = tempDatabase.Where("name LIKE ?", name)

	if sortField != "" && (sortDir == "ASC" || sortDir == "DESC") {
		tempDatabase.DB = tempDatabase.Order(sortField + " " + sortDir)
	}

	tempDatabase.Limit(limit).Offset(offset).Find(&actors)

	for key := range actors {
		actors[key].Movies = actors[key].GetMovies()
	}

	c.JSON(http.StatusOK, actors)
}

func UpdateActor(c *gin.Context) {
	id := c.Param("id")
	var actor Actor

	if err := database.First(&actor, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.BindJSON(&actor)
	if err := database.Save(&actor).Error; err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, database.BeautifyError(err))
		return
	}

	actor.Movies = actor.GetMovies()
	c.JSON(http.StatusOK, actor)
}

//extra
func (actor *Actor) GetMovies() []*Movie {
	var movies []*Movie
	database.Model(&actor).Association("Movies").Find(&movies)
	return movies
}

func (actor *Actor) Validate() error {
	if len(actor.Name) == 0 {
		return errors.New("name required")
	}
	return nil
}

func (actor *Actor) GetJSONBody() string {
	body := `{
		"name": "` + actor.Name + `",
		"active": ` + strconv.FormatBool(actor.Active) + `
	}`

	return body
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
