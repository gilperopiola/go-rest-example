package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type DirectorActions interface {
	GetMovies() []*Movie

	Validate() error
	GetJSONBody() string
}

type Director struct {
	ID          uint      `json:"id" gorm:"auto_increment;unique;not null"`
	Name        string    `json:"name" gorm:"unique;not null"`
	Active      bool      `json:"active" gorm:"default: 1"`
	DateCreated time.Time `json:"date_created" gorm:"default: current_timestamp"`

	Movies []*Movie `json:"movies,omitempty" gorm:"-"`
}

func CreateDirector(c *gin.Context) {
	var director Director
	c.BindJSON(&director)

	if err := director.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := database.Create(&director).Error; err != nil {
		c.JSON(http.StatusBadRequest, database.BeautifyError(err))
		return
	}

	director.Movies = []*Movie{}
	c.JSON(http.StatusOK, director)
}

func ReadDirector(c *gin.Context) {
	id := c.Param("id")
	var director Director
	if err := database.First(&director, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	director.Movies = director.GetMovies()
	c.JSON(http.StatusOK, director)
}

func ReadDirectors(c *gin.Context) {
	var directors []Director

	id, name, limit, offset, sortField, sortDir := getReadDirectorsParameters(c)

	tempDatabase := database

	if id > 0 {
		tempDatabase.DB = tempDatabase.Where("id = ?", id)
	}

	tempDatabase.DB = tempDatabase.Where("name LIKE ?", name)

	if sortField != "" && (sortDir == "ASC" || sortDir == "DESC") {
		tempDatabase.DB = tempDatabase.Order(sortField + " " + sortDir)
	}

	tempDatabase.Limit(limit).Offset(offset).Find(&directors)

	for key := range directors {
		directors[key].Movies = directors[key].GetMovies()
	}

	c.JSON(http.StatusOK, directors)
}

func UpdateDirector(c *gin.Context) {
	id := c.Param("id")
	var director Director
	if err := database.First(&director, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.BindJSON(&director)
	if err := database.Save(&director).Error; err != nil {
		c.JSON(http.StatusBadRequest, database.BeautifyError(err))
		return
	}

	director.Movies = director.GetMovies()
	c.JSON(http.StatusOK, director)
}

//extra
func (director *Director) GetMovies() []*Movie {
	var movies []*Movie
	database.Where("director_id = ?", director.ID).Find(&movies)

	for key := range movies {
		movies[key].DirectorID = 0
	}

	return movies
}

func (director *Director) Validate() error {
	if len(director.Name) == 0 {
		return errors.New("name required")
	}
	return nil
}

func (director *Director) GetJSONBody() string {
	body := `{
		"name": "` + director.Name + `",
		"active": ` + strconv.FormatBool(director.Active) + `
	}`

	return body
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
