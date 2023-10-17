package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/common/requests"
	"github.com/gin-gonic/gin"
)

/* Routes */

func (router *Router) SetPublicEndpoints(transport TransportLayer) {
	public := router.Group("/")
	{
		public.POST("/health", healthCheck)

		public.POST("/signup", transport.Signup)
		public.POST("/login", transport.Login)
	}
}

func (router *Router) SetUserEndpoints(transport TransportLayer, authI auth.AuthI) {
	users := router.Group("/users", authI.ValidateToken(auth.AnyRole, true))
	{
		users.GET("/:user_id", transport.GetUser)
		users.PATCH("/:user_id", transport.UpdateUser)
		users.DELETE("/:user_id", transport.DeleteUser)
	}

	posts := users.Group("/:user_id/posts")
	{
		posts.POST("", transport.CreateUserPost)
	}
}

func (router *Router) SetAdminEndpoints(transport TransportLayer, authI auth.AuthI) {
	admin := router.Group("/admin", authI.ValidateToken(auth.AdminRole, false))
	{
		admin.POST("/user", transport.CreateUser)
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"status": "API is up and running :)"})
}

/* Auth */

func (t Transport) Signup(c *gin.Context) {
	HandleRequest(t, c, requests.MakeSignupRequest, t.Service.Signup)
}

func (t Transport) Login(c *gin.Context) {
	HandleRequest(t, c, requests.MakeLoginRequest, t.Service.Login)
}

/* Users */

func (t Transport) CreateUser(c *gin.Context) {
	HandleRequest(t, c, requests.MakeCreateUserRequest, t.Service.CreateUser)
}

func (t Transport) GetUser(c *gin.Context) {
	HandleRequest(t, c, requests.MakeGetUserRequest, t.Service.GetUser)
}

func (t Transport) UpdateUser(c *gin.Context) {
	HandleRequest(t, c, requests.MakeUpdateUserRequest, t.Service.UpdateUser)
}

func (t Transport) DeleteUser(c *gin.Context) {
	HandleRequest(t, c, requests.MakeDeleteUserRequest, t.Service.DeleteUser)
}

func (t Transport) CreateUserPost(c *gin.Context) {
	HandleRequest(t, c, requests.MakeCreateUserPostRequest, t.Service.CreateUserPost)
}
