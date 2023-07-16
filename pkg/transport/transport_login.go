package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/utils"
	"github.com/gin-gonic/gin"
)

// Login takes {username_or_email, password}, checks if the user exists and returns it
func (e Endpoints) Login(c *gin.Context) {

	// Validate and get request
	loginRequest, err := makeLoginRequest(c)
	if err != nil {
		c.JSON(e.ErrorsMapper.Map(err))
		return
	}

	// Transform request to user credentials, checking if the input was a username or an email
	userCredentials := loginRequest.ToUserCredentials()

	// Call service with those credentials
	loginResponse, err := e.Service.Login(userCredentials)
	if err != nil {
		c.JSON(e.ErrorsMapper.Map(err))
		return
	}

	// Return OK
	c.JSON(returnOK(loginResponse))
}

func makeLoginRequest(c *gin.Context) (entities.LoginRequest, error) {

	// Bind request
	var loginRequest entities.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		return entities.LoginRequest{}, utils.JoinErrors(entities.ErrBindingRequest, err)
	}

	// Validate request
	if err := loginRequest.Validate(); err != nil {
		return entities.LoginRequest{}, err
	}

	// Return request
	return loginRequest, nil
}
