package transport

import (
	"net/http"

	"github.com/gilperopiola/go-rest-example/pkg/entities"
	repository "github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/service"

	"github.com/gin-gonic/gin"
)

type Endpoints interface {
	Signup(c *gin.Context)
	Login(c *gin.Context)
}

type EndpointsHandler struct {
	Database     *repository.Database
	Service      *service.Service
	ErrorsMapper *ErrorsMapper
}

// Signup creates a new user and returns it
func (e EndpointsHandler) Signup(c *gin.Context) {

	// Bind request
	var signupRequest entities.SignupRequest
	c.BindJSON(&signupRequest)

	// Validate request
	if err := signupRequest.Validate(); err != nil {
		c.JSON(e.ErrorsMapper.Map(err))
		return
	}

	// Perform action
	response, err := e.Service.Signup(signupRequest)
	if err != nil {
		c.JSON(e.ErrorsMapper.Map(err))
		return
	}

	// Return OK
	c.JSON(http.StatusOK, HTTPResponse{Success: true, Content: response})
}

// Login takes {username, password}, checks if the user exists and returns it
func (e EndpointsHandler) Login(c *gin.Context) {
	c.JSON(http.StatusOK, nil)

	//	var user models.User
	//	c.BindJSON(&user)
	//
	//	if err := user.ValidateLogIn(); err != nil {
	//		c.JSON(http.StatusBadRequest, err.Error())
	//		return
	//	}
	//
	//	var databaseUser User
	//	database.Where("username = ?", user.Username).First(&databaseUser)
	//	if databaseUser.ID == 0 {
	//		c.JSON(http.StatusBadRequest, "wrong username")
	//		return
	//	}
	//
	//	if databaseUser.Password != hash(databaseUser.Email, user.Password) {
	//		c.JSON(http.StatusBadRequest, "wrong password")
	//		return
	//	}
	//
	//	databaseUser.Token = generateToken(databaseUser)
	//	databaseUser.Password = ""
	//	c.JSON(http.StatusOK, databaseUser)
	//}
	//
	//func validateToken(requiredRole string) gin.HandlerFunc {
	//	return func(c *gin.Context) {
	//		tokenString := c.Request.Header.Get("Authorization")
	//
	//		if len(tokenString) < 40 {
	//			c.JSON(http.StatusUnauthorized, "authentication error")
	//			c.Abort()
	//			return
	//		}
	//
	//		token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
	//			return []byte(config.JWT.SECRET), nil
	//		})
	//		if err != nil {
	//			c.JSON(http.StatusUnauthorized, "authentication error")
	//			c.Abort()
	//			return
	//		}
	//
	//		if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
	//			if claims.Subject != requiredRole {
	//				c.JSON(http.StatusUnauthorized, "authentication error")
	//				c.Abort()
	//				return
	//			}
	//
	//			c.Set("ID", claims.Id)
	//			c.Set("Email", claims.Audience)
	//			c.Set("Role", claims.Subject)
	//		} else {
	//			c.JSON(http.StatusUnauthorized, "authentication error")
	//			c.Abort()
	//		}
	//	}
	//}
	//
	//func generateToken(user User) string {
	//	var role string
	//	if user.Admin {
	//		role = "Admin"
	//	} else {
	//		role = "User"
	//	}
	//
	//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
	//		Id:        fmt.Sprint(user.ID),
	//		Audience:  user.Email,
	//		Subject:   role,
	//		IssuedAt:  time.Now().Unix(),
	//		ExpiresAt: time.Now().Add(time.Hour * 24 * time.Duration(config.JWT.SESSION_DURATION)).Unix(),
	//	})
	//	tokenString, _ := token.SignedString([]byte(config.JWT.SECRET))
	//	return tokenString
	//}
	//
	//func generateTestingToken(role string) string {
	//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
	//		Id:        "1",
	//		Audience:  "test@test.com",
	//		Subject:   role,
	//		IssuedAt:  time.Now().Unix(),
	//		ExpiresAt: time.Now().Add(time.Minute).Unix(),
	//	})
	//	tokenString, _ := token.SignedString([]byte(config.JWT.SECRET))
	//	return tokenString
}
