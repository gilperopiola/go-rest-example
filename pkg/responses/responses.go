package responses

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
)

// These are the responses that the service layer returns to the transport layer

// - Auth
type SignupResponse struct {
	User entities.User `json:"user"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// - Users
type CreateUserResponse struct {
	User entities.User `json:"user"`
}

type GetUserResponse struct {
	User entities.User `json:"user"`
}

type UpdateUserResponse struct {
	User entities.User `json:"user"`
}

type DeleteUserResponse struct {
	User entities.User `json:"user"`
}