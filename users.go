package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type UserActions interface {
	Validate() error
	ValidateLogIn() error
}

type User struct {
	ID          uint      `json:"id" gorm:"auto_increment;unique;not null"`
	Username    string    `json:"username" gorm:"unique;not null"`
	Email       string    `json:"email" gorm:"unique;not null"`
	Password    string    `json:"password,omitempty" gorm:"not null"`
	Admin       bool      `json:"admin" gorm:"default: 0"`
	Active      bool      `json:"active"`
	DateCreated time.Time `json:"date_created" gorm:"default: current_timestamp"`

	Token string `json:"token,omitempty" gorm:"-"`
}

func CreateUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	log.Println(user)

	if err := user.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user.Password = hash(user.Email, user.Password)

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, beautifyDatabaseError(err))
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, user)
}

func ReadUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, user)
}

func ReadUsers(c *gin.Context) {
	var users []User

	id, username, email, limit, offset, sortField, sortDir := getReadUsersParameters(c)

	tempDB := db

	if id > 0 {
		tempDB = tempDB.Where("id = ?", id)
	}

	tempDB = tempDB.Where("username LIKE ? AND email LIKE ?", username, email)

	if sortField != "" && (sortDir == "ASC" || sortDir == "DESC") {
		tempDB = tempDB.Order(sortField + " " + sortDir)
	}

	tempDB.Limit(limit).Offset(offset).Find(&users)

	for key := range users {
		users[key].Password = ""
	}

	c.JSON(http.StatusOK, users)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.BindJSON(&user)
	if err := db.Save(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, beautifyDatabaseError(err))
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, user)
}

//extra

func (user *User) Validate() error {
	if len(user.Username) == 0 || len(user.Email) == 0 || len(user.Password) == 0 {
		return errors.New("all fields required")
	}

	if len(user.Username) < config.USERS.USERNAME_MIN_CHARACTERS || len(user.Username) > config.USERS.USERNAME_MAX_CHARACTERS {
		return errors.New("username must have between " + string(config.USERS.USERNAME_MIN_CHARACTERS) + " and " + string(config.USERS.USERNAME_MAX_CHARACTERS) + " characters")
	}

	if !strings.Contains(user.Email, "@") {
		return errors.New("email format invalid")
	}

	return nil
}

func (user *User) ValidateLogIn() error {
	if len(user.Username) == 0 || len(user.Password) == 0 {
		return errors.New("both fields required")
	}
	return nil
}

func getReadUsersParameters(c *gin.Context) (int, string, string, string, string, string, string) {
	id, _ := strconv.Atoi(c.Query("ID"))
	username := "%" + c.Query("Username") + "%"
	email := "%" + c.Query("Email") + "%"
	limit := c.Query("Limit")
	offset := c.Query("Offset")
	sortField := c.Query("SortField")
	sortDir := c.Query("SortDir")
	return id, username, email, limit, offset, sortField, sortDir
}
