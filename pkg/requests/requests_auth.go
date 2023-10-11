package requests

import (
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/models"
)

// Here live the input requests that are used in the transport & service layers.
// They are validated on the transport layer, and then passed to the service layer.
// Where they usually end up being converted to models.

type SignupRequest struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

func (r *SignupRequest) ToUserModel() models.User {
	return models.User{
		Email:     r.Email,
		Username:  r.Username,
		Password:  r.Password,
		Deleted:   false,
		IsAdmin:   false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type LoginRequest struct {
	UsernameOrEmail string `json:"username_or_email"`
	Password        string `json:"password"`
}

func (r *LoginRequest) ToUserModel() models.User {
	user := models.User{
		Password: r.Password,
	}

	if validEmailRegex.MatchString(r.UsernameOrEmail) {
		user.Email = r.UsernameOrEmail
	} else {
		user.Username = r.UsernameOrEmail
	}

	return user
}
