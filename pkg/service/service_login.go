package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/utils"
)

func (s *Service) Login(loginRequest entities.LoginRequest) (entities.LoginResponse, error) {

	// Transform LoginRequest to user model
	userToLogin := s.Codec.FromLoginRequestToUserModel(loginRequest)

	// Get user from database
	userToLogin, err := s.Repository.GetUser(userToLogin, true)
	if err != nil {
		return entities.LoginResponse{}, s.ErrorsMapper.Map(err)
	}

	// Check if passwords match, if not return error
	if !passwordsMatch(userToLogin.Password, loginRequest.Password, userToLogin.Email) {
		return entities.LoginResponse{}, s.ErrorsMapper.Map(entities.ErrWrongPassword)
	}

	// Transform user model to entity
	userEntity := s.Codec.FromUserModelToEntities(userToLogin)

	// Return generated token on the response
	return entities.LoginResponse{Token: s.Auth.GenerateToken(userEntity)}, nil
}

func passwordsMatch(password, otherPassword, email string) bool {
	return password == utils.Hash(email, otherPassword)
}
