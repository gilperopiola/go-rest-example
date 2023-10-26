package requests

import (
	"github.com/gilperopiola/go-rest-example/pkg/common"
)

type SignupRequest struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`

	// User Detail
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func MakeSignupRequest(c GinI) (request SignupRequest, err error) {
	if err := request.Build(c); err != nil {
		return SignupRequest{}, common.Wrap("makeSignupRequest", err)
	}
	if err = request.Validate(); err != nil {
		return SignupRequest{}, common.Wrap("makeSignupRequest", err)
	}
	return request, nil
}

type LoginRequest struct {
	UsernameOrEmail string `json:"username_or_email"`
	Password        string `json:"password"`
}

func MakeLoginRequest(c GinI) (request LoginRequest, err error) {
	if err := request.Build(c); err != nil {
		return LoginRequest{}, common.Wrap("makeLoginRequest", err)
	}
	if err = request.Validate(); err != nil {
		return LoginRequest{}, common.Wrap("makeLoginRequest", err)
	}
	return request, nil
}
