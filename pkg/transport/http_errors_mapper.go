package transport

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	customErrors "github.com/gilperopiola/go-rest-example/pkg/common/errors"
	"github.com/gilperopiola/go-rest-example/pkg/common/middleware"
)

// The errorsMapper maps errors to HTTP status codes
// It also logs errors and warnings

type errorsMapper struct {
	logger middleware.LoggerI
}

type errorsMapperI interface {
	Map(err error) (status int, response common.HTTPResponse)
}

func NewErrorsMapper(logger middleware.LoggerI) errorsMapper {
	return errorsMapper{logger: logger}
}

// This method will define the response of the transport layer
func (e errorsMapper) Map(err error) (statusCode int, response common.HTTPResponse) {

	// If we're here we shouldn't have a nil error
	if err == nil {
		err = customErrors.ErrNilError
	}

	// We get the HTTP code depending on the error, defaulting to 500
	statusCode = getStatusCodeFromError(err)

	// We log 500's as errors, and 400's as warnings
	e.logWarningOrError(err, statusCode)

	return statusCode, common.HTTPResponse{
		Success: false,
		Content: nil,
		Error:   getFirstError(err),
	}
}

func getFirstError(err error) string {
	var messages []string
	for err != nil {
		messages = append(messages, err.Error())
		err = errors.Unwrap(err)
	}
	return messages[len(messages)-1]
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
	customErrors.ErrBindingRequest:        400,
	customErrors.ErrAllFieldsRequired:     400,
	customErrors.ErrPasswordsDontMatch:    400,
	customErrors.ErrInvalidEmailFormat:    400,
	customErrors.ErrInvalidUsernameLength: 400,
	customErrors.ErrInvalidPasswordLength: 400,
	customErrors.ErrInvalidValue:          400,

	// 401 - Unauthorized
	customErrors.ErrUnauthorized:  401,
	customErrors.ErrWrongPassword: 401,

	// 404 - Not Found
	customErrors.ErrUserNotFound:       404,
	customErrors.ErrUserAlreadyDeleted: 404,

	// 409 - Conflict
	customErrors.ErrUsernameOrEmailAlreadyInUse: 409,

	// 500 - Internal Server Error
	customErrors.ErrCreatingUser:     500,
	customErrors.ErrGettingUser:      500,
	customErrors.ErrUpdatingUser:     500,
	customErrors.ErrDeletingUser:     500,
	customErrors.ErrSearchingUsers:   500,
	customErrors.ErrUnknown:          500,
	customErrors.ErrNilError:         500,
	customErrors.ErrCreatingUserPost: 500,
}

func (e errorsMapper) logWarningOrError(err error, responseStatusCode int) {
	chain := err.Error()

	logFn := e.logger.Info
	if responseStatusCode >= 500 {
		logFn = e.logger.Error
	}

	logFn(fmt.Sprint(responseStatusCode) + ": " + chain)
}
