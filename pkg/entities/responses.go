package entities

// Here are the output responses that are used in the transport/service layer.
// They are generated on the service layer, and then passed to the transport layer.

type SignupResponse struct {
	User User `json:"user"`
}
type LoginResponse struct {
	Token string `json:"token"`
}

type CreateUserResponse struct {
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
