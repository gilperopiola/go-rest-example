package transport

import (
	"net/http"

	"github.com/gilperopiola/go-rest-example/pkg/common"
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
	makeRequest func(*gin.Context) (req, error), serviceCall func(req) (resp, error)) {

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
	common.SignupRequest |
		common.LoginRequest |
		common.CreateUserRequest |
		common.GetUserRequest |
		common.UpdateUserRequest |
		common.DeleteUserRequest
}
type Response interface {
	common.SignupResponse |
		common.LoginResponse |
		common.CreateUserResponse |
		common.GetUserResponse |
		common.UpdateUserResponse |
		common.DeleteUserResponse
}
