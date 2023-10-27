package requests

import (
	"regexp"

	"github.com/gilperopiola/go-rest-example/pkg/common"
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

//----------------------
//	 AUTH VALIDATIONS
//----------------------

func (req SignupRequest) Validate() error {
	if err := validateUsernameEmailAndPassword(req.Username, req.Email, req.Password); err != nil {
		return err
	}

	if req.Password != req.RepeatPassword {
		return common.ErrPasswordsDontMatch
	}

	return nil
}

func (req LoginRequest) Validate() error {
	if req.UsernameOrEmail == "" || req.Password == "" {
		return common.ErrAllFieldsRequired
	}

	return nil
}

//-------------------------
//	  USER VALIDATIONS
//-------------------------

func (req CreateUserRequest) Validate() error {
	return validateUsernameEmailAndPassword(req.Username, req.Email, req.Password)
}

func (req GetUserRequest) Validate() error {
	if req.UserID == 0 {
		return common.ErrAllFieldsRequired
	}
	return nil
}

func (req UpdateUserRequest) Validate() error {
	if req.UserID == 0 || (req.Email == "" && req.Username == "") {
		return common.ErrAllFieldsRequired
	}

	if req.Email != "" && !validEmailRegex.MatchString(req.Email) {
		return common.ErrInvalidEmailFormat
	}

	if req.Username != "" {
		if len(req.Username) < usernameMinLength || len(req.Username) > usernameMaxLength {
			return common.ErrInvalidUsernameLength
		}
	}

	return nil
}

func (req DeleteUserRequest) Validate() error {
	if req.UserID == 0 {
		return common.ErrAllFieldsRequired
	}

	return nil
}

func (req SearchUsersRequest) Validate() error {
	if req.Page < 0 || req.PerPage <= 0 {
		return common.ErrAllFieldsRequired
	}

	return nil
}

func (req ChangePasswordRequest) Validate() error {
	if req.UserID == 0 {
		return common.ErrAllFieldsRequired
	}

	if req.OldPassword == "" || req.NewPassword == "" || req.RepeatPassword == "" {
		return common.ErrAllFieldsRequired
	}

	if len(req.NewPassword) < passwordMinLength || len(req.NewPassword) > passwordMaxLength {
		return common.ErrInvalidPasswordLength
	}

	if req.NewPassword != req.RepeatPassword {
		return common.ErrPasswordsDontMatch
	}

	return nil
}

//-----------------------------
//	  USER POSTS VALIDATIONS
//-----------------------------

func (req CreateUserPostRequest) Validate() error {
	if req.UserID == 0 || req.Title == "" {
		return common.ErrAllFieldsRequired
	}
	return nil
}

//-----------------
//	  HELPERS
//-----------------

func validateUsernameEmailAndPassword(username, email, password string) error {
	if email == "" || username == "" || password == "" {
		return common.ErrAllFieldsRequired
	}

	if !validEmailRegex.MatchString(email) {
		return common.ErrInvalidEmailFormat
	}

	if len(username) < usernameMinLength || len(username) > usernameMaxLength {
		return common.ErrInvalidUsernameLength
	}

	if len(password) < passwordMinLength || len(password) > passwordMaxLength {
		return common.ErrInvalidPasswordLength
	}

	return nil
}
