package requests

import (
	"regexp"
	"strconv"

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

/*---------------
//  Pagination
/--------------*/

// PaginatedRequest defines the interface for requests that require pagination.
type PaginatedRequest interface {
	SetPage(page int)
	SetPerPage(perPage int)
}

// parseAndValidatePagination handles the parsing and validation of pagination parameters.
func parseAndValidatePagination(c common.GinI, req PaginatedRequest, defaultPage, defaultPerPage int) error {
	page, err := getQueryInt(c, "page", defaultPage)
	if err != nil {
		return common.ErrInvalidValue("page")
	}

	perPage, err := getQueryInt(c, "per_page", defaultPerPage)
	if err != nil {
		return common.ErrInvalidValue("per_page")
	}

	req.SetPage(page)
	req.SetPerPage(perPage)

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
	err := validate.Struct(request)
	if err == nil {
		return nil
	}

	// If we're here that means there have been errors
	if validationErrs, ok := err.(validator.ValidationErrors); ok { // TODO Fully fledge error messages
		firstErr := validationErrs[0]
		return common.Wrap(err.Error(), common.ErrInvalidValue(firstErr.StructField()))
	}

	return common.Wrap(err.Error(), common.ErrValidatingRequest)
}

func modelDeps(config *config.Config, repository models.RepositoryI) *models.ModelDependencies {
	return &models.ModelDependencies{
		Config:     config,
		Repository: repository,
	}
}

func getPtrStrValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// isEmail checks if the given string is an email.
func isEmail(str string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(str)
}

// getQueryInt parses an int value from query parameters with a default value.
func getQueryInt(c common.GinI, key string, defaultValue int) (int, error) {
	valueStr := c.DefaultQuery(key, strconv.Itoa(defaultValue))
	return strconv.Atoi(valueStr)
}
