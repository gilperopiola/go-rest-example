package requests

import (
	"github.com/gilperopiola/go-rest-example/pkg/common"
)

//---------------
//	CREATE USER
//---------------

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`

	// User Detail
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func MakeCreateUserRequest(c GinI) (request CreateUserRequest, err error) {
	if err = request.Build(c); err != nil {
		return CreateUserRequest{}, common.Wrap("makeCreateUserRequest", err)
	}
	if err = request.Validate(); err != nil {
		return CreateUserRequest{}, common.Wrap("makeCreateUserRequest", err)
	}
	return request, nil
}

//--------------
//	GET USER
//--------------

type GetUserRequest struct {
	UserID int `json:"user_id"`
}

func MakeGetUserRequest(c GinI) (request GetUserRequest, err error) {
	if err = request.Build(c); err != nil {
		return GetUserRequest{}, common.Wrap("makeGetUserRequest", err)
	}
	if err = request.Validate(); err != nil {
		return GetUserRequest{}, common.Wrap("makeGetUserRequest", err)
	}
	return request, nil
}

//-----------------
//	UPDATE USER
//-----------------

type UpdateUserRequest struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`

	// User Detail
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}

func MakeUpdateUserRequest(c GinI) (request UpdateUserRequest, err error) {
	if err = request.Build(c); err != nil {
		return UpdateUserRequest{}, common.Wrap("makeUpdateUserRequest", err)
	}
	if err = request.Validate(); err != nil {
		return UpdateUserRequest{}, common.Wrap("makeUpdateUserRequest", err)
	}
	return request, nil
}

//----------------
//	DELETE USER
//----------------

type DeleteUserRequest struct {
	UserID int `json:"user_id"`
}

func MakeDeleteUserRequest(c GinI) (request DeleteUserRequest, err error) {
	if err = request.Build(c); err != nil {
		return DeleteUserRequest{}, common.Wrap("makeDeleteUserRequest", err)
	}
	if err = request.Validate(); err != nil {
		return DeleteUserRequest{}, common.Wrap("makeDeleteUserRequest", err)
	}
	return request, nil
}

//-----------------
//	SEARCH USERS
//-----------------

type SearchUsersRequest struct {
	Username string `json:"username"`
	Page     int    `json:"page"`
	PerPage  int    `json:"per_page"`
}

func MakeSearchUsersRequest(c GinI) (request SearchUsersRequest, err error) {
	if err = request.Build(c); err != nil {
		return SearchUsersRequest{}, common.Wrap("makeSearchUsersRequest", err)
	}
	if err = request.Validate(); err != nil {
		return SearchUsersRequest{}, common.Wrap("makeSearchUsersRequest", err)
	}
	return request, nil
}

//-------------------
//	CHANGE PASSWORD
//-------------------

type ChangePasswordRequest struct {
	UserID         int    `json:"user_id"`
	OldPassword    string `json:"old_password"`
	NewPassword    string `json:"new_password"`
	RepeatPassword string `json:"repeat_password"`
}

func MakeChangePasswordRequest(c GinI) (request ChangePasswordRequest, err error) {
	if err = request.Build(c); err != nil {
		return ChangePasswordRequest{}, common.Wrap("makeChangePasswordRequest", err)
	}
	if err = request.Validate(); err != nil {
		return ChangePasswordRequest{}, common.Wrap("makeChangePasswordRequest", err)
	}
	return request, nil
}

//---------------------
//	CREATE USER POST
//---------------------

type CreateUserPostRequest struct {
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func MakeCreateUserPostRequest(c GinI) (request CreateUserPostRequest, err error) {
	if err = request.Build(c); err != nil {
		return CreateUserPostRequest{}, common.Wrap("makeCreateUserPostRequest", err)
	}
	if err = request.Validate(); err != nil {
		return CreateUserPostRequest{}, common.Wrap("makeCreateUserPostRequest", err)
	}
	return request, nil
}
