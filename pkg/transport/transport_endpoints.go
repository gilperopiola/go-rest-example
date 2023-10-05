package transport

import (
	"github.com/gin-gonic/gin"
)

/* Auth */

func (t Transport) Signup(c *gin.Context) {
	HandleRequest(t, c, makeSignupRequest, t.Service.Signup)
}

func (t Transport) Login(c *gin.Context) {
	HandleRequest(t, c, makeLoginRequest, t.Service.Login)
}

/* Users */

func (t Transport) CreateUser(c *gin.Context) {
	HandleRequest(t, c, makeCreateUserRequest, t.Service.CreateUser)
}

func (t Transport) GetUser(c *gin.Context) {
	HandleRequest(t, c, makeGetUserRequest, t.Service.GetUser)
}

func (t Transport) UpdateUser(c *gin.Context) {
	HandleRequest(t, c, makeUpdateUserRequest, t.Service.UpdateUser)
}

func (t Transport) DeleteUser(c *gin.Context) {
	HandleRequest(t, c, makeDeleteUserRequest, t.Service.DeleteUser)
}
