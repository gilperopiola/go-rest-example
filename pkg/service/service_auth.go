package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/utils"
)

func (s *Service) Signup(signupRequest entities.SignupRequest) (entities.SignupResponse, error) {

	// Validate user doesn't exist
	if s.Repository.UserExists(signupRequest.Email, signupRequest.Username, false) {
		return entities.SignupResponse{}, s.ErrorsMapper.Map(entities.ErrUsernameOrEmailAlreadyInUse)
	}

	// Hash password
	hashedPassword := utils.Hash(signupRequest.Email, signupRequest.Password)

	// Transform request to model
	userToSignup := s.Codec.FromSignupRequestToUserModel(signupRequest, hashedPassword)

	// Create user model on the database
	createdUser, err := s.Repository.CreateUser(userToSignup)
	if err != nil {
		return entities.SignupResponse{}, s.ErrorsMapper.Map(err)
	}

	// Return response
	return entities.SignupResponse{
		User: s.Codec.FromUserModelToEntities(createdUser), // Transform user to entities
	}, nil
}

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
