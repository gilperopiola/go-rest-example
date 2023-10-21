package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/common/requests"

	"github.com/gin-gonic/gin"
)

/*
	This is the entrypoint of the HTTP requests
		1. The HTTP Request arrivas and matches one of the URLs down below
		2. It calls ValidateToken if necessary
			a. If it's valid, it sets the user info in the context
		3. The corresponding function below in this file is called
		4. The HandleRequest method in transport.go is called
			a. It calls a function inside requests_xxx.go
				1. This function makes our Custom Request from the gin context
				2. It also validates it
			b. Now that we have our Custom Request, we use it to call the service
				1. The corresponding method in service_xxx.go is called
				2. Our Custom Request is converted to a Model
				3. The model's methods are called, which in turn call the repository_xxx.go file
					a. The repository_xxx.go file interacts with the database and returns a Model
				6. The resulting Model is converted to a Response Model
				7. The Response is returned to the HandleRequest method in transport.go
			c. The HandleRequest method in transport.go returns the HTTP Request
			d. Errors are also handled here, on HandleRequest
*/

//-----------------------------
//      ROUTES / ENDPOINTS
//-----------------------------

func (router *router) setPublicEndpoints(transport TransportLayer) {
	public := router.Group("/")
	{
		public.GET("/health", healthCheck)

		public.POST("/signup", transport.Signup)
		public.POST("/login", transport.Login)
	}
}

func (router *router) setUserEndpoints(transport TransportLayer, authI auth.AuthI) {
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

func (router *router) setAdminEndpoints(transport TransportLayer, authI auth.AuthI) {
	admin := router.Group("/admin", authI.ValidateToken(auth.AdminRole, false))
	{
		admin.POST("/user", transport.CreateUser)
		admin.GET("/users", transport.SearchUsers)
	}
}

//-------------------
//       AUTH
//-------------------

func (t transport) Signup(c *gin.Context) {
	HandleRequest(t, c, requests.MakeSignupRequest, t.Service().Signup)
}

func (t transport) Login(c *gin.Context) {
	HandleRequest(t, c, requests.MakeLoginRequest, t.Service().Login)
}

//-------------------
//      USERS
//-------------------

func (t transport) CreateUser(c *gin.Context) {
	HandleRequest(t, c, requests.MakeCreateUserRequest, t.Service().CreateUser)
}

func (t transport) GetUser(c *gin.Context) {
	HandleRequest(t, c, requests.MakeGetUserRequest, t.Service().GetUser)
}

func (t transport) UpdateUser(c *gin.Context) {
	HandleRequest(t, c, requests.MakeUpdateUserRequest, t.Service().UpdateUser)
}

func (t transport) DeleteUser(c *gin.Context) {
	HandleRequest(t, c, requests.MakeDeleteUserRequest, t.Service().DeleteUser)
}

func (t transport) SearchUsers(c *gin.Context) {
	HandleRequest(t, c, requests.MakeSearchUsersRequest, t.Service().SearchUsers)
}

//-------------------
//      POSTS
//-------------------

func (t transport) CreateUserPost(c *gin.Context) {
	HandleRequest(t, c, requests.MakeCreateUserPostRequest, t.Service().CreateUserPost)
}

//-------------------
//       MISC
//-------------------

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"status": "API is up and running :)"})
}
