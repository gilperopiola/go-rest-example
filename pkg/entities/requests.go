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

/* ----------------------------- */

func (req *SignupRequest) Validate() error {
	if req.Email == "" || req.Password == "" || req.RepeatPassword == "" {
		return ErrAllFieldsRequired
	}

	if req.Password != req.RepeatPassword {
		return ErrPasswordsDontMatch
	}

	return nil
}
