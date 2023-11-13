package transport

import (
	"net/http"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/requests"
	"github.com/gilperopiola/go-rest-example/pkg/common/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

/*----------------------
// HandleRequest takes:
//
//   - a transport and a gin context
//   - an empty request struct (just for type checking)
//   - a validator
//   - a function that makes a request struct from the gin context
//   - a function that calls the service with that request
//
// It writes an HTTP response with the result of the service call.
// It also adds errors to the context so they can be retrieved by the error handler.
//----------------------------------------------------------------------------------*/

func HandleRequest[reqT requests.All, respT responses.All](c *gin.Context, emptyReq reqT, val *validator.Validate,
	makeRequestFn func(*gin.Context, reqT, *validator.Validate) (reqT, error), serviceFn func(*gin.Context, reqT) (respT, error)) {

	// Build, validate and get request
	request, err := makeRequestFn(c, emptyReq, val)
	if err != nil {
		c.Error(err) // Transport error
		return
	}

	// Call service with that request
	response, err := serviceFn(c, request)
	if err != nil {
		c.Error(err) // Service / Repository error
		return
	}

	// Return OK
	c.JSON(http.StatusOK, common.HTTPResponse{
		Success: true,
		Content: response,
	})
}

/*-------------------
//     Handlers
//----------------*/

func (t transport) signup(c *gin.Context) {
	HandleRequest(c, &requests.SignupRequest{}, t.validate, requests.MakeRequest[*requests.SignupRequest], t.Signup)
}

func (t transport) login(c *gin.Context) {
	HandleRequest(c, &requests.LoginRequest{}, t.validate, requests.MakeRequest[*requests.LoginRequest], t.Login)
}

func (t transport) createUser(c *gin.Context) {
	HandleRequest(c, &requests.CreateUserRequest{}, t.validate, requests.MakeRequest[*requests.CreateUserRequest], t.CreateUser)
}

func (t transport) getUser(c *gin.Context) {
	HandleRequest(c, &requests.GetUserRequest{}, t.validate, requests.MakeRequest[*requests.GetUserRequest], t.GetUser)
}

func (t transport) updateUser(c *gin.Context) {
	HandleRequest(c, &requests.UpdateUserRequest{}, t.validate, requests.MakeRequest[*requests.UpdateUserRequest], t.UpdateUser)
}

func (t transport) deleteUser(c *gin.Context) {
	HandleRequest(c, &requests.DeleteUserRequest{}, t.validate, requests.MakeRequest[*requests.DeleteUserRequest], t.DeleteUser)
}

func (t transport) searchUsers(c *gin.Context) {
	HandleRequest(c, &requests.SearchUsersRequest{}, t.validate, requests.MakeRequest[*requests.SearchUsersRequest], t.SearchUsers)
}

func (t transport) changePassword(c *gin.Context) {
	HandleRequest(c, &requests.ChangePasswordRequest{}, t.validate, requests.MakeRequest[*requests.ChangePasswordRequest], t.ChangePassword)
}

func (t transport) createUserPost(c *gin.Context) {
	HandleRequest(c, &requests.CreateUserPostRequest{}, t.validate, requests.MakeRequest[*requests.CreateUserPostRequest], t.CreateUserPost)
}

/*--------------------
//       Misc
//------------------*/

func (t transport) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, common.HTTPResponse{
		Success: true,
		Content: "service is up and running :)",
	})
}
