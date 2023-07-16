package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
)

func (s *Service) DeleteUser(deleteUserRequest entities.DeleteUserRequest) (entities.DeleteUserResponse, error) {

	// Set the user's Deleted = true
	userModel, err := s.Repository.DeleteUser(deleteUserRequest.ID)
	if err != nil {
		return entities.DeleteUserResponse{}, s.ErrorsMapper.Map(err)
	}

	// Transform user model to entity
	userEntity := s.Codec.FromUserModelToEntities(userModel)

	// Return user
	return entities.DeleteUserResponse{User: userEntity}, nil
}
