package requests

import (
	"strconv"

	"github.com/gilperopiola/go-rest-example/pkg/common"
)

const (
	contextUserIDKey = "UserID"
	pathUserIDKey    = "user_id"
)

/*---------------
//    Signup
//-------------*/

type SignupRequest struct {
	Username       string `json:"username" validate:"required,min=4,max=32"`
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required,min=8,max=64"`
	RepeatPassword string `json:"repeat_password" validate:"required,eqfield=Password"`

	// User Detail
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (req *SignupRequest) Build(c common.GinI) error {
	return bindRequestBody(c, req)
}

/*--------------
//    Login
//------------*/

type LoginRequest struct {
	UsernameOrEmail string `json:"username_or_email" validate:"required"`
	Password        string `json:"password" validate:"required,min=8,max=64"`
}

func (req *LoginRequest) Build(c common.GinI) error {
	return bindRequestBody(c, req)
}

/*---------------------
//    Create User
--------------------*/

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=4,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=64"`
	IsAdmin  bool   `json:"is_admin"`

	// User Detail
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (req *CreateUserRequest) Build(c common.GinI) error {
	return bindRequestBody(c, req)
}

/*--------------------
//     Get User
//------------------*/

type GetUserRequest struct {
	UserID int `json:"user_id" validate:"required,min=1"`
}

func (req *GetUserRequest) Build(c common.GinI) error {
	req.UserID = c.GetInt(contextUserIDKey)
	return nil
}

/*--------------------
//    Update User
//------------------*/

type UpdateUserRequest struct {
	UserID   int    `json:"user_id" validate:"required,min=1"`
	Username string `json:"username" validate:"required_without_all=Email FirstName LastName,omitempty,min=4,max=32"`
	Email    string `json:"email" validate:"required_without_all=Username FirstName LastName,omitempty,email"`

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

/*--------------------
//    Delete User
//------------------*/

type DeleteUserRequest struct {
	UserID int `json:"user_id" validate:"required,min=1"`
}

func (req *DeleteUserRequest) Build(c common.GinI) error {
	req.UserID = c.GetInt(contextUserIDKey)
	return nil
}

/*--------------------
//    Search Users
//------------------*/

type SearchUsersRequest struct {
	Username string `json:"username"`
	Page     int    `json:"page" validate:"omitempty,min=0"`
	PerPage  int    `json:"per_page" validate:"omitempty,min=1,max=100"`
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

/*-----------------------
//    Change Password
//---------------------*/

type ChangePasswordRequest struct {
	UserID         int    `json:"user_id" validate:"required,min=1"`
	OldPassword    string `json:"old_password" validate:"required"`
	NewPassword    string `json:"new_password" validate:"required,min=8,max=64"`
	RepeatPassword string `json:"repeat_password" validate:"required,eqfield=NewPassword"`
}

func (req *ChangePasswordRequest) Build(c common.GinI) error {
	if err := bindRequestBody(c, req); err != nil {
		return err
	}

	req.UserID = c.GetInt(contextUserIDKey)

	return nil
}

/*------------------------
//   Create User Post
//----------------------*/

type CreateUserPostRequest struct {
	UserID int    `json:"user_id" validate:"required,min=1"`
	Title  string `json:"title" validate:"required"`
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
