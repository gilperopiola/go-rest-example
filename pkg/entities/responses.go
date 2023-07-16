package entities

type SignupResponse struct {
	User User `json:"user"`
}
type LoginResponse struct {
	Token string `json:"token"`
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
