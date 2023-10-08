package auth

import (
	"fmt"
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/entities"

	"github.com/golang-jwt/jwt/v4"
)

func (auth *Auth) GenerateToken(user entities.User, role entities.Role) string {

	var (
		issuedAt  = time.Now().Unix()
		expiresAt = time.Now().Add(time.Hour * 24 * time.Duration(auth.sessionDurationDays)).Unix()
	)

	claims := &CustomClaims{
		Username: user.Username,
		Email:    user.Email,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			Id:        fmt.Sprint(user.ID),
			Audience:  user.Email,
			IssuedAt:  issuedAt,
			ExpiresAt: expiresAt,
		},
	}

	// Generate token (struct)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token (string)
	tokenString, _ := token.SignedString([]byte(auth.secret))

	return tokenString
}
