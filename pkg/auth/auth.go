package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/entities"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	UNAUTHORIZED = "unauthorized"
)

type Auth struct {
	secret              string
	sessionDurationDays int
}

type AuthProvider interface {
	GenerateToken(user entities.User) string
	ValidateToken() gin.HandlerFunc

	decodeToken(tokenString string) (*jwt.Token, error)
	getTokenStringFromHeaders(c *gin.Context) string
}

func NewAuth(secret string, sessionDurationDays int) *Auth {
	return &Auth{
		secret:              secret,
		sessionDurationDays: sessionDurationDays,
	}
}

/* ----------------------- */

func (auth *Auth) GenerateToken(user entities.User) string {

	issuedAt := time.Now().Unix()
	expiresAt := time.Now().Add(time.Hour * 24 * time.Duration(auth.sessionDurationDays)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.StandardClaims{
			Id:        fmt.Sprint(user.ID),
			Audience:  user.Email,
			IssuedAt:  issuedAt,
			ExpiresAt: expiresAt,
		},
	)

	tokenString, _ := token.SignedString([]byte(auth.secret))
	return tokenString
}

func (auth *Auth) ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get token string from context
		tokenString := auth.getTokenStringFromHeaders(c)

		// Decode string into actual *jwt.Token
		token, err := auth.decodeToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, UNAUTHORIZED)
			c.Abort()
			return
		}

		// Check if token is valid, then set ID and Email in context
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
func (auth *Auth) decodeToken(tokenString string) (*jwt.Token, error) {
	if len(tokenString) < 40 {
		return &jwt.Token{}, nil
	}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(auth.secret), nil
	}

	return jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, keyFunc)
}

func (auth *Auth) getTokenStringFromHeaders(c *gin.Context) string {
	tokenString := c.Request.Header.Get("Authorization")
	return strings.TrimPrefix(tokenString, "Bearer ")
}
