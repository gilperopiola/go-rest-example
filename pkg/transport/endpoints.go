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
//   - an empty request struct (just for type checking)
//   - a function that makes a request struct from the gin context
//   - a function that calls the service with that request
//
// It writes an HTTP response with the result of the service call.
// It also adds errors to the context so they can be retrieved by the middleware.

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

/*-------------------
//      AUTH
//-----------------*/

func (t transport) signup(c *gin.Context) {
	HandleRequest(c, &requests.SignupRequest{}, requests.MakeRequest[*requests.SignupRequest], t.Signup)
}

func (t transport) login(c *gin.Context) {
	HandleRequest(c, &requests.LoginRequest{}, requests.MakeRequest[*requests.LoginRequest], t.Login)
}

/*-------------------
//      USERS
//----------------*/

func (t transport) createUser(c *gin.Context) {
	HandleRequest(c, &requests.CreateUserRequest{}, requests.MakeRequest[*requests.CreateUserRequest], t.CreateUser)
}

func (t transport) getUser(c *gin.Context) {
	HandleRequest(c, &requests.GetUserRequest{}, requests.MakeRequest[*requests.GetUserRequest], t.GetUser)
}

func (t transport) updateUser(c *gin.Context) {
	HandleRequest(c, &requests.UpdateUserRequest{}, requests.MakeRequest[*requests.UpdateUserRequest], t.UpdateUser)
}

func (t transport) deleteUser(c *gin.Context) {
	HandleRequest(c, &requests.DeleteUserRequest{}, requests.MakeRequest[*requests.DeleteUserRequest], t.DeleteUser)
}

func (t transport) searchUsers(c *gin.Context) {
	HandleRequest(c, &requests.SearchUsersRequest{}, requests.MakeRequest[*requests.SearchUsersRequest], t.SearchUsers)
}

func (t transport) changePassword(c *gin.Context) {
	HandleRequest(c, &requests.ChangePasswordRequest{}, requests.MakeRequest[*requests.ChangePasswordRequest], t.ChangePassword)
}

/*--------------------
/       POSTS
//-----------------*/

func (t transport) createUserPost(c *gin.Context) {
	HandleRequest(c, &requests.CreateUserPostRequest{}, requests.MakeRequest[*requests.CreateUserPostRequest], t.CreateUserPost)
}

/*--------------------
//       MISC
//-----------------*/

func (t transport) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, common.HTTPResponse{
		Success: true,
		Content: "service is up and running :)",
	})
}
