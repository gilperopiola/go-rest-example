package transport

import (
	"errors"
	"net/http"

	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/utils"
)

type ErrorsMapper struct{}

type ErrorsMapperInterface interface {
	Map(err error) (status int, response HTTPResponse)
	MapWithType(errType, err error) (status int, response HTTPResponse)
}

func (e ErrorsMapper) MapWithType(errType, err error) (status int, response HTTPResponse) {
	return e.Map(utils.JoinErrors(errType, err))
}

func (e ErrorsMapper) Map(err error) (status int, response HTTPResponse) {

	// Generic errors

	if errors.Is(err, entities.ErrUnauthorized) {
		return returnErrorResponse(http.StatusUnauthorized, err)
	}

	if errors.Is(err, entities.ErrBindingRequest) {
		return returnErrorResponse(http.StatusBadRequest, err)
	}

	if errors.Is(err, entities.ErrAllFieldsRequired) {
		return returnErrorResponse(http.StatusBadRequest, err)
	}

	// Signup

	if errors.Is(err, entities.ErrPasswordsDontMatch) {
		return returnErrorResponse(http.StatusBadRequest, err)
	}

	if errors.Is(err, entities.ErrUsernameOrEmailAlreadyInUse) {
		return returnErrorResponse(http.StatusBadRequest, err)
	}

	// Default to internal server error
	return returnErrorResponse(http.StatusInternalServerError, err)
}
