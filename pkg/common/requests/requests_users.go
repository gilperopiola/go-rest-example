package requests

import (
	"regexp"
	"strconv"

	"github.com/gilperopiola/go-rest-example/pkg/common"
)

var (
	contextUserIDKey = "UserID"
	pathUserIDKey    = "user_id"

	usernameMinLength = 4
	usernameMaxLength = 32
	passwordMinLength = 8
	passwordMaxLength = 64

	validEmailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
)

/*---------------
//    Signup
//-------------*/

type SignupRequest struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`

	// User Detail
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (req *SignupRequest) Build(c common.GinI) error {
	return bindRequestBody(c, req)
}

func (req SignupRequest) Validate() error {
	if err := validateUsernameEmailAndPassword(req.Username, req.Email, req.Password); err != nil {
		return err
	}

	if req.Password != req.RepeatPassword {
		return common.ErrPasswordsDontMatch
	}

	return nil
}

/*--------------
//    Login
//------------*/

type LoginRequest struct {
	UsernameOrEmail string `json:"username_or_email"`
	Password        string `json:"password"`
}

func (req *LoginRequest) Build(c common.GinI) error {
	return bindRequestBody(c, req)
}

func (req LoginRequest) Validate() error {
	if req.UsernameOrEmail == "" || req.Password == "" {
		return common.ErrAllFieldsRequired
	}

	return nil
}

/*---------------------
//    Create User
--------------------*/

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`

	// User Detail
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (req *CreateUserRequest) Build(c common.GinI) error {
	return bindRequestBody(c, req)
}

func (req CreateUserRequest) Validate() error {
	return validateUsernameEmailAndPassword(req.Username, req.Email, req.Password)
}

/*--------------------
//     Get User
//------------------*/

type GetUserRequest struct {
	UserID int `json:"user_id"`
}

func (req *GetUserRequest) Build(c common.GinI) error {
	req.UserID = c.GetInt(contextUserIDKey)
	return nil
}

func (req GetUserRequest) Validate() error {
	if req.UserID == 0 {
		return common.ErrAllFieldsRequired
	}
	return nil
}

/*--------------------
//    Update User
//------------------*/

type UpdateUserRequest struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`

	// User Detail
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}

func (req *UpdateUserRequest) Build(c common.GinI) error {
	if err := bindRequestBody(c, req); err != nil {
		return err
	}

	req.UserID = c.GetInt(contextUserIDKey)

	return nil
}

func (req UpdateUserRequest) Validate() error {
	if req.UserID == 0 || (req.Email == "" && req.Username == "" && req.FirstName == nil && req.LastName == nil) {
		return common.ErrAllFieldsRequired
	}

	if req.Email != "" && !validEmailRegex.MatchString(req.Email) {
		return common.ErrInvalidEmailFormat
	}

	if req.Username != "" {
		if len(req.Username) < usernameMinLength || len(req.Username) > usernameMaxLength {
			return common.ErrInvalidUsernameLength(usernameMinLength, usernameMaxLength)
		}
	}

	return nil
}

/*--------------------
//    Delete User
//------------------*/

type DeleteUserRequest struct {
	UserID int `json:"user_id"`
}

func (req *DeleteUserRequest) Build(c common.GinI) error {
	req.UserID = c.GetInt(contextUserIDKey)
	return nil
}

func (req DeleteUserRequest) Validate() error {
	if req.UserID == 0 {
		return common.ErrAllFieldsRequired
	}
	return nil
}

/*--------------------
//    Search Users
//------------------*/

type SearchUsersRequest struct {
	Username string `json:"username"`
	Page     int    `json:"page"`
	PerPage  int    `json:"per_page"`
}

func (req *SearchUsersRequest) Build(c common.GinI) error {
	var (
		err            = error(nil)
		defaultPage    = "0"
		defaultPerPage = "10"
	)

	req.Username = c.Query("username")

	req.Page, err = strconv.Atoi(c.DefaultQuery("page", defaultPage))
	if err != nil {
		return common.ErrInvalidValue("page")
	}

	req.PerPage, err = strconv.Atoi(c.DefaultQuery("per_page", defaultPerPage))
	if err != nil {
		return common.ErrInvalidValue("per_page")
	}

	return nil
}

func (req SearchUsersRequest) Validate() error {
	if req.Page < 0 || req.PerPage <= 0 {
		return common.ErrAllFieldsRequired
	}

	return nil
}

/*-----------------------
//    Change Password
//---------------------*/

type ChangePasswordRequest struct {
	UserID         int    `json:"user_id"`
	OldPassword    string `json:"old_password"`
	NewPassword    string `json:"new_password"`
	RepeatPassword string `json:"repeat_password"`
}

func (req *ChangePasswordRequest) Build(c common.GinI) error {
	if err := bindRequestBody(c, req); err != nil {
		return err
	}

	req.UserID = c.GetInt(contextUserIDKey)

	return nil
}

func (req ChangePasswordRequest) Validate() error {
	if req.UserID == 0 || req.OldPassword == "" || req.NewPassword == "" || req.RepeatPassword == "" {
		return common.ErrAllFieldsRequired
	}

	if len(req.NewPassword) < passwordMinLength || len(req.NewPassword) > passwordMaxLength {
		return common.ErrInvalidPasswordLength(passwordMinLength, passwordMaxLength)
	}

	if req.NewPassword != req.RepeatPassword {
		return common.ErrPasswordsDontMatch
	}

	return nil
}

/*------------------------
//   Create User Post
//----------------------*/

type CreateUserPostRequest struct {
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (req *CreateUserPostRequest) Build(c common.GinI) error {
	err := bindRequestBody(c, req)
	if err != nil {
		return err
	}

	req.UserID = c.GetInt(contextUserIDKey)

	return nil
}

func (req CreateUserPostRequest) Validate() error {
	if req.UserID == 0 || req.Title == "" {
		return common.ErrAllFieldsRequired
	}
	return nil
}

/*--------------
//   Helpers
/-------------*/

func bindRequestBody(c common.GinI, request Request) error {
	if err := c.ShouldBindJSON(&request); err != nil {
		return common.Wrap(err.Error(), common.ErrBindingRequest)
	}
	return nil
}

func validateUsernameEmailAndPassword(username, email, password string) error {
	if email == "" || username == "" || password == "" {
		return common.ErrAllFieldsRequired
	}

	if !validEmailRegex.MatchString(email) {
		return common.ErrInvalidEmailFormat
	}

	if len(username) < usernameMinLength || len(username) > usernameMaxLength {
		return common.ErrInvalidUsernameLength(usernameMinLength, usernameMaxLength)
	}

	if len(password) < passwordMinLength || len(password) > passwordMaxLength {
		return common.ErrInvalidPasswordLength(passwordMinLength, passwordMaxLength)
	}

	return nil
}
