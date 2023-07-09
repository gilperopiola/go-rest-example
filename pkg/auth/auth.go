package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/config"
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go"
)

const (
	UNAUTHORIZED = "unauthorized"
)

func GenerateToken(user entities.User, sessionDurationDays int, secret string) string {

	issuedAt := time.Now().Unix()
	expiresAt := time.Now().Add(time.Hour * 24 * time.Duration(sessionDurationDays)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.StandardClaims{
			Id:        fmt.Sprint(user.ID),
			Audience:  user.Email,
			IssuedAt:  issuedAt,
			ExpiresAt: expiresAt,
		},
	)

	tokenString, _ := token.SignedString([]byte(secret))
	return tokenString
}

func ValidateToken(config config.JWTConfig) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get token string from context
		tokenString := getTokenString(c)

		// Decode string into actual *jwt.Token
		token, err := decodeToken(tokenString, config.SECRET)
		if err != nil {
			c.JSON(http.StatusUnauthorized, UNAUTHORIZED)
			c.Abort()
			return
		}

		// Check if token is valid, set ID and Email in context
		if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
			c.Set("ID", claims.Id)
			c.Set("Email", claims.Audience)
		} else {
			c.JSON(http.StatusUnauthorized, UNAUTHORIZED)
			c.Abort()
		}
	}
}

// decodeToken decodes a JWT token string into a *jwt.Token
func decodeToken(tokenString, jwtSecret string) (*jwt.Token, error) {
	if len(tokenString) < 40 {
		return &jwt.Token{}, nil
	}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	}

	return jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, keyFunc)
}

func getTokenString(c *gin.Context) string {
	tokenString := c.Request.Header.Get("Authorization")
	return strings.TrimPrefix(tokenString, "Bearer ")
}
