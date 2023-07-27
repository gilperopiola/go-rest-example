package entities

import (
	"regexp"
)

const (

	// Signup

	USERNAME_MIN_LENGTH = 4
	USERNAME_MAX_LENGTH = 32
	PASSWORD_MIN_LENGTH = 8
	PASSWORD_MAX_LENGTH = 64
	VALID_EMAIL_REGEX   = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
)

// SignupRequest Validate {Email, Username, Password, RepeatPassword}
func (request *SignupRequest) Validate() error {

	// Empty fields
	if request.Email == "" || request.Username == "" || request.Password == "" || request.RepeatPassword == "" {
		return ErrAllFieldsRequired
	}

	// Matching passwords
	if request.Password != request.RepeatPassword {
		return ErrPasswordsDontMatch
	}

	// Valid email format
	matched, err := regexp.MatchString(VALID_EMAIL_REGEX, request.Email)
	if err != nil || !matched {
		return ErrInvalidEmailFormat
	}

	// Valid username and password length
	if len(request.Username) < USERNAME_MIN_LENGTH || len(request.Username) > USERNAME_MAX_LENGTH {
		return ErrInvalidUsernameLength
	}

	if len(request.Password) < PASSWORD_MIN_LENGTH || len(request.Password) > PASSWORD_MAX_LENGTH {
		return ErrInvalidPasswordLength
	}

	// Return OK
	return nil
}

// LoginRequest Validate {UsernameOrEmail, Password}
func (req *LoginRequest) Validate() error {

	// Empty fields
	if req.UsernameOrEmail == "" || req.Password == "" {
		return ErrAllFieldsRequired
	}

	return nil
}

func (req *GetUserRequest) Validate() error {

	// Empty ID
	if req.ID == 0 {
		return ErrAllFieldsRequired
	}

	return nil
}

// UpdateUserRequest Validate {ID, Email, Username}
func (request *UpdateUserRequest) Validate() error {

	// Empty fields
	if request.ID == 0 || (request.Email == "" && request.Username == "") {
		return ErrAllFieldsRequired
	}

	// Valid email format
	if request.Email != "" {
		matched, err := regexp.MatchString(VALID_EMAIL_REGEX, request.Email)
		if err != nil || !matched {
			return ErrInvalidEmailFormat
		}
	}

	// Valid username length
	if request.Username != "" {
		if len(request.Username) < USERNAME_MIN_LENGTH || len(request.Username) > USERNAME_MAX_LENGTH {
			return ErrInvalidUsernameLength
		}
	}

	// Return OK
	return nil
}

// DeleteUserRequest Validate {ID}
func (req *DeleteUserRequest) Validate() error {

	// Empty ID
	if req.ID == 0 {
		return ErrAllFieldsRequired
	}

	return nil
}
