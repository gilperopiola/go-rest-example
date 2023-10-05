package auth

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	UnauthorizedMsg = "unauthorized"
)

// ValidateToken validates a token for any role and sets ID and Email in context
func (auth *Auth) ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get token first as a string and then as a *jwt.Token
		token := auth.getTokenStructFromContext(c)

		// Check if token is valid, then set ID and Email in context
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			c.Set("ID", claims.Id)
			c.Set("Email", claims.Audience)
			return
		}

		c.JSON(http.StatusUnauthorized, UnauthorizedMsg)
		c.Abort()
	}
}

// ValidateRole validates a token for a specific role and sets ID and Email in context
func (auth *Auth) ValidateRole(role AuthRole) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get token first as a string and then as a *jwt.Token
		token := auth.getTokenStructFromContext(c)

		// Get custom claims from token
		customClaims, ok := token.Claims.(*CustomClaims)

		// Check if token is valid, then set ID and Email in context
		if ok && token.Valid && customClaims.Role == role {
			c.Set("ID", customClaims.Id)
			c.Set("Email", customClaims.Audience)
			return
		}

		c.JSON(http.StatusUnauthorized, UnauthorizedMsg)
		c.Abort()
	}
}

func (auth *Auth) getTokenStructFromContext(c *gin.Context) *jwt.Token {

	// Get token string from context
	tokenString := removeBearerPrefix(auth.getJWTStringFromHeader(c.Request.Header))

	// Decode string into actual *jwt.Token
	token, err := auth.decodeTokenString(tokenString)
	if err == nil {
		return token
	}

	// Error decoding token
	c.JSON(http.StatusUnauthorized, UnauthorizedMsg)
	c.Abort()
	return nil
}

// decodeTokenString decodes a JWT token string into a *jwt.Token
func (auth *Auth) decodeTokenString(tokenString string) (*jwt.Token, error) {

	// Check length
	if len(tokenString) < 40 {
		return &jwt.Token{}, nil
	}

	// Make key function
	keyFunc := func(token *jwt.Token) (interface{}, error) { return []byte(auth.secret), nil }

	// Parse
	return jwt.ParseWithClaims(tokenString, &CustomClaims{}, keyFunc)
}

func (auth *Auth) getJWTStringFromHeader(header http.Header) string {
	return header.Get("Authorization")
}

func removeBearerPrefix(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}
