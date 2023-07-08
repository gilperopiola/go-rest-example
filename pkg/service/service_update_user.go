package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/models"
)

func (s *Service) UpdateUser(updateUserRequest entities.UpdateUserRequest) (entities.UpdateUserResponse, error) {

	// Check if username or email are available
	if s.Repository.UserExists(updateUserRequest.Email, updateUserRequest.Username) {
		return entities.UpdateUserResponse{}, s.ErrorsMapper.Map(entities.ErrUsernameOrEmailAlreadyInUse)
	}

	// If they are available, create userModel model for DB searching
	userModel := models.User{ID: updateUserRequest.ID}

	// Get user from database
	userModel, err := s.Repository.GetUser(userModel)
	if err != nil {
		return entities.UpdateUserResponse{}, s.ErrorsMapper.Map(err)
	}

	// Replace fields
	if updateUserRequest.Username != "" {
		userModel.Username = updateUserRequest.Username
	}

	if updateUserRequest.Email != "" {
		userModel.Email = updateUserRequest.Email
	}

	// Update user on the DB
	if userModel, err = s.Repository.UpdateUser(userModel); err != nil {
		return entities.UpdateUserResponse{}, s.ErrorsMapper.Map(err)
	}

	// Transform user model to entity
	userEntity := s.Codec.FromUserModelToEntities(userModel)

	// Return user
	return entities.UpdateUserResponse{User: userEntity}, nil
}
