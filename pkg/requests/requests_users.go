package requests

import (
	"regexp"
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/models"
)

// Here live the input requests that are used in the transport & service layers.
// They are validated on the transport layer, and then passed to the service layer.
// Where they usually end up being converted to models.

const (
	usernameMinLength = 4
	usernameMaxLength = 32
	passwordMinLength = 8
	passwordMaxLength = 64
)

var (
	validEmailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
)

// - Users
type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

func (r *CreateUserRequest) ToUserModel() models.User {
	return models.User{
		Email:     r.Email,
		Username:  r.Username,
		Password:  r.Password,
		Deleted:   false,
		IsAdmin:   r.IsAdmin,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (req CreateUserRequest) Validate() error {
	return validateUsernameEmailAndPassword(req.Username, req.Email, req.Password)
}

type GetUserRequest struct {
	ID int `json:"id"`
}

func (r *GetUserRequest) ToUserModel() models.User {
	return models.User{ID: r.ID}
}

func (req GetUserRequest) Validate() error {
	if req.ID == 0 {
		return entities.ErrAllFieldsRequired
	}
	return nil
}

type UpdateUserRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (r *UpdateUserRequest) ToUserModel() models.User {
	return models.User{ID: r.ID, Username: r.Username, Email: r.Email}
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

type DeleteUserRequest struct {
	ID int `json:"id"`
}

func (r *DeleteUserRequest) ToUserModel() models.User {
	return models.User{ID: r.ID}
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
