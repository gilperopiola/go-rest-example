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

	// User Detail
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
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

	// User Detail
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type GetUserRequest struct {
	ID int `json:"id"`
}

type UpdateUserRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`

	// User Detail
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}

type DeleteUserRequest struct {
	ID int `json:"id"`
}
