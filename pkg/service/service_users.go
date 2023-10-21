package service

import (
	"fmt"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	customErrors "github.com/gilperopiola/go-rest-example/pkg/common/errors"
	"github.com/gilperopiola/go-rest-example/pkg/common/requests"
	"github.com/gilperopiola/go-rest-example/pkg/common/responses"
	"github.com/gilperopiola/go-rest-example/pkg/repository/options"
)

//-------------------------
//       CREATE USER
//-------------------------

// CreateUser is an admins only endpoint
func (s *service) CreateUser(createUserRequest requests.CreateUserRequest) (responses.CreateUserResponse, error) {
	user := createUserRequest.ToUserModel()

	if user.Exists(s.repository) {
		return responses.CreateUserResponse{}, common.Wrap(fmt.Errorf("CreateUser: user.Exists"), customErrors.ErrUsernameOrEmailAlreadyInUse)
	}

	user.HashPassword(s.config.JWT.HashSalt)

	if err := user.Create(s.repository); err != nil {
		return responses.CreateUserResponse{}, common.Wrap(fmt.Errorf("CreateUser: user.Create"), err)
	}

	return responses.CreateUserResponse{User: user.ToResponseModel()}, nil
}

//-----------------------
//       GET USER
//-----------------------

func (s *service) GetUser(getUserRequest requests.GetUserRequest) (responses.GetUserResponse, error) {
	user := getUserRequest.ToUserModel()

	if err := user.Get(s.repository, options.WithoutDeleted); err != nil {
		return responses.GetUserResponse{}, common.Wrap(fmt.Errorf("GetUser: user.Get"), err)
	}

	return responses.GetUserResponse{User: user.ToResponseModel()}, nil
}

//--------------------------
//       UPDATE USER
//--------------------------

func (s *service) UpdateUser(updateUserRequest requests.UpdateUserRequest) (responses.UpdateUserResponse, error) {
	user := updateUserRequest.ToUserModel()

	if user.Exists(s.repository) {
		return responses.UpdateUserResponse{}, common.Wrap(fmt.Errorf("UpdateUser: user.Exists"), customErrors.ErrUsernameOrEmailAlreadyInUse)
	}

	if err := user.Get(s.repository, options.WithoutDeleted); err != nil {
		return responses.UpdateUserResponse{}, common.Wrap(fmt.Errorf("UpdateUser: user.Get"), err)
	}

	user.OverwriteFields(updateUserRequest.Username, updateUserRequest.Email, "")
	user.OverwriteDetails(updateUserRequest.FirstName, updateUserRequest.LastName)

	if err := user.Update(s.repository); err != nil {
		return responses.UpdateUserResponse{}, common.Wrap(fmt.Errorf("UpdateUser: user.Update"), err)
	}

	return responses.UpdateUserResponse{User: user.ToResponseModel()}, nil
}

//--------------------------
//       DELETE USER
//--------------------------

func (s *service) DeleteUser(deleteUserRequest requests.DeleteUserRequest) (responses.DeleteUserResponse, error) {
	user := deleteUserRequest.ToUserModel()

	// This returns an error if the user is already deleted
	if err := user.Delete(s.repository); err != nil {
		return responses.DeleteUserResponse{}, common.Wrap(fmt.Errorf("DeleteUser: user.Delete"), err)
	}

	return responses.DeleteUserResponse{User: user.ToResponseModel()}, nil
}

//--------------------------
//      SEARCH USERS
//--------------------------

func (s *service) SearchUsers(searchUsersRequest requests.SearchUsersRequest) (responses.SearchUsersResponse, error) {
	var (
		page    = searchUsersRequest.Page
		perPage = searchUsersRequest.PerPage
	)

	user := searchUsersRequest.ToUserModel()

	users, err := user.Search(s.repository, page, perPage)
	if err != nil {
		return responses.SearchUsersResponse{}, common.Wrap(fmt.Errorf("SearchUsers: user.Search"), err)
	}

	return responses.SearchUsersResponse{
		Users:   users.ToResponseModel(),
		Page:    page,
		PerPage: perPage,
	}, nil
}
