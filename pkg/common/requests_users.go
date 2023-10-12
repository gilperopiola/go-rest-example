package common

import (
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/models"
)

// Here live the input requests that are used in the transport & service layers.
// They are validated on the transport layer, and then passed to the service layer.
// Where they usually end up being converted to models.

// - Users
type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

func (r *CreateUserRequest) ToUserModel() models.User {
	return models.User{
		Email:     r.Email,
		Username:  r.Username,
		Password:  r.Password,
		Deleted:   false,
		IsAdmin:   r.IsAdmin,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type GetUserRequest struct {
	ID int `json:"id"`
}

func (r *GetUserRequest) ToUserModel() models.User {
	return models.User{ID: r.ID}
}

type UpdateUserRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (r *UpdateUserRequest) ToUserModel() models.User {
	return models.User{ID: r.ID, Username: r.Username, Email: r.Email}
}

type DeleteUserRequest struct {
	ID int `json:"id"`
}

func (r *DeleteUserRequest) ToUserModel() models.User {
	return models.User{ID: r.ID}
}
