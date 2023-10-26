package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/common/requests"
	"github.com/gin-gonic/gin"
)

//-------------------
//       AUTH
//-------------------

func (t transport) signup(c *gin.Context) {
	HandleRequest(t, c, requests.MakeSignupRequest, t.Service().Signup)
}

func (t transport) login(c *gin.Context) {
	HandleRequest(t, c, requests.MakeLoginRequest, t.Service().Login)
}

//-------------------
//      USERS
//-------------------

func (t transport) createUser(c *gin.Context) {
	HandleRequest(t, c, requests.MakeCreateUserRequest, t.Service().CreateUser)
}

func (t transport) getUser(c *gin.Context) {
	HandleRequest(t, c, requests.MakeGetUserRequest, t.Service().GetUser)
}

func (t transport) updateUser(c *gin.Context) {
	HandleRequest(t, c, requests.MakeUpdateUserRequest, t.Service().UpdateUser)
}

func (t transport) deleteUser(c *gin.Context) {
	HandleRequest(t, c, requests.MakeDeleteUserRequest, t.Service().DeleteUser)
}

func (t transport) searchUsers(c *gin.Context) {
	HandleRequest(t, c, requests.MakeSearchUsersRequest, t.Service().SearchUsers)
}

//-------------------
//      POSTS
//-------------------

func (t transport) createUserPost(c *gin.Context) {
	HandleRequest(t, c, requests.MakeCreateUserPostRequest, t.Service().CreateUserPost)
}

//-------------------
//      MISC
//-------------------

func (t transport) healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"status": "API is up and running :)"})
}
