package requests

import (
	"regexp"
	"strconv"
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"
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

// bindRequestBody just binds the request body to the request struct
func bindRequestBody(c common.GinI, request interface{}) error {
	if err := c.ShouldBindJSON(&request); err != nil {
		return common.Wrap(err.Error(), common.ErrBindingRequest)
	}
	return nil
}

/*---------------------
//    CREATE USER
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

func (req CreateUserRequest) Validate() error {
	return validateUsernameEmailAndPassword(req.Username, req.Email, req.Password)
}

func (req *CreateUserRequest) Build(c common.GinI) error {
	return bindRequestBody(c, req)
}

func (r *CreateUserRequest) ToUserModel() models.User {
	return models.User{
		Email:    r.Email,
		Username: r.Username,
		Password: r.Password,
		Deleted:  false,
		Details: models.UserDetail{
			FirstName: r.FirstName,
			LastName:  r.LastName,
		},
		IsAdmin:   r.IsAdmin,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

/*--------------------
//     GET USER
//------------------*/

type GetUserRequest struct {
	UserID int `json:"user_id"`
}

func (req GetUserRequest) Validate() error {
	if req.UserID == 0 {
		return common.ErrAllFieldsRequired
	}
	return nil
}

func (req *GetUserRequest) Build(c common.GinI) error {
	req.UserID = c.GetInt(contextUserIDKey)
	return nil
}

func (r *GetUserRequest) ToUserModel() models.User {
	return models.User{ID: r.UserID}
}

/*--------------------
//    UPDATE USER
//------------------*/

type UpdateUserRequest struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`

	// User Detail
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
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

func (req *UpdateUserRequest) Build(c common.GinI) error {
	if err := bindRequestBody(c, req); err != nil {
		return err
	}

	req.UserID = c.GetInt(contextUserIDKey)

	return nil
}

func (r *UpdateUserRequest) ToUserModel() models.User {
	firstName, lastName := "", ""
	if r.FirstName != nil {
		firstName = *r.FirstName
	}
	if r.LastName != nil {
		lastName = *r.LastName
	}

	return models.User{
		ID:       r.UserID,
		Username: r.Username,
		Email:    r.Email,
		Details: models.UserDetail{
			FirstName: firstName,
			LastName:  lastName,
		},
	}
}

/*--------------------
//    DELETE USER
//------------------*/

type DeleteUserRequest struct {
	UserID int `json:"user_id"`
}

func (req DeleteUserRequest) Validate() error {
	if req.UserID == 0 {
		return common.ErrAllFieldsRequired
	}

	return nil
}

func (req *DeleteUserRequest) Build(c common.GinI) error {
	req.UserID = c.GetInt(contextUserIDKey)
	return nil
}

func (r *DeleteUserRequest) ToUserModel() models.User {
	return models.User{ID: r.UserID}
}

/*--------------------
//    SEARCH USERS
//------------------*/

type SearchUsersRequest struct {
	Username string `json:"username"`
	Page     int    `json:"page"`
	PerPage  int    `json:"per_page"`
}

func (req SearchUsersRequest) Validate() error {
	if req.Page < 0 || req.PerPage <= 0 {
		return common.ErrAllFieldsRequired
	}

	return nil
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

func (r *SearchUsersRequest) ToUserModel() models.User {
	return models.User{Username: r.Username}
}

/*-----------------------
//    CHANGE PASSWORD
//---------------------*/

type ChangePasswordRequest struct {
	UserID         int    `json:"user_id"`
	OldPassword    string `json:"old_password"`
	NewPassword    string `json:"new_password"`
	RepeatPassword string `json:"repeat_password"`
}

func (req ChangePasswordRequest) Validate() error {
	if req.UserID == 0 {
		return common.ErrAllFieldsRequired
	}

	if req.OldPassword == "" || req.NewPassword == "" || req.RepeatPassword == "" {
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

func (req *ChangePasswordRequest) Build(c common.GinI) error {
	err := bindRequestBody(c, req)
	if err != nil {
		return err
	}

	req.UserID = c.GetInt(contextUserIDKey)

	return nil
}

func (r *ChangePasswordRequest) ToUserModel() models.User {
	return models.User{
		ID:       r.UserID,
		Password: r.OldPassword,
	}
}

/*------------------------
//    CREATE USER POST
//----------------------*/

type CreateUserPostRequest struct {
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (req CreateUserPostRequest) Validate() error {
	if req.UserID == 0 || req.Title == "" {
		return common.ErrAllFieldsRequired
	}
	return nil
}

func (req *CreateUserPostRequest) Build(c common.GinI) error {
	err := bindRequestBody(c, req)
	if err != nil {
		return err
	}

	req.UserID = c.GetInt(contextUserIDKey)

	return nil
}

func (r *CreateUserPostRequest) ToUserPostModel() models.UserPost {
	return models.UserPost{
		UserID: r.UserID,
		Title:  r.Title,
		Body:   r.Body,
	}
}

/*-----------------------
//       HELPERS
//---------------------*/

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
