package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"

	"github.com/gin-gonic/gin"
)

// DeleteUser deletes the logged user
func (e Endpoints) DeleteUser(c *gin.Context) {

	// Validate and get request
	deleteUserRequest, err := makeDeleteUserRequest(c)
	if err != nil {
		c.JSON(e.ErrorsMapper.Map(err))
		return
	}

	// Call service with that request
	deleteUserResponse, err := e.Service.DeleteUser(deleteUserRequest)
	if err != nil {
		c.JSON(e.ErrorsMapper.Map(err))
		return
	}

	// Return OK
	c.JSON(returnOK(deleteUserResponse))
}

func makeDeleteUserRequest(c *gin.Context) (entities.DeleteUserRequest, error) {

	// Get info from context and URL, check if user IDs match
	userToDeleteID, err := checkIfRequestUserIDsMatch(c)
	if err != nil {
		return entities.DeleteUserRequest{}, err
	}

	// Create & validate request
	deleteUserRequest := entities.DeleteUserRequest{ID: userToDeleteID}

	if err := deleteUserRequest.Validate(); err != nil {
		return entities.DeleteUserRequest{}, err
	}

	// Return request
	return deleteUserRequest, nil
}
