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

func (codec *Codec) FromLoginRequestToUserModel(request entities.LoginRequest) models.User {
	out := models.User{Password: request.Password}
	usernameOrEmail := request.UsernameOrEmail

	// If it's an email, login with email. Otherwise login with username
	if matchesEmailFormat, _ := regexp.MatchString(VALID_EMAIL_REGEX, usernameOrEmail); matchesEmailFormat {
		out.Email = usernameOrEmail
	} else {
		out.Username = usernameOrEmail
	}

	return out
}

func (codec *Codec) FromCreateUserRequestToUserModel(request entities.CreateUserRequest, hashedPassword string) models.User {
	return models.User{
		Email:     request.Email,
		Username:  request.Username,
		Password:  hashedPassword,
		Deleted:   false,
		IsAdmin:   request.IsAdmin,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (codec *Codec) FromGetUserRequestToUserModel(request entities.GetUserRequest) models.User {
	return models.User{ID: request.ID}
}

func (codec *Codec) FromUpdateUserRequestToUserModel(request entities.UpdateUserRequest) models.User {
	return models.User{
		ID: request.ID,
	}
}
