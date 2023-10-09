package entities

import (
	"regexp"
)

const (
	// - Signup & Users
	usernameMinLength = 4
	usernameMaxLength = 32
	passwordMinLength = 8
	passwordMaxLength = 64
)

var (
	validEmailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
)

// SignupRequest Validate {Email, Username, Password, RepeatPassword}
func (req SignupRequest) Validate() error {
	// Check username, email and password basic validations
	if err := validateUsernameEmailAndPassword(req.Username, req.Email, req.Password); err != nil {
		return err
	}

	// Check passwords match
	if req.Password != req.RepeatPassword {
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
func (req CreateUserRequest) Validate() error {
	return validateUsernameEmailAndPassword(req.Username, req.Email, req.Password)
}

func (req GetUserRequest) Validate() error {
	// Empty ID
	if req.ID == 0 {
		return ErrAllFieldsRequired
	}
	return nil
}

// UpdateUserRequest Validate {ID, Email, Username}
func (req UpdateUserRequest) Validate() error {
	// Empty fields
	if req.ID == 0 || (req.Email == "" && req.Username == "") {
		return ErrAllFieldsRequired
	}

	// Valid email format
	if req.Email != "" && !validEmailRegex.MatchString(req.Email) {
		return ErrInvalidEmailFormat
	}

	// Valid username length
	if req.Username != "" {
		if len(req.Username) < usernameMinLength || len(req.Username) > usernameMaxLength {
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

func validateUsernameEmailAndPassword(username, email, password string) error {
	// Empty fields
	if email == "" || username == "" || password == "" {
		return ErrAllFieldsRequired
	}

	// Valid email format
	if !validEmailRegex.MatchString(email) {
		return ErrInvalidEmailFormat
	}

	// Valid username and password length
	if len(username) < usernameMinLength || len(username) > usernameMaxLength {
		return ErrInvalidUsernameLength
	}

	if len(password) < passwordMinLength || len(password) > passwordMaxLength {
		return ErrInvalidPasswordLength
	}

	// Return OK
	return nil
}
