package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/requests"
	"github.com/gilperopiola/go-rest-example/pkg/responses"
	"github.com/gilperopiola/go-rest-example/pkg/utils"
)

// CreateUser is an admins only endpoint
func (s *Service) CreateUser(createUserRequest requests.CreateUserRequest) (responses.CreateUserResponse, error) {

	// Check if username or email are already in use
	if s.Repository.UserExists(createUserRequest.Email, createUserRequest.Username) {
		return responses.CreateUserResponse{}, entities.ErrUsernameOrEmailAlreadyInUse
	}

	createUserRequest.Password = utils.Hash(createUserRequest.Email, createUserRequest.Password)

	userFromDB, err := s.Repository.CreateUser(createUserRequest.ToUserModel())
	if err != nil {
		return responses.CreateUserResponse{}, s.ErrorsMapper.Map(err)
	}

	// Return response
	return responses.CreateUserResponse{User: userFromDB.ToEntity()}, nil
}

func (s *Service) GetUser(getUserRequest requests.GetUserRequest) (responses.GetUserResponse, error) {
	userFromDB, err := s.Repository.GetUser(getUserRequest.ToUserModel(), repository.WithoutDeleted)
	if err != nil {
		return responses.GetUserResponse{}, s.ErrorsMapper.Map(err)
	}

	return responses.GetUserResponse{User: userFromDB.ToEntity()}, nil
}

func (s *Service) UpdateUser(updateUserRequest requests.UpdateUserRequest) (responses.UpdateUserResponse, error) {

	// Check if username or email are already in use
	if s.Repository.UserExists(updateUserRequest.Email, updateUserRequest.Username) {
		return responses.UpdateUserResponse{}, entities.ErrUsernameOrEmailAlreadyInUse
	}

	// Get user from database
	userFromDB, err := s.Repository.GetUser(updateUserRequest.ToUserModel(), repository.WithoutDeleted)
	if err != nil {
		return responses.UpdateUserResponse{}, s.ErrorsMapper.Map(err)
	}

	// Replace fields with the new ones
	userFromDB.OverwriteFields(updateUserRequest.Username, updateUserRequest.Email)

	if userFromDB, err = s.Repository.UpdateUser(userFromDB); err != nil {
		return responses.UpdateUserResponse{}, s.ErrorsMapper.Map(err)
	}

	return responses.UpdateUserResponse{User: userFromDB.ToEntity()}, nil
}

func (s *Service) DeleteUser(deleteUserRequest requests.DeleteUserRequest) (responses.DeleteUserResponse, error) {
	// Set the user's Deleted field to true
	// This returns an error if the user is already deleted
	userFromDB, err := s.Repository.DeleteUser(deleteUserRequest.ID)
	if err != nil {
		return responses.DeleteUserResponse{}, s.ErrorsMapper.Map(err)
	}

	// Return user
	return responses.DeleteUserResponse{User: userFromDB.ToEntity()}, nil
}
