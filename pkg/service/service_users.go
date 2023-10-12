package service

import (
	"fmt"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	customErrors "github.com/gilperopiola/go-rest-example/pkg/errors"
	"github.com/gilperopiola/go-rest-example/pkg/handlers"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
)

// CreateUser is an admins only endpoint
func (s *Service) CreateUser(createUserRequest common.CreateUserRequest) (common.CreateUserResponse, error) {
	user := handlers.New(createUserRequest.ToUserModel())

	if user.Exists(s.Repository) {
		return common.CreateUserResponse{}, common.Wrap(fmt.Errorf("user.Exists"), customErrors.ErrUsernameOrEmailAlreadyInUse)
	}

	user.HashPassword()

	if err := user.Create(s.Repository); err != nil {
		return common.CreateUserResponse{}, common.Wrap(fmt.Errorf("user.Create"), err)
	}

	return common.CreateUserResponse{User: user.ToEntity()}, nil
}

func (s *Service) GetUser(getUserRequest common.GetUserRequest) (common.GetUserResponse, error) {
	user := handlers.New(getUserRequest.ToUserModel())

	if err := user.Get(s.Repository, repository.WithoutDeleted); err != nil {
		return common.GetUserResponse{}, common.Wrap(fmt.Errorf("user.Get"), err)
	}

	return common.GetUserResponse{User: user.ToEntity()}, nil
}

func (s *Service) UpdateUser(updateUserRequest common.UpdateUserRequest) (common.UpdateUserResponse, error) {
	user := handlers.New(updateUserRequest.ToUserModel())

	if user.Exists(s.Repository) {
		return common.UpdateUserResponse{}, common.Wrap(fmt.Errorf("user.Exists"), customErrors.ErrUsernameOrEmailAlreadyInUse)
	}

	if err := user.Get(s.Repository, repository.WithoutDeleted); err != nil {
		return common.UpdateUserResponse{}, common.Wrap(fmt.Errorf("user.Get"), err)
	}

	user.OverwriteFields(updateUserRequest.Username, updateUserRequest.Email, "")

	if err := user.Update(s.Repository); err != nil {
		return common.UpdateUserResponse{}, common.Wrap(fmt.Errorf("user.Update"), err)
	}

	return common.UpdateUserResponse{User: user.ToEntity()}, nil
}

func (s *Service) DeleteUser(deleteUserRequest common.DeleteUserRequest) (common.DeleteUserResponse, error) {
	user := handlers.New(deleteUserRequest.ToUserModel())

	// This returns an error if the user is already deleted
	if err := user.Delete(s.Repository); err != nil {
		return common.DeleteUserResponse{}, common.Wrap(fmt.Errorf("user.Delete"), err)
	}

	return common.DeleteUserResponse{User: user.ToEntity()}, nil
}
