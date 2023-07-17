package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/models"
)

func (s *Service) UpdateUser(updateUserRequest entities.UpdateUserRequest) (entities.UpdateUserResponse, error) {

	// Check if username and/or email are available
	if s.Repository.UserExists(updateUserRequest.Email, updateUserRequest.Username, false) {
		return entities.UpdateUserResponse{}, s.ErrorsMapper.Map(entities.ErrUsernameOrEmailAlreadyInUse)
	}

	// If they are available, create userToUpdate model for DB searching
	userToUpdate := models.User{ID: updateUserRequest.ID}

	// Get user from database
	userToUpdate, err := s.Repository.GetUser(userToUpdate, true)
	if err != nil {
		return entities.UpdateUserResponse{}, s.ErrorsMapper.Map(err)
	}

	// Replace fields
	userToUpdate.FillFields(updateUserRequest.Username, updateUserRequest.Email)

	// Update user on the DB
	if userToUpdate, err = s.Repository.UpdateUser(userToUpdate); err != nil {
		return entities.UpdateUserResponse{}, s.ErrorsMapper.Map(err)
	}

	// Return user
	return entities.UpdateUserResponse{
		User: s.Codec.FromUserModelToEntities(userToUpdate), // Transform user model to entity
	}, nil
}
