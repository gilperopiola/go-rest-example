package middleware

import (
	"errors"
	"net/http"

	"github.com/gilperopiola/go-rest-example/pkg/common"

	"github.com/gin-gonic/gin"
)

func NewErrorHandlerMiddleware(logger *LoggerAdapter) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Wait until the request is finished
		c.Next()

		// Then, check the context for errors
		if len(c.Errors) == 0 {
			return
		}

		// If there are errors, get the last one
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
	customErr, ok := lastErr.(*common.Error)
	if !ok {
		// Return 500 if not custom error
		return http.StatusInternalServerError, err.Error(), stackTrace
	}

	// Return status, custom error (second-to-last one) and stack trace
	return customErr.Status(), messages[len(messages)-1], stackTrace
}

func logStackTrace(logger *LoggerAdapter, status int, stackTrace, path, method string) {
	logContext := logger.WithField("status", status).WithField("path", path).WithField("method", method)
	logContext.Error(stackTrace)
}
