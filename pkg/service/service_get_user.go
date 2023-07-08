package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/models"
)

func (s *Service) GetUser(getUserRequest entities.GetUserRequest) (entities.GetUserResponse, error) {

	// Create userModel model for DB searching
	userModel := models.User{ID: getUserRequest.ID}

	// Get user from database
	userModel, err := s.Repository.GetUser(userModel)
	if err != nil {
		return entities.GetUserResponse{}, s.ErrorsMapper.Map(err)
	}

	// Transform user model to entity
	userEntity := s.Codec.FromUserModelToEntities(userModel)

	// Return user
	return entities.GetUserResponse{User: userEntity}, nil
}
