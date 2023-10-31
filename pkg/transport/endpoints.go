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
//   - an empty request struct
//   - a function that makes a request struct from the gin context
//   - a function that calls the service with that request
//
// It writes an HTTP response with the result of the service call.

func HandleRequest[req requests.All, resp responses.All](c *gin.Context, emptyReq req, makeRequestFn func(common.GinI, req) (req, error), serviceCallFn func(req) (resp, error)) {

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

//-------------------
//       AUTH
//-------------------

func (t transport) signup(c *gin.Context) {
	HandleRequest(c, &requests.SignupRequest{}, requests.MakeRequest[*requests.SignupRequest], t.Service().Signup)
}

func (t transport) login(c *gin.Context) {
	HandleRequest(c, &requests.LoginRequest{}, requests.MakeRequest[*requests.LoginRequest], t.Service().Login)
}

//------------------
//      USERS
//------------------

func (t transport) createUser(c *gin.Context) {
	HandleRequest(c, &requests.CreateUserRequest{}, requests.MakeRequest[*requests.CreateUserRequest], t.Service().CreateUser)
}

func (t transport) getUser(c *gin.Context) {
	HandleRequest(c, &requests.GetUserRequest{}, requests.MakeRequest[*requests.GetUserRequest], t.Service().GetUser)
}

func (t transport) updateUser(c *gin.Context) {
	HandleRequest(c, &requests.UpdateUserRequest{}, requests.MakeRequest[*requests.UpdateUserRequest], t.Service().UpdateUser)
}

func (t transport) deleteUser(c *gin.Context) {
	HandleRequest(c, &requests.DeleteUserRequest{}, requests.MakeRequest[*requests.DeleteUserRequest], t.Service().DeleteUser)
}

func (t transport) searchUsers(c *gin.Context) {
	HandleRequest(c, &requests.SearchUsersRequest{}, requests.MakeRequest[*requests.SearchUsersRequest], t.Service().SearchUsers)
}

func (t transport) changePassword(c *gin.Context) {
	HandleRequest(c, &requests.ChangePasswordRequest{}, requests.MakeRequest[*requests.ChangePasswordRequest], t.Service().ChangePassword)
}

//-------------------
//      POSTS
//-------------------

func (t transport) createUserPost(c *gin.Context) {
	HandleRequest(c, &requests.CreateUserPostRequest{}, requests.MakeRequest[*requests.CreateUserPostRequest], t.Service().CreateUserPost)
}

//-------------------
//       MISC
//-------------------

func (t transport) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, common.HTTPResponse{
		Success: true,
		Content: "service is up and running :)",
	})
}
