package entities

type SignupRequest struct {
	Email          string `json:"email"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

type LoginRequest struct {
	EmailOrUsername string
	Password        string
}
