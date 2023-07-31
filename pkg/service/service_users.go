package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
)

func (s *Service) GetUser(getUserRequest entities.GetUserRequest) (entities.GetUserResponse, error) {

	// Create userToGet model for DB searching
	userToGet := s.Codec.FromGetUserRequestToUserModel(getUserRequest)

	// Get user from database
	userToGet, err := s.Repository.GetUser(userToGet, true)
	if err != nil {
		return entities.GetUserResponse{}, s.ErrorsMapper.Map(err)
	}

	// Return user
	return entities.GetUserResponse{
		User: s.Codec.FromUserModelToEntities(userToGet),
	}, nil
}

func (s *Service) UpdateUser(updateUserRequest entities.UpdateUserRequest) (entities.UpdateUserResponse, error) {

	// Check if username and/or email are available
	if s.Repository.UserExists(updateUserRequest.Email, updateUserRequest.Username, false) {
		return entities.UpdateUserResponse{}, s.ErrorsMapper.Map(entities.ErrUsernameOrEmailAlreadyInUse)
	}

	// If they are available, create userToUpdate model for DB searching
	userToUpdate := s.Codec.FromUpdateUserRequestToUserModel(updateUserRequest)

	// Get user from database
	userToUpdate, err := s.Repository.GetUser(userToUpdate, true)
	if err != nil {
		return entities.UpdateUserResponse{}, s.ErrorsMapper.Map(err)
	}

	// Replace fields
	userToUpdate.Fill(updateUserRequest.Username, updateUserRequest.Email)

	// Update user on the DB
	if userToUpdate, err = s.Repository.UpdateUser(userToUpdate); err != nil {
		return entities.UpdateUserResponse{}, s.ErrorsMapper.Map(err)
	}

	// Return user
	return entities.UpdateUserResponse{
		User: s.Codec.FromUserModelToEntities(userToUpdate), // Transform user model to entity
	}, nil
}

func (s *Service) DeleteUser(deleteUserRequest entities.DeleteUserRequest) (entities.DeleteUserResponse, error) {

	// Set the user's Deleted field to true
	userModel, err := s.Repository.DeleteUser(deleteUserRequest.ID)
	if err != nil {
		return entities.DeleteUserResponse{}, s.ErrorsMapper.Map(err)
	}

	// Transform user model to entity
	userEntity := s.Codec.FromUserModelToEntities(userModel)

	// Return user
	return entities.DeleteUserResponse{User: userEntity}, nil
}
