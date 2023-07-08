package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/utils"
)

func (s *Service) Login(userCredentials entities.UserCredentials) (entities.LoginResponse, error) {

	// Transform UserCredentials entity to model
	userModel := s.Codec.FromUserCredentialsToUserModel(userCredentials)

	// Get user from database
	user, err := s.Repository.GetUser(userModel)
	if err != nil {
		return entities.LoginResponse{}, s.ErrorsMapper.Map(err)
	}

	// Check if passwords match, if not return error
	if user.Password != utils.Hash(user.Email, userCredentials.Password) {
		return entities.LoginResponse{}, s.ErrorsMapper.Map(entities.ErrWrongPassword)
	}

	// Transform user model to entity
	userEntity := s.Codec.FromUserModelToEntities(user)

	// Generate and return token
	return entities.LoginResponse{
		Token: auth.GenerateToken(userEntity, s.Config.JWT.SESSION_DURATION_DAYS, s.Config.JWT.SECRET),
	}, nil
}
