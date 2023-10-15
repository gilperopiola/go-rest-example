package requests

// Here live the input requests that are used in the Transport & Service layers.
// They are created & validated on the Transport Layer, and then passed to the Service.
// Where they usually end up being converted to models.

// - Auth

type SignupRequest struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

type LoginRequest struct {
	UsernameOrEmail string `json:"username_or_email"`
	Password        string `json:"password"`
}

// - Users

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

type GetUserRequest struct {
	ID int `json:"id"`
}

type UpdateUserRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type DeleteUserRequest struct {
	ID int `json:"id"`
}
