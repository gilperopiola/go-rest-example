package entities

import (
	"regexp"
)

const (
	// - Signup & Users
	USERNAME_MIN_LENGTH = 4
	USERNAME_MAX_LENGTH = 32
	PASSWORD_MIN_LENGTH = 8
	PASSWORD_MAX_LENGTH = 64
	VALID_EMAIL_REGEX   = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
)

// SignupRequest Validate {Email, Username, Password, RepeatPassword}
func (request SignupRequest) Validate() error {

	// Check username, email and password basic validations
	if err := usernameEmailAndPasswordValidation(request.Username, request.Email, request.Password); err != nil {
		return err
	}

	// Check passwords match
	if request.Password != request.RepeatPassword {
		return ErrPasswordsDontMatch
	}

	// Return OK
	return nil
}

// LoginRequest Validate {UsernameOrEmail, Password}
func (req LoginRequest) Validate() error {

	// Empty fields
	if req.UsernameOrEmail == "" || req.Password == "" {
		return ErrAllFieldsRequired
	}

	return nil
}

// CreateUserRequest Validate {Email, Username, Password, IsAdmin}
func (request CreateUserRequest) Validate() error {
	return usernameEmailAndPasswordValidation(request.Username, request.Email, request.Password)
}

func (req GetUserRequest) Validate() error {

	// Empty ID
	if req.ID == 0 {
		return ErrAllFieldsRequired
	}

	return nil
}

// UpdateUserRequest Validate {ID, Email, Username}
func (request UpdateUserRequest) Validate() error {

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
func (req DeleteUserRequest) Validate() error {

	// Empty ID
	if req.ID == 0 {
		return ErrAllFieldsRequired
	}

	return nil
}

func usernameEmailAndPasswordValidation(username, email, password string) error {

	// Empty fields
	if email == "" || username == "" || password == "" {
		return ErrAllFieldsRequired
	}

	// Valid email format
	matched, err := regexp.MatchString(VALID_EMAIL_REGEX, email)
	if err != nil || !matched {
		return ErrInvalidEmailFormat
	}

	// Valid username and password length
	if len(username) < USERNAME_MIN_LENGTH || len(username) > USERNAME_MAX_LENGTH {
		return ErrInvalidUsernameLength
	}

	if len(password) < PASSWORD_MIN_LENGTH || len(password) > PASSWORD_MAX_LENGTH {
		return ErrInvalidPasswordLength
	}

	// Return OK
	return nil
}
