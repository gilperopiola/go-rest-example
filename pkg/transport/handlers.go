package transport

import (
	"net/http"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/requests"
	"github.com/gilperopiola/go-rest-example/pkg/common/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

/*----------------------------------------------------------------------------------------
//           THIS IS THE MOST IMPORTANT FUNCTION ON THE WHOLE PROJECT
//
// After a route is called and a handler is executed, this function is called.
// It unifies all endpoints under a simple logic:
//  - 1. Build & Validate the Request.
//  - 2. Call the Service, get the returning Response.
//  - 3. Add errors to Context. If there are none, write back the OK HTTP Response to the client.
//
// handleRequest takes:
//   - 1. A transport and a gin context
//   - 2. An empty request struct (the concrete type)
//   - 3. A go-playground validator (https://github.com/go-playground/validator/)
//   - 4. A function that makes a XXXRequest struct from the gin context
//   - 5. A function that calls the Service with that XXXRequest and returns a XXXResponse
//----------------------------------------------------------------------------------*/

func handleRequest[reqT requests.All, respT responses.All](c *gin.Context, emptyReq reqT, val *validator.Validate,
	makeRequestFn func(*gin.Context, reqT, *validator.Validate) (reqT, error), serviceFn func(*gin.Context, reqT) (respT, error)) {

	// Build, validate and get request
	request, err := makeRequestFn(c, emptyReq, val)
	if err != nil {
		c.Error(err)
		return // Transport error
	}

	// Call service with that request
	response, err := serviceFn(c, request)
	if err != nil {
		c.Error(err)
		return // Service & Repository errors
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
	handleRequest(c, &requests.SignupRequest{}, t.validate, requests.MakeRequest[*requests.SignupRequest], t.Signup)
}

func (t transport) login(c *gin.Context) {
	handleRequest(c, &requests.LoginRequest{}, t.validate, requests.MakeRequest[*requests.LoginRequest], t.Login)
}

func (t transport) createUser(c *gin.Context) {
	handleRequest(c, &requests.CreateUserRequest{}, t.validate, requests.MakeRequest[*requests.CreateUserRequest], t.CreateUser)
}

func (t transport) getUser(c *gin.Context) {
	handleRequest(c, &requests.GetUserRequest{}, t.validate, requests.MakeRequest[*requests.GetUserRequest], t.GetUser)
}

func (t transport) updateUser(c *gin.Context) {
	handleRequest(c, &requests.UpdateUserRequest{}, t.validate, requests.MakeRequest[*requests.UpdateUserRequest], t.UpdateUser)
}

func (t transport) deleteUser(c *gin.Context) {
	handleRequest(c, &requests.DeleteUserRequest{}, t.validate, requests.MakeRequest[*requests.DeleteUserRequest], t.DeleteUser)
}

func (t transport) searchUsers(c *gin.Context) {
	handleRequest(c, &requests.SearchUsersRequest{}, t.validate, requests.MakeRequest[*requests.SearchUsersRequest], t.SearchUsers)
}

func (t transport) changePassword(c *gin.Context) {
	handleRequest(c, &requests.ChangePasswordRequest{}, t.validate, requests.MakeRequest[*requests.ChangePasswordRequest], t.ChangePassword)
}

func (t transport) createUserPost(c *gin.Context) {
	handleRequest(c, &requests.CreateUserPostRequest{}, t.validate, requests.MakeRequest[*requests.CreateUserPostRequest], t.CreateUserPost)
}

/*--------------------
//       Misc
//------------------*/

func (t transport) healthCheck(c *gin.Context) {

	success, content := true, "service is up and running :)"

	if err := t.sqlDB.Ping(); err != nil {
		success, content = false, "error pinging database :("
	}

	c.JSON(http.StatusOK, common.HTTPResponse{Success: success, Content: content})
}
