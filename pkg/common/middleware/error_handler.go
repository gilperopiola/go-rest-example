package middleware

import (
	"errors"
	"net/http"

	"github.com/gilperopiola/go-rest-example/pkg/common"

	"github.com/gin-gonic/gin"
)

/*--------------------------------------------------------------------------------------------------------
// This Error Handler is an important part of the system. It is one of the last middlewares to be called,
// and it is responsible for logging the errors and returning a human-readable error to the client.
//------------------------------------------------------------*/

func NewErrorHandlerMiddleware(logger *LoggerAdapter) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Wait until the request is finished
		c.Next()

		// Then, check the context for errors
		if len(c.Errors) == 0 {
			return
		}

		// Get the last error
		err := c.Errors.Last()

		// Get all the info we need
		url := c.Request.URL.Path
		method := c.Request.Method
		statusCode, humanReadable, stackTrace := getErrorInfo(err)

		// Log the error
		logStackTrace(logger, statusCode, stackTrace, url, method)

		c.JSON(statusCode, common.HTTPResponse{
			Success: false,
			Content: nil,
			Error:   humanReadable,
		})
	}
}

func logStackTrace(logger *LoggerAdapter, httpStatus int, stackTrace, path, method string) {
	logger.WithFields(map[string]interface{}{
		"status": httpStatus,
		"path":   path,
		"method": method,
	}).Error(stackTrace)
}

// getErrorInfo returns the status, the human-readable string & the stack trace of the error
func getErrorInfo(err error) (int, string, string) {
	var (
		stackTrace      = err.Error()
		messages        []string
		secondToLastErr error
	)

	// Keep unwrapping the error until you reach the last one. Append every message to the messages slice. Save the second-to-last error.
	for err != nil {
		messages = append(messages, err.Error())
		secondToLastErr = err
		err = errors.Unwrap(err)
	}

	customErr, ok := secondToLastErr.(*common.Error)
	if !ok {
		return http.StatusInternalServerError, err.Error(), stackTrace
	}

	// The second-to-last error on a chain is the one that contains the human-readable string
	return customErr.Status(), messages[len(messages)-1], stackTrace
}
