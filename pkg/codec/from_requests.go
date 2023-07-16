package codec

import (
	"regexp"
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/models"
)

const (
	VALID_EMAIL_REGEX = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
)

func (codec *Codec) FromSignupRequestToUserModel(request entities.SignupRequest, hashedPassword string) models.User {
	return models.User{
		Email:     request.Email,
		Username:  request.Username,
		Password:  hashedPassword,
		Deleted:   false,
		IsAdmin:   false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (codec *Codec) FromLoginRequestToUserCredentials(request entities.LoginRequest) entities.UserCredentials {
	out := entities.UserCredentials{Password: request.Password}

	// If it's an email, login with email. Otherwise login with username
	if matchesEmailFormat, _ := regexp.MatchString(VALID_EMAIL_REGEX, request.UsernameOrEmail); matchesEmailFormat {
		out.Email = request.UsernameOrEmail
	} else {
		out.Username = request.UsernameOrEmail
	}

	return out
}
