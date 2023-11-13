package requests

import (
	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"
)

type Request interface {
	Build(c common.GinI) error
	Validate(validate *validator.Validate) error
}

/*---------------------------------------------------------------------------
// All Requests must be here. All Requests must have Build and Validate methods.
------------------------*/

type All interface {
	Request

	*SignupRequest |
		*LoginRequest |
		*CreateUserRequest |
		*GetUserRequest |
		*UpdateUserRequest |
		*DeleteUserRequest |
		*SearchUsersRequest |
		*ChangePasswordRequest |
		*CreateUserPostRequest
}

/*---------------------------------------------------------------------------
// The MakeRequest function is very important, it uses generics to orchestrate
// the generation and validation of the different types of requests.
------------------------*/

func MakeRequest[req All](c *gin.Context, request req, validate *validator.Validate) (req, error) {
	if err := makeRequest(c, request, validate); err != nil {
		return req(nil), err
	}
	return request, nil
}

func makeRequest[req All](c *gin.Context, request req, validate *validator.Validate) error {
	if err := request.Build(c); err != nil {
		return common.Wrap("request.Build", err)
	}
	if err := request.Validate(validate); err != nil {
		return common.Wrap("request.Validate", err)
	}
	return nil
}
