package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/requests"
	"github.com/gilperopiola/go-rest-example/pkg/common/responses"
	"github.com/gilperopiola/go-rest-example/pkg/repository/options"
)

//-----------------------
//       SIGNUP
//-----------------------

func (s *service) Signup(request *requests.SignupRequest) (responses.SignupResponse, error) {
	user := request.ToUserModel()

	if user.Exists(s.repository) {
		return responses.SignupResponse{}, common.Wrap("Signup: user.Exists", common.ErrUsernameOrEmailAlreadyInUse)
	}

	user.HashPassword(s.config.JWT.HashSalt)

	if err := user.Create(s.repository); err != nil {
		return responses.SignupResponse{}, common.Wrap("Signup: user.Create", err)
	}

	return responses.SignupResponse{User: user.ToResponseModel()}, nil
}

//---------------------
//       LOGIN
//---------------------

func (s *service) Login(request *requests.LoginRequest) (responses.LoginResponse, error) {
	user := request.ToUserModel()

	if err := user.Get(s.repository, options.WithoutDeleted); err != nil {
		return responses.LoginResponse{}, common.Wrap("Login: user.Get", err)
	}

	if !user.PasswordMatches(request.Password, s.config.JWT.HashSalt) {
		return responses.LoginResponse{}, common.Wrap("Login: !user.PasswordMatches", common.ErrWrongPassword)
	}

	tokenString, err := user.GenerateTokenString(s.auth)
	if err != nil {
		return responses.LoginResponse{}, common.Wrap("Login: user.GenerateTokenString", common.ErrUnauthorized)
	}

	return responses.LoginResponse{Token: tokenString}, nil
}
