package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/common/requests"

	"github.com/gin-gonic/gin"
)

//-------------------
//       AUTH
//-------------------

func (t transport) signup(c *gin.Context) {
	HandleRequest(c, requests.MakeSignupRequest, t.Service().Signup, t.ErrorsMapper())
}

func (t transport) login(c *gin.Context) {
	HandleRequest(c, requests.MakeLoginRequest, t.Service().Login, t.ErrorsMapper())
}

//------------------
//      USERS
//------------------

func (t transport) createUser(c *gin.Context) {
	HandleRequest(c, requests.MakeCreateUserRequest, t.Service().CreateUser, t.ErrorsMapper())
}

func (t transport) getUser(c *gin.Context) {
	HandleRequest(c, requests.MakeGetUserRequest, t.Service().GetUser, t.ErrorsMapper())
}

func (t transport) updateUser(c *gin.Context) {
	HandleRequest(c, requests.MakeUpdateUserRequest, t.Service().UpdateUser, t.ErrorsMapper())
}

func (t transport) deleteUser(c *gin.Context) {
	HandleRequest(c, requests.MakeDeleteUserRequest, t.Service().DeleteUser, t.ErrorsMapper())
}

func (t transport) searchUsers(c *gin.Context) {
	HandleRequest(c, requests.MakeSearchUsersRequest, t.Service().SearchUsers, t.ErrorsMapper())
}

func (t transport) changePassword(c *gin.Context) {
	HandleRequest(c, requests.MakeChangePasswordRequest, t.Service().ChangePassword, t.ErrorsMapper())
}

//-------------------
//      POSTS
//-------------------

func (t transport) createUserPost(c *gin.Context) {
	HandleRequest(c, requests.MakeCreateUserPostRequest, t.Service().CreateUserPost, t.ErrorsMapper())
}
