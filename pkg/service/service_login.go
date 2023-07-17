package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/utils"
)

func (s *Service) Login(userCredentials entities.UserCredentials) (entities.LoginResponse, error) {
	var err error

	// Transform UserCredentials entity to model
	userToLogin := s.Codec.FromUserCredentialsToUserModel(userCredentials)

	// Get user from database
	userToLogin, err = s.Repository.GetUser(userToLogin, true)
	if err != nil {
		return entities.LoginResponse{}, s.ErrorsMapper.Map(err)
	}

	// Check if passwords match, if not return error
	if !passwordsMatch(userToLogin.Password, userCredentials.Password, userToLogin.Email) {
		return entities.LoginResponse{}, s.ErrorsMapper.Map(entities.ErrWrongPassword)
	}

	// Transform user model to entity
	userEntity := s.Codec.FromUserModelToEntities(userToLogin)

	// Return generated token on the response
	return entities.LoginResponse{
		Token: auth.GenerateToken(userEntity, s.Config.JWT.SESSION_DURATION_DAYS, s.Config.JWT.SECRET),
	}, nil
}

func passwordsMatch(password, otherPassword, email string) bool {
	return password == utils.Hash(email, otherPassword)
}
