package requests

import (
	"regexp"

	"github.com/gilperopiola/go-rest-example/pkg/entities"
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

// Users

func (req CreateUserRequest) Validate() error {
	return validateUsernameEmailAndPassword(req.Username, req.Email, req.Password)
}

func (req GetUserRequest) Validate() error {
	if req.ID == 0 {
		return entities.ErrAllFieldsRequired
	}
	return nil
}

func (req UpdateUserRequest) Validate() error {
	if req.ID == 0 || (req.Email == "" && req.Username == "") {
		return entities.ErrAllFieldsRequired
	}

	if req.Email != "" && !validEmailRegex.MatchString(req.Email) {
		return entities.ErrInvalidEmailFormat
	}

	if req.Username != "" {
		if len(req.Username) < usernameMinLength || len(req.Username) > usernameMaxLength {
			return entities.ErrInvalidUsernameLength
		}
	}

	return nil
}

func (req DeleteUserRequest) Validate() error {
	if req.ID == 0 {
		return entities.ErrAllFieldsRequired
	}

	return nil
}

func validateUsernameEmailAndPassword(username, email, password string) error {
	if email == "" || username == "" || password == "" {
		return entities.ErrAllFieldsRequired
	}

	if !validEmailRegex.MatchString(email) {
		return entities.ErrInvalidEmailFormat
	}

	if len(username) < usernameMinLength || len(username) > usernameMaxLength {
		return entities.ErrInvalidUsernameLength
	}

	if len(password) < passwordMinLength || len(password) > passwordMaxLength {
		return entities.ErrInvalidPasswordLength
	}

	return nil
}

// Auth

func (req SignupRequest) Validate() error {
	if err := validateUsernameEmailAndPassword(req.Username, req.Email, req.Password); err != nil {
		return err
	}

	if req.Password != req.RepeatPassword {
		return entities.ErrPasswordsDontMatch
	}

	return nil
}

func (req LoginRequest) Validate() error {
	if req.UsernameOrEmail == "" || req.Password == "" {
		return entities.ErrAllFieldsRequired
	}

	return nil
}
