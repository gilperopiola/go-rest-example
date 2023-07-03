package entities

import (
	"regexp"
)

const (
	// Signup - Login
	USERNAME_MIN_LENGTH = 4
	USERNAME_MAX_LENGTH = 32
	PASSWORD_MIN_LENGTH = 8
	PASSWORD_MAX_LENGTH = 64
)

func (req *SignupRequest) Validate() error {

	// Empty fields
	if req.Email == "" || req.Username == "" || req.Password == "" || req.RepeatPassword == "" {
		return ErrAllFieldsRequired
	}

	// Matching passwords
	if req.Password != req.RepeatPassword {
		return ErrPasswordsDontMatch
	}

	// Valid email format
	matched, err := regexp.MatchString(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`, req.Email)
	if err != nil || !matched {
		return ErrInvalidEmailFormat
	}

	// Valid username and password length
	if len(req.Username) < USERNAME_MIN_LENGTH || len(req.Username) > USERNAME_MAX_LENGTH {
		return ErrInvalidUsernameLength
	}

	if len(req.Password) < PASSWORD_MIN_LENGTH || len(req.Password) > PASSWORD_MAX_LENGTH {
		return ErrInvalidPasswordLength
	}

	// Return OK
	return nil
}
