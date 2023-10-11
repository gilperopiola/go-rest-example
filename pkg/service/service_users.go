package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/handlers"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/requests"
	"github.com/gilperopiola/go-rest-example/pkg/responses"
)

// CreateUser is an admins only endpoint
func (s *Service) CreateUser(createUserRequest requests.CreateUserRequest) (responses.CreateUserResponse, error) {
	userHandler := handlers.New(createUserRequest.ToUserModel())

	if userHandler.Exists(s.Repository) {
		return responses.CreateUserResponse{}, entities.ErrUsernameOrEmailAlreadyInUse
	}

	userHandler.HashPassword()

	if err := userHandler.Create(s.Repository); err != nil {
		return responses.CreateUserResponse{}, s.ErrorsMapper.Map(err)
	}

	return responses.CreateUserResponse{User: userHandler.ToEntity()}, nil
}

func (s *Service) GetUser(getUserRequest requests.GetUserRequest) (responses.GetUserResponse, error) {
	userHandler := handlers.New(getUserRequest.ToUserModel())

	if err := userHandler.Get(s.Repository, repository.WithoutDeleted); err != nil {
		return responses.GetUserResponse{}, s.ErrorsMapper.Map(err)
	}

	return responses.GetUserResponse{User: userHandler.ToEntity()}, nil
}

func (s *Service) UpdateUser(updateUserRequest requests.UpdateUserRequest) (responses.UpdateUserResponse, error) {
	userHandler := handlers.New(updateUserRequest.ToUserModel())

	if userHandler.Exists(s.Repository) {
		return responses.UpdateUserResponse{}, entities.ErrUsernameOrEmailAlreadyInUse
	}

	if err := userHandler.Get(s.Repository, repository.WithoutDeleted); err != nil {
		return responses.UpdateUserResponse{}, s.ErrorsMapper.Map(err)
	}

	// Replace fields with the new ones
	userHandler.OverwriteFields(updateUserRequest.Username, updateUserRequest.Email, "")

	if err := userHandler.Update(s.Repository); err != nil {
		return responses.UpdateUserResponse{}, s.ErrorsMapper.Map(err)
	}

	return responses.UpdateUserResponse{User: userHandler.ToEntity()}, nil
}

func (s *Service) DeleteUser(deleteUserRequest requests.DeleteUserRequest) (responses.DeleteUserResponse, error) {
	userHandler := handlers.New(deleteUserRequest.ToUserModel())

	// This returns an error if the user is already deleted
	if err := userHandler.Delete(s.Repository); err != nil {
		return responses.DeleteUserResponse{}, s.ErrorsMapper.Map(err)
	}

	return responses.DeleteUserResponse{User: userHandler.ToEntity()}, nil
}
