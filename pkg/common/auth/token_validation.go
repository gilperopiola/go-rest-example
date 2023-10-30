package auth

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gilperopiola/go-rest-example/pkg/common"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var (
	pathUserIDKey = "user_id"
)

// ValidateToken validates a token for a specific role and sets ID and Email in context
func (auth *Auth) ValidateToken(role Role, shouldMatchUserID bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get token string and then convert it to a *jwt.Token
		token, err := auth.getTokenStructFromContext(c)
		if err != nil {
			c.Error(common.Wrap("auth.getTokenStructFromContext", common.ErrUnauthorized))
			c.Abort()
			return
		}

		// Get custom claims from token
		customClaims, ok := token.Claims.(*CustomClaims)

		// Check if claims and token and role are valid
		if !ok || !token.Valid || customClaims.Valid() != nil || (role != AnyRole && customClaims.Role != role) {
			c.Error(common.Wrap("!token.Valid || customClaims.Role != role", common.ErrUnauthorized))
			c.Abort()
			return
		}

		// Check if user ID in URL matches user ID in token
		if shouldMatchUserID {
			urlUserID, err := getIntFromURLPath(c.Params, pathUserIDKey)
			if err != nil || customClaims.ID != fmt.Sprint(urlUserID) {
				c.Error(common.Wrap("!shouldMatchUserID", common.ErrUnauthorized))
				c.Abort()
				return
			}
		}

		// If OK, set UserID, Username and Email inside of context
		userID, _ := strconv.Atoi(customClaims.ID)
		addUserInfoToContext(c, userID, customClaims.Username, customClaims.Email)
	}
}

func addUserInfoToContext(c *gin.Context, id int, username, email string) {
	c.Set("UserID", id)
	c.Set("Username", username)
	c.Set("Email", email)
}

func (auth *Auth) getTokenStructFromContext(c *gin.Context) (*jwt.Token, error) {

	// Get token string from headers
	tokenString := strings.TrimPrefix(c.Request.Header.Get("Authorization"), "Bearer ")

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
		return &jwt.Token{}, common.ErrUnauthorized
	}

	// Make key function
	keyFunc := func(token *jwt.Token) (interface{}, error) { return []byte(auth.secret), nil }

	// Parse
	return jwt.ParseWithClaims(tokenString, &CustomClaims{}, keyFunc)
}

func getIntFromURLPath(params gin.Params, key string) (int, error) {
	value, ok := params.Get(key)
	if !ok {
		return 0, fmt.Errorf("error getting %s from URL params", key)
	}

	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("error converting %s from string to int", key)
	}

	return valueInt, nil
}
