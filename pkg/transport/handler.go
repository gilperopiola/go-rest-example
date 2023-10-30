package transport

import (
	"net/http"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/requests"
	"github.com/gilperopiola/go-rest-example/pkg/common/responses"

	"github.com/gin-gonic/gin"
)

// HandleRequest takes:
//
//   - a transport and a gin context
//   - a function that makes a request from the gin context
//   - a function that calls the service with that request
//
// It writes an HTTP response with the result of the service call.

func HandleRequest[req requests.All, resp responses.All](c *gin.Context, emptyReq req, makeRequestFn func(requests.GinI, req) (req, error), serviceCallFn func(req) (resp, error)) {

	// Build, validate and get request
	request, err := makeRequestFn(c, emptyReq)
	if err != nil {
		c.Error(err)
		return
	}

	// Call service with that request
	response, err := serviceCallFn(request)
	if err != nil {
		c.Error(err)
		return
	}

	// Return OK
	c.JSON(http.StatusOK, common.HTTPResponse{
		Success: true,
		Content: response,
	})
}
