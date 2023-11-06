package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/requests"
	"github.com/gilperopiola/go-rest-example/pkg/common/responses"
	"github.com/gilperopiola/go-rest-example/pkg/repository/options"
)

/*-----------------------
//       SIGNUP
//---------------------*/

func (s *service) Signup(request *requests.SignupRequest) (responses.SignupResponse, error) {
	user := request.ToUserModel()

	user.HashPassword(s.config.Auth.HashSalt)

	if err := user.Create(s.repository); err != nil {
		return responses.SignupResponse{}, common.Wrap("Signup: user.Create", err)
	}

	return responses.SignupResponse{User: user.ToResponseModel()}, nil
}

/*---------------------
//       LOGIN
//-------------------*/

func (s *service) Login(request *requests.LoginRequest) (responses.LoginResponse, error) {
	user := request.ToUserModel()

	// Get user
	if err := user.Get(s.repository, options.WithoutDeleted()); err != nil {
		return responses.LoginResponse{}, common.Wrap("Login: user.Get", err)
	}

	// Check password
	if !user.PasswordMatches(request.Password, s.config.Auth.HashSalt) {
		return responses.LoginResponse{}, common.Wrap("Login: !user.PasswordMatches", common.ErrWrongPassword)
	}

	// Generate token
	tokenString, err := user.GenerateTokenString(s.auth)
	if err != nil {
		return responses.LoginResponse{}, common.Wrap("Login: user.GenerateTokenString", common.ErrUnauthorized)
	}

	return responses.LoginResponse{Token: tokenString}, nil
}
