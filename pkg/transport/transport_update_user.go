package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/utils"
	"github.com/gin-gonic/gin"
)

// Login takes {username_or_email, password}, checks if the user exists and returns it
func (e Endpoints) UpdateUser(c *gin.Context) {

	// Validate and get request
	updateUserRequest, err := makeUpdateUserRequest(c)
	if err != nil {
		c.JSON(e.ErrorsMapper.Map(err))
		return
	}

	// Call service with the request
	updateUserResponse, err := e.Service.UpdateUser(updateUserRequest)
	if err != nil {
		c.JSON(e.ErrorsMapper.Map(err))
		return
	}

	// Return OK
	c.JSON(returnOK(updateUserResponse))
}

func makeUpdateUserRequest(c *gin.Context) (entities.UpdateUserRequest, error) {

	// Validate and get User ID
	userToUpdateID, err := checkIfRequestUserIDsMatch(c)
	if err != nil {
		return entities.UpdateUserRequest{}, err
	}

	// Bind request
	var updateUserRequest entities.UpdateUserRequest
	if err := c.BindJSON(&updateUserRequest); err != nil {
		return entities.UpdateUserRequest{}, utils.JoinErrors(entities.ErrBindingRequest, err)
	}

	// Assign User ID to request
	updateUserRequest.ID = userToUpdateID

	// Validate request
	if err := updateUserRequest.Validate(); err != nil {
		return entities.UpdateUserRequest{}, err
	}

	// Return request
	return updateUserRequest, nil
}
