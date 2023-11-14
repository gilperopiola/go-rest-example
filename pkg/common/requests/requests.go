package requests

import (
	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/config"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Request interface {
	Build(c common.GinI) error
}

/*---------------------------------------------------------------------------
// All Requests must be here. All Requests must have a Build method.
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
	if err := validateRequest(validate, request); err != nil {
		return common.Wrap("validateRequest", err)
	}
	return nil
}

/*--------------
//   Helpers
/-------------*/

func bindRequestBody(c common.GinI, request Request) error {
	if err := c.ShouldBindJSON(&request); err != nil {
		return common.Wrap(err.Error(), common.ErrBindingRequest)
	}
	return nil
}

func validateRequest(validate *validator.Validate, request Request) error {
	if err := validate.Struct(request); err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok { // TODO Fully fledge error messages
			firstErr := validationErrs[0]
			return common.Wrap(err.Error(), common.ErrInvalidValue(firstErr.StructField()))
		}
		return common.Wrap(err.Error(), common.ErrValidatingRequest)
	}
	return nil
}

func modelDeps(config *config.Config, repository models.RepositoryI) *models.ModelDependencies {
	return &models.ModelDependencies{
		Config:     config,
		Repository: repository,
	}
}
