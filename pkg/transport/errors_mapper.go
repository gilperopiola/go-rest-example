package transport

import (
	"net/http"
	"strings"

	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/utils"

	"github.com/sirupsen/logrus"
)

type ErrorsMapper struct {
	logger *logrus.Logger
}

func NewErrorsMapper(logger *logrus.Logger) ErrorsMapper {
	return ErrorsMapper{logger: logger}
}

type ErrorsMapperInterface interface {
	Map(err error) (status int, response HTTPResponse)
	MapWithType(errType, err error) (status int, response HTTPResponse)
}

func (e ErrorsMapper) MapWithType(errType, err error) (status int, response HTTPResponse) {
	return e.Map(utils.JoinErrors(errType, err))
}

// This method will define the response of the transport layer
func (e ErrorsMapper) Map(err error) (status int, response HTTPResponse) {

	// If we're here we shouldn't have a nil error
	if err == nil {
		err = entities.ErrNilError
	}

	// We check if the specific error msg is in the error chain
	// Assigning then the HTTP code and defaulting to 500
	responseStatusCode := http.StatusInternalServerError

	for key, value := range errorsMapToHTTPCode {
		if strings.Contains(err.Error(), key.Error()) {
			responseStatusCode = value
			break
		}
	}

	// We log 500's as errors, and 400's as warnings
	logWarningOrError(e.logger, err, responseStatusCode)

	return returnErrorResponse(responseStatusCode, err)
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
	entities.ErrNilError:     500,
	entities.ErrUnknown:      500,
}

func logWarningOrError(logger *logrus.Logger, err error, responseStatusCode int) {
	if responseStatusCode >= 500 {
		logger.Error(err.Error())
	} else {
		logger.Warn(err.Error())
	}
}
