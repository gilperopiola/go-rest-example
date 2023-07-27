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
