package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/common/requests"

	"github.com/gin-gonic/gin"
)

/*
	This is the entrypoint of the HTTP requests.
		1. It matches the URL
		2. It calls ValidateToken if necessary
			a. If it's valid, it sets the user info in the context
		3. The corresponding function below in this file is called
		4. The HandleRequest method in transport.go is called
			a. It calls requests_maker.go
				1. It makes the request from the gin context
				2. It validates the request
			b. Now that we have our custom Request, we use it to call the service
				1. The corresponding method in service_xxx.go is called
				2. Our custom Request is converted to a Model
				3. A custom Handler is created from that Model, allowing us to interact with it
				4. The Handler is called, which in turn calls the repository_xxx.go file
					a. The repository_xxx.go file interacts with the database and returns a Model
				5. The Handler returns the obtained Model
				6. The Model is converted to a Response	Model
				7. The Response is returned to the HandleRequest method in transport.go
			c. The HandleRequest method in transport.go returns the HTTP Request
			d. Errors are also handled here, on HandleRequest
*/

//-----------------------------
//      ROUTES / ENDPOINTS
//-----------------------------

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

//-------------------
//       AUTH
//-------------------

func (t Transport) Signup(c *gin.Context) {
	HandleRequest(t, c, requests.MakeSignupRequest, t.Service.Signup)
}

func (t Transport) Login(c *gin.Context) {
	HandleRequest(t, c, requests.MakeLoginRequest, t.Service.Login)
}

//-------------------
//      USERS
//-------------------

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

//-------------------
//      POSTS
//-------------------

func (t Transport) CreateUserPost(c *gin.Context) {
	HandleRequest(t, c, requests.MakeCreateUserPostRequest, t.Service.CreateUserPost)
}

//-------------------
//       MISC
//-------------------

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"status": "API is up and running :)"})
}
