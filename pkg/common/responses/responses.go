package responses

import "github.com/gilperopiola/go-rest-example/pkg/common/entities"

// These aren't the HTTP Responses that the API will return, but the responses that the Service Layer
// returns to the Transport Layer

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
