package requests

import (
	"regexp"
	"strconv"
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/config"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"
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

func (r *SignupRequest) ToUserModel(config *config.Config, repository models.RepositoryI) models.User {
	return models.User{
		Username: r.Username,
		Email:    r.Email,
		Password: common.Hash(r.Password, config.Auth.HashSalt),
		Details: models.UserDetail{
			FirstName: r.FirstName,
			LastName:  r.LastName,
		},
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		ModelDependencies: modelDeps(config, repository),
	}
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

func (r *LoginRequest) ToUserModel(config *config.Config, repository models.RepositoryI) models.User {
	out := models.User{
		Password:          r.Password,
		ModelDependencies: modelDeps(config, repository),
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`).MatchString(r.UsernameOrEmail) {
		out.Username = r.UsernameOrEmail
	} else {
		out.Email = r.UsernameOrEmail
	}

	return out
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

func (r *CreateUserRequest) ToUserModel(config *config.Config, repository models.RepositoryI) models.User {
	return models.User{
		Email:    r.Email,
		Username: r.Username,
		Password: common.Hash(r.Password, config.Auth.HashSalt),
		Details: models.UserDetail{
			FirstName: r.FirstName,
			LastName:  r.LastName,
		},
		IsAdmin:           r.IsAdmin,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		ModelDependencies: modelDeps(config, repository),
	}
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

func (r *GetUserRequest) ToUserModel(repository models.RepositoryI) models.User {
	return models.User{
		ID:                r.UserID,
		ModelDependencies: modelDeps(nil, repository),
	}
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

func (r *UpdateUserRequest) ToUserModel(repository models.RepositoryI) models.User {
	return models.User{
		ID:       r.UserID,
		Username: r.Username,
		Email:    r.Email,
		Details: models.UserDetail{
			FirstName: getPtrStrValue(r.FirstName),
			LastName:  getPtrStrValue(r.LastName),
		},
		ModelDependencies: modelDeps(nil, repository),
	}
}

func getPtrStrValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
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

func (r *DeleteUserRequest) ToUserModel(repository models.RepositoryI) models.User {
	return models.User{
		ID:                r.UserID,
		ModelDependencies: modelDeps(nil, repository),
	}
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

	if req.Page, err = strconv.Atoi(c.DefaultQuery("page", defaultPage)); err != nil {
		return common.ErrInvalidValue("page")
	}

	if req.PerPage, err = strconv.Atoi(c.DefaultQuery("per_page", defaultPerPage)); err != nil {
		return common.ErrInvalidValue("per_page")
	}

	return nil
}

func (r *SearchUsersRequest) ToUserModel(repository models.RepositoryI) models.User {
	return models.User{
		Username:          r.Username,
		ModelDependencies: modelDeps(nil, repository),
	}
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

func (r *ChangePasswordRequest) ToUserModel(config *config.Config, repository models.RepositoryI) models.User {
	return models.User{
		ID:                r.UserID,
		Password:          r.OldPassword,
		ModelDependencies: modelDeps(config, repository),
	}
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

func (r *CreateUserPostRequest) ToUserPostModel(repository models.RepositoryI) models.UserPost {
	return models.UserPost{
		UserID:            r.UserID,
		Title:             r.Title,
		Body:              r.Body,
		ModelDependencies: modelDeps(nil, repository),
	}
}
