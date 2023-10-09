package transport

import (
	"net/http"
	"strings"

	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/logger"
)

// The errorsMapper maps errors to HTTP status codes
// It also logs errors and warnings

type errorsMapper struct {
	logger logger.LoggerI
}

type errorsMapperI interface {
	Map(err error) (status int, response HTTPResponse)
}

func NewErrorsMapper(logger logger.LoggerI) errorsMapper {
	return errorsMapper{logger: logger}
}

// This method will define the response of the transport layer
func (e errorsMapper) Map(err error) (statusCode int, response HTTPResponse) {

	// If we're here we shouldn't have a nil error
	if err == nil {
		err = entities.ErrNilError
	}

	// We get the HTTP code depending on the error, defaulting to 500
	statusCode = getStatusCodeFromError(err)

	// We log 500's as errors, and 400's as warnings
	e.logWarningOrError(err, statusCode)

	return statusCode, HTTPResponse{
		Success: false,
		Content: nil,
		Error:   err.Error(),
	}
}

func getStatusCodeFromError(err error) int {
	responseStatusCode := http.StatusInternalServerError

	// This is done through strings comparison!!! (not ideal)
	for key, value := range errorsMapToHTTPCode {
		if strings.Contains(err.Error(), key.Error()) {
			responseStatusCode = value
			break
		}
	}

	return responseStatusCode
}

var errorsMapToHTTPCode = map[error]int{
	// 400 - Bad Request
	entities.ErrBindingRequest:        400,
	entities.ErrAllFieldsRequired:     400,
	entities.ErrPasswordsDontMatch:    400,
	entities.ErrInvalidEmailFormat:    400,
	entities.ErrInvalidUsernameLength: 400,
	entities.ErrInvalidPasswordLength: 400,

	// 401 - Unauthorized
	entities.ErrUnauthorized:  401,
	entities.ErrWrongPassword: 401,

	// 404 - Not Found
	entities.ErrUserNotFound: 404,

	// 409 - Conflict
	entities.ErrUsernameOrEmailAlreadyInUse: 409,

	// 500 - Internal Server Error
	entities.ErrCreatingUser: 500,
	entities.ErrGettingUser:  500,
	entities.ErrNilError:     500,
	entities.ErrUnknown:      500,
}

func (e errorsMapper) logWarningOrError(err error, responseStatusCode int) {
	logFn := e.logger.Warn
	if responseStatusCode >= 500 {
		logFn = e.logger.Error
	}
	logFn(err.Error())
}
