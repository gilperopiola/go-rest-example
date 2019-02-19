package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//SignUp takes {username, email, password}, creates a user and returns it
func SignUp(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	user.Password = hash(user.Email, user.Password)

	if err := db.Create(&user).Error; err != nil {
		c.JSON(400, beautifyDatabaseError(err))
		return
	}

	user.Token = generateToken(user)
	user.Password = ""
	c.JSON(200, user)
}

//LogIn takes {username, password}, checks if the user exists and returns it
func LogIn(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	var dbUser User
	db.Where("username = ?", user.Username).First(&dbUser)
	if dbUser.ID == 0 {
		c.JSON(400, "wrong username")
		return
	}

	if dbUser.Password != hash(dbUser.Email, user.Password) {
		c.JSON(400, "wrong password")
		return
	}

	dbUser.Token = generateToken(dbUser)
	dbUser.Password = ""
	c.JSON(200, dbUser)
}

func generateToken(user User) string {
	claims := jwt.StandardClaims{
		Id:       fmt.Sprint(user.ID),
		Audience: user.Email,

		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 24 * time.Duration(config.JWT.SESSION_DURATION)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(config.JWT.SECRET))
	return tokenString
}

func hash(salt string, data string) string {
	hasher := sha1.New()
	hasher.Write([]byte(salt + data))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
