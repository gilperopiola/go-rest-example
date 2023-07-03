package transport

import (
	"errors"
	"net/http"

	"github.com/gilperopiola/go-rest-example/pkg/entities"
)

type ErrorsMapperer interface {
	Map(err error) (status int, response HTTPResponse)
}

type ErrorsMapper struct {
}

func (e ErrorsMapper) Map(err error) (status int, response HTTPResponse) {

	// Signup
	if errors.Is(err, entities.ErrAllFieldsRequired) {
		return returnErrorResponseFunc(http.StatusBadRequest, err)()
	}
	if errors.Is(err, entities.ErrPasswordsDontMatch) {
		return returnErrorResponseFunc(http.StatusBadRequest, err)()
	}
	if errors.Is(err, entities.ErrUsernameOrEmailAlreadyInUse) {
		return returnErrorResponseFunc(http.StatusBadRequest, err)()
	}

	// Default to internal server error
	return returnErrorResponseFunc(http.StatusInternalServerError, err)()
}
