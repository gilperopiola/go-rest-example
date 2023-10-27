package transport

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gilperopiola/go-rest-example/pkg/common"
)

// The errorsMapper maps errors to HTTP status codes
// It also logs errors and warnings

type errorsMapper struct {
	logger common.LoggerI
}

type errorsMapperI interface {
	Map(err error) (status int, response common.HTTPResponse)
}

func NewErrorsMapper(logger common.LoggerI) errorsMapper {
	return errorsMapper{logger: logger}
}

// This method will define the response of the transport layer
func (e errorsMapper) Map(err error) (statusCode int, response common.HTTPResponse) {

	// If we're here we shouldn't have a nil error
	if err == nil {
		err = common.ErrNilError
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
	common.ErrBindingRequest:        400,
	common.ErrAllFieldsRequired:     400,
	common.ErrPasswordsDontMatch:    400,
	common.ErrInvalidEmailFormat:    400,
	common.ErrInvalidUsernameLength: 400,
	common.ErrInvalidPasswordLength: 400,
	common.ErrInvalidValue:          400,

	// 401 - Unauthorized
	common.ErrUnauthorized:  401,
	common.ErrWrongPassword: 401,

	// 404 - Not Found
	common.ErrUserNotFound:       404,
	common.ErrUserAlreadyDeleted: 404,

	// 409 - Conflict
	common.ErrUsernameOrEmailAlreadyInUse: 409,

	// 500 - Internal Server Error
	common.ErrCreatingUser:     500,
	common.ErrGettingUser:      500,
	common.ErrUpdatingUser:     500,
	common.ErrDeletingUser:     500,
	common.ErrSearchingUsers:   500,
	common.ErrUnknown:          500,
	common.ErrNilError:         500,
	common.ErrCreatingUserPost: 500,
}

func (e errorsMapper) logWarningOrError(err error, responseStatusCode int) {
	chain := err.Error()

	logFn := e.logger.Info
	if responseStatusCode >= 500 {
		logFn = e.logger.Error
	}

	logFn(fmt.Sprint(responseStatusCode) + ": " + chain)
}
