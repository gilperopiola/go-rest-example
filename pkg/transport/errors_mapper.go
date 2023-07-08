package transport

import (
	"errors"
	"net/http"

	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/utils"
)

type ErrorsMapper struct{}

type ErrorsMapperIface interface {
	Map(err error) (status int, response HTTPResponse)
	MapWithType(errType, err error) (status int, response HTTPResponse)
}

func (e ErrorsMapper) MapWithType(errType, err error) (status int, response HTTPResponse) {
	return e.Map(utils.JoinErrors(errType, err))
}

func (e ErrorsMapper) Map(err error) (status int, response HTTPResponse) {

	// Generic errors

	if errors.Is(err, entities.ErrUnauthorized) {
		return returnErrorResponseFunc(http.StatusUnauthorized, err)()
	}

	if errors.Is(err, entities.ErrBindingRequest) {
		return returnErrorResponseFunc(http.StatusBadRequest, err)()
	}

	if errors.Is(err, entities.ErrAllFieldsRequired) {
		return returnErrorResponseFunc(http.StatusBadRequest, err)()
	}

	// Signup

	if errors.Is(err, entities.ErrPasswordsDontMatch) {
		return returnErrorResponseFunc(http.StatusBadRequest, err)()
	}

	if errors.Is(err, entities.ErrUsernameOrEmailAlreadyInUse) {
		return returnErrorResponseFunc(http.StatusBadRequest, err)()
	}

	// Default to internal server error
	return returnErrorResponseFunc(http.StatusInternalServerError, err)()
}
