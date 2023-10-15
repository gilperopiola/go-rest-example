package requests

import (
	"regexp"

	customErrors "github.com/gilperopiola/go-rest-example/pkg/errors"
)

const (
	usernameMinLength = 4
	usernameMaxLength = 32
	passwordMinLength = 8
	passwordMaxLength = 64
)

var (
	validEmailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
)

// - Auth

func (req SignupRequest) Validate() error {
	if err := validateUsernameEmailAndPassword(req.Username, req.Email, req.Password); err != nil {
		return err
	}

	if req.Password != req.RepeatPassword {
		return customErrors.ErrPasswordsDontMatch
	}

	return nil
}

func (req LoginRequest) Validate() error {
	if req.UsernameOrEmail == "" || req.Password == "" {
		return customErrors.ErrAllFieldsRequired
	}

	return nil
}

// - Users

func (req CreateUserRequest) Validate() error {
	return validateUsernameEmailAndPassword(req.Username, req.Email, req.Password)
}

func (req GetUserRequest) Validate() error {
	if req.ID == 0 {
		return customErrors.ErrAllFieldsRequired
	}
	return nil
}

func (req UpdateUserRequest) Validate() error {
	if req.ID == 0 || (req.Email == "" && req.Username == "") {
		return customErrors.ErrAllFieldsRequired
	}

	if req.Email != "" && !validEmailRegex.MatchString(req.Email) {
		return customErrors.ErrInvalidEmailFormat
	}

	if req.Username != "" {
		if len(req.Username) < usernameMinLength || len(req.Username) > usernameMaxLength {
			return customErrors.ErrInvalidUsernameLength
		}
	}

	return nil
}

func (req DeleteUserRequest) Validate() error {
	if req.ID == 0 {
		return customErrors.ErrAllFieldsRequired
	}

	return nil
}

func validateUsernameEmailAndPassword(username, email, password string) error {
	if email == "" || username == "" || password == "" {
		return customErrors.ErrAllFieldsRequired
	}

	if !validEmailRegex.MatchString(email) {
		return customErrors.ErrInvalidEmailFormat
	}

	if len(username) < usernameMinLength || len(username) > usernameMaxLength {
		return customErrors.ErrInvalidUsernameLength
	}

	if len(password) < passwordMinLength || len(password) > passwordMaxLength {
		return customErrors.ErrInvalidPasswordLength
	}

	return nil
}
