package service

import (
	"fmt"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	customErrors "github.com/gilperopiola/go-rest-example/pkg/errors"
	"github.com/gilperopiola/go-rest-example/pkg/handlers"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/requests"
	"github.com/gilperopiola/go-rest-example/pkg/responses"
)

// CreateUser is an admins only endpoint
func (s *Service) CreateUser(createUserRequest requests.CreateUserRequest) (responses.CreateUserResponse, error) {
	user := handlers.New(createUserRequest.ToUserModel())

	if user.Exists(s.Repository) {
		return responses.CreateUserResponse{}, common.Wrap(fmt.Errorf("user.Exists"), customErrors.ErrUsernameOrEmailAlreadyInUse)
	}

	user.HashPassword()

	if err := user.Create(s.Repository); err != nil {
		return responses.CreateUserResponse{}, common.Wrap(fmt.Errorf("user.Create"), err)
	}

	return responses.CreateUserResponse{User: user.ToEntity()}, nil
}

func (s *Service) GetUser(getUserRequest requests.GetUserRequest) (responses.GetUserResponse, error) {
	user := handlers.New(getUserRequest.ToUserModel())

	if err := user.Get(s.Repository, repository.WithoutDeleted); err != nil {
		return responses.GetUserResponse{}, common.Wrap(fmt.Errorf("user.Get"), err)
	}

	return responses.GetUserResponse{User: user.ToEntity()}, nil
}

func (s *Service) UpdateUser(updateUserRequest requests.UpdateUserRequest) (responses.UpdateUserResponse, error) {
	user := handlers.New(updateUserRequest.ToUserModel())

	if user.Exists(s.Repository) {
		return responses.UpdateUserResponse{}, common.Wrap(fmt.Errorf("user.Exists"), customErrors.ErrUsernameOrEmailAlreadyInUse)
	}

	if err := user.Get(s.Repository, repository.WithoutDeleted); err != nil {
		return responses.UpdateUserResponse{}, common.Wrap(fmt.Errorf("user.Get"), err)
	}

	user.OverwriteFields(updateUserRequest.Username, updateUserRequest.Email, "")

	if err := user.Update(s.Repository); err != nil {
		return responses.UpdateUserResponse{}, common.Wrap(fmt.Errorf("user.Update"), err)
	}

	return responses.UpdateUserResponse{User: user.ToEntity()}, nil
}

func (s *Service) DeleteUser(deleteUserRequest requests.DeleteUserRequest) (responses.DeleteUserResponse, error) {
	user := handlers.New(deleteUserRequest.ToUserModel())

	// This returns an error if the user is already deleted
	if err := user.Delete(s.Repository); err != nil {
		return responses.DeleteUserResponse{}, common.Wrap(fmt.Errorf("user.Delete"), err)
	}

	return responses.DeleteUserResponse{User: user.ToEntity()}, nil
}
