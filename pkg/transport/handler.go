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
// It returns a response with the result of the service call.
func HandleRequest[req Request, resp Response](t TransportLayer, c *gin.Context,
	makeRequest func(requests.GinI) (req, error), serviceCall func(req) (resp, error)) {

	// Make, validate and get request
	request, err := makeRequest(c)
	if err != nil {
		c.JSON(t.ErrorsMapper().Map(err))
		return
	}

	// Call service with that request
	response, err := serviceCall(request)
	if err != nil {
		c.JSON(t.ErrorsMapper().Map(err))
		return
	}

	// Return OK
	c.JSON(http.StatusOK, common.HTTPResponse{
		Success: true,
		Content: response,
	})
}

type Request interface {
	requests.SignupRequest |
		requests.LoginRequest |
		requests.CreateUserRequest |
		requests.GetUserRequest |
		requests.UpdateUserRequest |
		requests.DeleteUserRequest |
		requests.SearchUsersRequest |
		requests.CreateUserPostRequest
}
type Response interface {
	responses.SignupResponse |
		responses.LoginResponse |
		responses.CreateUserResponse |
		responses.GetUserResponse |
		responses.UpdateUserResponse |
		responses.DeleteUserResponse |
		responses.SearchUsersResponse |
		responses.CreateUserPostResponse
}
