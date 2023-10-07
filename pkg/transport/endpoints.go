package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gin-gonic/gin"
)

/* Routes */

func (router *Router) SetPublicEndpoints(transport TransportLayer) {
	public := router.Group("/")
	{
		public.POST("/signup", transport.Signup)
		public.POST("/login", transport.Login)
	}
}

func (router *Router) SetUserEndpoints(transport TransportLayer, auth auth.AuthInterface) {
	users := router.Group("/users", auth.ValidateToken(entities.AnyRole))
	{
		users.GET("/:user_id", transport.GetUser)
		users.PATCH("/:user_id", transport.UpdateUser)
		users.DELETE("/:user_id", transport.DeleteUser)
	}
}

func (router *Router) SetAdminEndpoints(transport TransportLayer, auth auth.AuthInterface) {
	admin := router.Group("/admin", auth.ValidateToken(entities.AdminRole))
	{
		admin.POST("/user", transport.CreateUser)
	}
}

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
