package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/utils"

	"github.com/gin-gonic/gin"
)

// Signup takes {username, email, password, repeat_password} and creates a new user
func (e Endpoints) Signup(c *gin.Context) {

	// Validate and get request
	signupRequest, err := makeSignupRequest(c)
	if err != nil {
		c.JSON(e.ErrorsMapper.Map(err))
		return
	}

	// Call service with that request
	signupResponse, err := e.Service.Signup(signupRequest)
	if err != nil {
		c.JSON(e.ErrorsMapper.Map(err))
		return
	}

	// Return OK
	c.JSON(returnOK(signupResponse))
}

func makeSignupRequest(c *gin.Context) (entities.SignupRequest, error) {

	// Bind & validate request
	var signupRequest entities.SignupRequest
	if err := c.ShouldBindJSON(&signupRequest); err != nil {
		return entities.SignupRequest{}, utils.JoinErrors(entities.ErrBindingRequest, err)
	}

	if err := signupRequest.Validate(); err != nil {
		return entities.SignupRequest{}, err
	}

	// Return request
	return signupRequest, nil
}
