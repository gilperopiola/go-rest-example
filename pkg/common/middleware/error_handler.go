package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gilperopiola/go-rest-example/pkg/common"

	"github.com/gin-gonic/gin"
)

func NewErrorHandlerMiddleware(logger common.LoggerI) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Wait until the request is finished
		c.Next()

		// Check the context for errors
		if len(c.Errors) == 0 {
			return
		}

		// If there are, get the last one
		err := c.Errors.Last()

		statusCode, stackTrace := getStatusCodeFromError(err), err.Error()

		// Log the error depending on severity
		logStackTrace(logger, statusCode, stackTrace, c.FullPath())

		c.JSON(statusCode, common.HTTPResponse{
			Success: false,
			Content: nil,
			Error:   getHumanReadableError(err),
		})
	}
}

func logStackTrace(logger common.LoggerI, status int, stackTrace, path string) {
	logContext := logger.WithField("status", status).WithField("path", path)
	if status >= http.StatusInternalServerError {
		logContext.Error(stackTrace)
	} else {
		logContext.Info(stackTrace)
	}
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
	common.ErrCreatingUserPost: 500,

	// Other
	common.ErrTooManyRequests: 429,
}

// getHumanReadableError returns the first error in the chain of errors
func getHumanReadableError(err error) string {
	var messages []string
	for err != nil {
		messages = append(messages, err.Error())
		err = errors.Unwrap(err)
	}
	return messages[len(messages)-1]
}

// getStatusCodeFromError returns the HTTP status code that corresponds to the error
func getStatusCodeFromError(err error) int {
	statusCode := http.StatusInternalServerError

	// This is done through strings comparison!!! (not ideal)
	for key, value := range errorsMapToHTTPCode {
		if strings.Contains(err.Error(), key.Error()) {
			statusCode = value
			break
		}
	}
	return statusCode
}
