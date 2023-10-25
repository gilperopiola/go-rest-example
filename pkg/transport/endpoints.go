package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/common/auth"
	"github.com/gilperopiola/go-rest-example/pkg/common/requests"
	"github.com/prometheus/client_golang/prometheus/promhttp"

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

func (router *router) setEndpoints(transport TransportLayer, authI auth.AuthI) {
	router.GET("/health", healthCheck)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	v1 := router.Group("/v1")
	{
		router.setV1Endpoints(v1, transport, authI)
	}

	/* Profiling
	pprofGroup := router.Group("/debug/pprof")
	{
		pprofGroup.GET("/", gin.WrapF(pprof.Index))
		pprofGroup.GET("/cmdline", gin.WrapF(pprof.Cmdline))
		pprofGroup.GET("/profile", gin.WrapF(pprof.Profile))
		pprofGroup.POST("/symbol", gin.WrapF(pprof.Symbol))
		pprofGroup.GET("/symbol", gin.WrapF(pprof.Symbol))
		pprofGroup.GET("/trace", gin.WrapF(pprof.Trace))
		pprofGroup.GET("/allocs", gin.WrapF(pprof.Handler("allocs").ServeHTTP))
		pprofGroup.GET("/block", gin.WrapF(pprof.Handler("block").ServeHTTP))
		pprofGroup.GET("/goroutine", gin.WrapF(pprof.Handler("goroutine").ServeHTTP))
		pprofGroup.GET("/heap", gin.WrapF(pprof.Handler("heap").ServeHTTP))
		pprofGroup.GET("/mutex", gin.WrapF(pprof.Handler("mutex").ServeHTTP))
		pprofGroup.GET("/threadcreate", gin.WrapF(pprof.Handler("threadcreate").ServeHTTP))
	}*/
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"status": "API is up and running :)"})
}

func (router *router) setV1Endpoints(v1 *gin.RouterGroup, transport TransportLayer, authI auth.AuthI) {

	// Auth
	v1.POST("/signup", transport.Signup)
	v1.POST("/login", transport.Login)

	// Users
	users := v1.Group("/users", authI.ValidateToken(auth.AnyRole, true))
	{
		users.GET("/:user_id", transport.GetUser)
		users.PATCH("/:user_id", transport.UpdateUser)
		users.DELETE("/:user_id", transport.DeleteUser)

		// User posts
		posts := users.Group("/:user_id/posts")
		{
			posts.POST("", transport.CreateUserPost)
		}
	}

	// Admin endpoints
	admin := v1.Group("/admin", authI.ValidateToken(auth.AdminRole, false))
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
