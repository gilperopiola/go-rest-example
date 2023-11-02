package middleware

import (
	"errors"
	"net/http"

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

		statusCode, humanReadable, stackTrace := getErrorInfo(err)
		method := c.Request.Method

		// Log the error depending on severity
		logStackTrace(logger, statusCode, stackTrace, c.Request.URL.Path, method)

		c.JSON(statusCode, common.HTTPResponse{
			Success: false,
			Content: nil,
			Error:   humanReadable,
		})
	}
}

// getErrorInfo returns the status, the human-readable string & the stack trace of the error
func getErrorInfo(err error) (int, string, string) {
	var (
		stackTrace = err.Error()
		messages   []string
		lastErr    error
	)

	// Unwrap the error and get all the messages
	for err != nil {
		messages = append(messages, err.Error())
		lastErr = err
		err = errors.Unwrap(err)
	}

	// Assert the type of the second-to-last error
	if customErr, ok := lastErr.(*common.Error); ok {
		return customErr.Status(), messages[len(messages)-1], stackTrace
	}

	// Not a custom error, send 500
	return http.StatusInternalServerError, err.Error(), stackTrace
}

func logStackTrace(logger common.LoggerI, status int, stackTrace, path, method string) {
	logContext := logger.WithField("status", status).WithField("path", path).WithField("method", method)
	if status >= http.StatusInternalServerError {
		logContext.Error(stackTrace)
	} else {
		logContext.Info(stackTrace)
	}
}
