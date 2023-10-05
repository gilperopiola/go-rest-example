package entities

// These are the responses that the service layer returns to the transport layer

// - Auth
type SignupResponse struct {
	User User `json:"user"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// - Users
type CreateUserResponse struct {
	User User `json:"user"`
}

type GetUserResponse struct {
	User User `json:"user"`
}

type UpdateUserResponse struct {
	User User `json:"user"`
}

type DeleteUserResponse struct {
	User User `json:"user"`
}
