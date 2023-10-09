package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/utils"
)

// CreateUser is an admins only endpoint
func (s *Service) CreateUser(createUserRequest entities.CreateUserRequest) (entities.CreateUserResponse, error) {

	// Codec functions
	var (
		toModel  = s.Codec.FromCreateUserRequestToUserModel
		toEntity = s.Codec.FromUserModelToEntities
	)

	// Validate user doesn't exist
	if s.Repository.UserExists(createUserRequest.Email, createUserRequest.Username) {
		return entities.CreateUserResponse{}, entities.ErrUsernameOrEmailAlreadyInUse
	}

	// Hash password
	hashedPassword := utils.Hash(createUserRequest.Email, createUserRequest.Password)

	// Transform request to model
	userToCreate := toModel(createUserRequest, hashedPassword)

	// Create user model on the database
	createdUser, err := s.Repository.CreateUser(userToCreate)
	if err != nil {
		return entities.CreateUserResponse{}, s.ErrorsMapper.Map(err)
	}

	// Return response
	return entities.CreateUserResponse{User: toEntity(createdUser)}, nil
}

func (s *Service) GetUser(getUserRequest entities.GetUserRequest) (entities.GetUserResponse, error) {

	// Codec functions
	var (
		toModel  = s.Codec.FromGetUserRequestToUserModel
		toEntity = s.Codec.FromUserModelToEntities
	)

	// Create userToGet model for DB searching
	userToGet := toModel(getUserRequest)

	// Get user from database
	userToGet, err := s.Repository.GetUser(userToGet, repository.WithoutDeleted)
	if err != nil {
		return entities.GetUserResponse{}, s.ErrorsMapper.Map(err)
	}

	// Return user
	return entities.GetUserResponse{User: toEntity(userToGet)}, nil
}

func (s *Service) UpdateUser(updateUserRequest entities.UpdateUserRequest) (entities.UpdateUserResponse, error) {

	// Codec functions
	var (
		toModel  = s.Codec.FromUpdateUserRequestToUserModel
		toEntity = s.Codec.FromUserModelToEntities
	)

	// Check if username and/or email are available
	if s.Repository.UserExists(updateUserRequest.Email, updateUserRequest.Username) {
		return entities.UpdateUserResponse{}, entities.ErrUsernameOrEmailAlreadyInUse
	}

	// Get user from database
	userToUpdate, err := s.Repository.GetUser(toModel(updateUserRequest), repository.WithoutDeleted)
	if err != nil {
		return entities.UpdateUserResponse{}, s.ErrorsMapper.Map(err)
	}

	// Replace fields
	userToUpdate.OverwriteFields(updateUserRequest.Username, updateUserRequest.Email)

	// Update user on the DB
	if userToUpdate, err = s.Repository.UpdateUser(userToUpdate); err != nil {
		return entities.UpdateUserResponse{}, s.ErrorsMapper.Map(err)
	}

	// Return user
	return entities.UpdateUserResponse{User: toEntity(userToUpdate)}, nil
}

func (s *Service) DeleteUser(deleteUserRequest entities.DeleteUserRequest) (entities.DeleteUserResponse, error) {

	// Codec functions
	var (
		toEntity = s.Codec.FromUserModelToEntities
	)

	// Set the user's Deleted field to true
	// This returns an error if the user is already deleted
	userModel, err := s.Repository.DeleteUser(deleteUserRequest.ID)
	if err != nil {
		return entities.DeleteUserResponse{}, s.ErrorsMapper.Map(err)
	}

	// Return user
	return entities.DeleteUserResponse{User: toEntity(userModel)}, nil
}
