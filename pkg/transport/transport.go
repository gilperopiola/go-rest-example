package transport

import (
	"net/http"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/requests"
	"github.com/gilperopiola/go-rest-example/pkg/common/responses"
	"github.com/gilperopiola/go-rest-example/pkg/service"

	"github.com/gin-gonic/gin"
)

type Transport struct {
	Service      service.ServiceLayer
	ErrorsMapper errorsMapperI
}

type TransportLayer interface {

	// - Auth
	Signup(c *gin.Context)
	Login(c *gin.Context)

	// - Users
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

func NewTransport(service service.ServiceLayer, errorsMapper errorsMapperI) Transport {
	return Transport{
		Service:      service,
		ErrorsMapper: errorsMapper,
	}
}

// HandleRequest takes:
//
//   - a transport and a gin context
//   - a function that makes a request from the gin context
//   - a function that calls the service with that request
//
// It returns a response with the result of the service call.
func HandleRequest[req Request, resp Response](t Transport, c *gin.Context,
	makeRequest func(requests.GinI) (req, error), serviceCall func(req) (resp, error)) {

	// Make, validate and get request
	request, err := makeRequest(c)
	if err != nil {
		c.JSON(t.ErrorsMapper.Map(err))
		return
	}

	// Call service with that request
	response, err := serviceCall(request)
	if err != nil {
		c.JSON(t.ErrorsMapper.Map(err))
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
		requests.DeleteUserRequest
}
type Response interface {
	responses.SignupResponse |
		responses.LoginResponse |
		responses.CreateUserResponse |
		responses.GetUserResponse |
		responses.UpdateUserResponse |
		responses.DeleteUserResponse
}
