package service

import (
	"fmt"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	customErrors "github.com/gilperopiola/go-rest-example/pkg/common/errors"
	"github.com/gilperopiola/go-rest-example/pkg/common/requests"
	"github.com/gilperopiola/go-rest-example/pkg/common/responses"
	"github.com/gilperopiola/go-rest-example/pkg/repository/options"
)

//-----------------------
//       SIGNUP
//-----------------------

func (s *service) Signup(signupRequest requests.SignupRequest) (responses.SignupResponse, error) {
	user := signupRequest.ToUserModel()

	if user.Exists(s.repository) {
		return responses.SignupResponse{}, common.Wrap(fmt.Errorf("Signup: user.Exists"), customErrors.ErrUsernameOrEmailAlreadyInUse)
	}

	user.HashPassword(s.config.JWT.HashSalt)

	if err := user.Create(s.repository); err != nil {
		return responses.SignupResponse{}, common.Wrap(fmt.Errorf("Signup: user.Create"), err)
	}

	return responses.SignupResponse{User: user.ToResponseModel()}, nil
}

//---------------------
//       LOGIN
//---------------------

func (s *service) Login(loginRequest requests.LoginRequest) (responses.LoginResponse, error) {
	user := loginRequest.ToUserModel()

	if err := user.Get(s.repository, options.WithoutDeleted); err != nil {
		return responses.LoginResponse{}, common.Wrap(fmt.Errorf("Login: user.Get"), err)
	}

	if !user.PasswordMatches(loginRequest.Password, s.config.JWT.HashSalt) {
		return responses.LoginResponse{}, common.Wrap(fmt.Errorf("Login: !user.PasswordMatches"), customErrors.ErrWrongPassword)
	}

	tokenString, err := user.GenerateTokenString(s.auth)
	if err != nil {
		return responses.LoginResponse{}, common.Wrap(fmt.Errorf("Login: user.GenerateTokenString"), customErrors.ErrUnauthorized)
	}

	return responses.LoginResponse{Token: tokenString}, nil
}
