package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//SignUp takes {username, email, password}, creates a user and returns it
func SignUp(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	if err := user.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user.Password = hash(user.Email, user.Password)
	user.Active = true

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, beautifyDatabaseError(err))
		return
	}

	user.Token = generateToken(user)
	user.Password = ""
	c.JSON(http.StatusOK, user)
}

//LogIn takes {username, password}, checks if the user exists and returns it
func LogIn(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	if err := user.ValidateLogIn(); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var dbUser User
	db.Where("username = ?", user.Username).First(&dbUser)
	if dbUser.ID == 0 {
		c.JSON(http.StatusBadRequest, "wrong username")
		return
	}

	if dbUser.Password != hash(dbUser.Email, user.Password) {
		c.JSON(http.StatusBadRequest, "wrong password")
		return
	}

	dbUser.Token = generateToken(dbUser)
	dbUser.Password = ""
	c.JSON(http.StatusOK, dbUser)
}

func validateToken(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")

		if len(tokenString) < 40 {
			c.JSON(http.StatusUnauthorized, "authentication error")
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JWT.SECRET), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, "authentication error")
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
			if claims.Subject != requiredRole {
				c.JSON(http.StatusUnauthorized, "authentication error")
				c.Abort()
				return
			}
			
			c.Set("ID", claims.Id)
			c.Set("Email", claims.Audience)
			c.Set("Role", claims.Subject)
		} else {
			c.JSON(http.StatusUnauthorized, "authentication error")
			c.Abort()
		}
	}
}

func generateToken(user User) string {
	var role string
	if user.Admin {
		role = "Admin"
	} else {
		role = "User"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        fmt.Sprint(user.ID),
		Audience:  user.Email,
		Subject:   role,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 24 * time.Duration(config.JWT.SESSION_DURATION)).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(config.JWT.SECRET))
	return tokenString
}

func generateTestingToken(role string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        "1",
		Audience:  "test@test.com",
		Subject:   role,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(config.JWT.SECRET))
	return tokenString
}

func hash(salt string, data string) string {
	hasher := sha1.New()
	hasher.Write([]byte(salt + data))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
