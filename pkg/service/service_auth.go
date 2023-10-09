package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/utils"
)

func (s *Service) Signup(signupRequest entities.SignupRequest) (entities.SignupResponse, error) {

	// Codec functions
	var (
		toModel  = s.Codec.FromSignupRequestToUserModel
		toEntity = s.Codec.FromUserModelToEntities
	)

	// Validate user doesn't exist
	if s.Repository.UserExists(signupRequest.Email, signupRequest.Username, false) {
		return entities.SignupResponse{}, entities.ErrUsernameOrEmailAlreadyInUse
	}

	// Hash password
	hashedPassword := utils.Hash(signupRequest.Email, signupRequest.Password)

	// Transform request to model
	userToSignup := toModel(signupRequest, hashedPassword)

	// Create user model on the database
	createdUser, err := s.Repository.CreateUser(userToSignup)
	if err != nil {
		return entities.SignupResponse{}, s.ErrorsMapper.Map(err)
	}

	// Return response
	return entities.SignupResponse{User: toEntity(createdUser)}, nil
}

func (s *Service) Login(loginRequest entities.LoginRequest) (entities.LoginResponse, error) {

	// Codec functions
	var (
		toModel  = s.Codec.FromLoginRequestToUserModel
		toEntity = s.Codec.FromUserModelToEntities
	)

	// Transform LoginRequest to user model
	userToLogin := toModel(loginRequest)

	// Get user from database
	userToLogin, err := s.Repository.GetUser(userToLogin, true)
	if err != nil {
		return entities.LoginResponse{}, s.ErrorsMapper.Map(err)
	}

	// Check if passwords match, if not return error
	if !userToLogin.PasswordMatches(loginRequest.Password) {
		return entities.LoginResponse{}, entities.ErrWrongPassword
	}

	// Transform user model to entity
	userEntity := toEntity(userToLogin)

	// Set the appropriate role
	authRole := entities.UserRole
	if userEntity.IsAdmin {
		authRole = entities.AdminRole
	}

	tokenString, err := s.Auth.GenerateToken(userEntity, authRole)
	if err != nil {
		return entities.LoginResponse{}, entities.ErrUnauthorized
	}

	// Return generated token on the response
	return entities.LoginResponse{Token: tokenString}, nil
}
