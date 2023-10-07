package auth

import (
	"net/http"
	"strings"

	"github.com/gilperopiola/go-rest-example/pkg/entities"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	UnauthorizedMsg = "unauthorized"
)

// ValidateToken validates a token for a specific role and sets ID and Email in context
func (auth *Auth) ValidateToken(role entities.Role) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get token first as a string and then as a *jwt.Token
		token := auth.getTokenStructFromContext(c)

		// Get custom claims from token
		customClaims, ok := token.Claims.(*CustomClaims)

		// Check if token is invalid
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, UnauthorizedMsg)
			c.Abort()
		}

		// Check if role is valid
		if role != entities.AnyRole && customClaims.Role != role {
			c.JSON(http.StatusUnauthorized, UnauthorizedMsg)
			c.Abort()
		}

		// If OK, set ID and Email inside of context
		addInfoToContext(c, customClaims)
	}
}

func (auth *Auth) getTokenStructFromContext(c *gin.Context) *jwt.Token {

	// Get token string from context
	tokenString := removeBearerPrefix(auth.getJWTStringFromHeader(c.Request.Header))

	// Decode string into actual *jwt.Token
	if token, err := auth.decodeTokenString(tokenString); err == nil {
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
		return &jwt.Token{}, entities.ErrUnauthorized
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

func addInfoToContext(c *gin.Context, claims *CustomClaims) {
	c.Set("ID", claims.Id)
	c.Set("Username", claims.Username)
	c.Set("Email", claims.Email)
}
