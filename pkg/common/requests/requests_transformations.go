package requests

import (
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/common/models"
)

// - Auth

func (r *SignupRequest) ToUserModel() models.User {
	return models.User{
		Email:    r.Email,
		Username: r.Username,
		Password: r.Password,
		Deleted:  false,
		Details: models.UserDetail{
			FirstName: r.FirstName,
			LastName:  r.LastName,
		},
		IsAdmin:   false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
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

// - Users

func (r *CreateUserRequest) ToUserModel() models.User {
	return models.User{
		Email:    r.Email,
		Username: r.Username,
		Password: r.Password,
		Deleted:  false,
		Details: models.UserDetail{
			FirstName: r.FirstName,
			LastName:  r.LastName,
		},
		IsAdmin:   r.IsAdmin,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (r *GetUserRequest) ToUserModel() models.User {
	return models.User{ID: r.ID}
}

func (r *UpdateUserRequest) ToUserModel() models.User {
	var (
		firstName = ""
		lastName  = ""
	)
	if r.FirstName != nil {
		firstName = *r.FirstName
	}
	if r.LastName != nil {
		lastName = *r.LastName
	}

	return models.User{
		ID:       r.ID,
		Username: r.Username,
		Email:    r.Email,
		Details: models.UserDetail{
			FirstName: firstName,
			LastName:  lastName,
		},
	}
}

func (r *DeleteUserRequest) ToUserModel() models.User {
	return models.User{ID: r.ID}
}
