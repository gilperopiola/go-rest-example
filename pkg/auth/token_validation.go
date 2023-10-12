package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gilperopiola/go-rest-example/pkg/entities"
	customErrors "github.com/gilperopiola/go-rest-example/pkg/errors"
	"github.com/gilperopiola/go-rest-example/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type HTTPResponse struct {
	Success bool        `json:"success"`
	Content interface{} `json:"content"`
	Error   string      `json:"error"`
}

// ValidateToken validates a token for a specific role and sets ID and Email in context
func (auth *Auth) ValidateToken(role entities.Role, shouldMatchUserID bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get token string and then convert it to a *jwt.Token
		token, err := auth.getTokenStructFromContext(c)
		if err != nil {
			abortRequest(c)
		}

		// Get custom claims from token
		customClaims, ok := token.Claims.(*CustomClaims)

		// Check if claims and token are valid
		if !ok || !token.Valid || customClaims.Valid() != nil {
			abortRequest(c)
		}

		// Check if role is valid
		if role != entities.AnyRole && customClaims.Role != role {
			abortRequest(c)
		}

		// Check if user ID in URL matches user ID in token
		if shouldMatchUserID {
			urlUserID, err := utils.GetIntFromContextURLParams(c.Params, "user_id")
			if err != nil {
				abortRequest(c)
			}

			if customClaims.ID != fmt.Sprint(urlUserID) {
				abortRequest(c)
			}
		}

		// If OK, set ID, Username and Email inside of context
		addUserInfoToContext(c, customClaims)
	}
}

func (auth *Auth) getTokenStructFromContext(c *gin.Context) (*jwt.Token, error) {

	// Get token string from context
	tokenString := removeBearerPrefix(getJWTStringFromHeader(c.Request.Header))

	// Decode string into actual *jwt.Token
	token, err := auth.decodeTokenString(tokenString)
	if err != nil {
		return nil, err
	}

	// Token decoded OK
	return token, nil
}

// decodeTokenString decodes a JWT token string into a *jwt.Token
func (auth *Auth) decodeTokenString(tokenString string) (*jwt.Token, error) {

	// Check length
	if len(tokenString) < 40 {
		return &jwt.Token{}, customErrors.ErrUnauthorized
	}

	// Make key function
	keyFunc := func(token *jwt.Token) (interface{}, error) { return []byte(auth.secret), nil }

	// Parse
	return jwt.ParseWithClaims(tokenString, &CustomClaims{}, keyFunc)
}

func addUserInfoToContext(c *gin.Context, claims *CustomClaims) {
	c.Set("ID", claims.ID)
	c.Set("Username", claims.Username)
	c.Set("Email", claims.Email)
}

func getJWTStringFromHeader(header http.Header) string {
	return header.Get("Authorization")
}

func removeBearerPrefix(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}

func abortRequest(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, HTTPResponse{
		Success: false,
		Content: nil,
		Error:   "unauthorized",
	})
	c.Abort()
}
