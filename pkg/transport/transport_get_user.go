package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/utils"

	"github.com/gin-gonic/gin"
)

// Get User gets the info of a user
func (e Transport) GetUser(c *gin.Context) {

	// Validate and get request
	getUserRequest, err := makeGetUserRequest(c)
	if err != nil {
		c.JSON(e.ErrorsMapper.Map(err))
		return
	}

	// Call service with that request
	getUserResponse, err := e.Service.GetUser(getUserRequest)
	if err != nil {
		c.JSON(e.ErrorsMapper.Map(err))
		return
	}

	// Return OK
	c.JSON(returnOK(getUserResponse))
}

func makeGetUserRequest(c *gin.Context) (entities.GetUserRequest, error) {

	// Get info from context and URL, check if user IDs match
	userToGetID, err := checkIfRequestUserIDsMatch(c)
	if err != nil {
		return entities.GetUserRequest{}, err
	}

	// Create & validate request
	getUserRequest := entities.GetUserRequest{ID: userToGetID}

	if err := getUserRequest.Validate(); err != nil {
		return entities.GetUserRequest{}, err
	}

	// Return request
	return getUserRequest, nil
}

func checkIfRequestUserIDsMatch(c *gin.Context) (int, error) {

	// Get info from context and URL
	loggedUserID, err := utils.GetIntFromContext(c, "ID")
	if err != nil {
		return 0, err
	}

	userToGetID, err := utils.GetIntFromURLParams(c, "user_id")
	if err != nil {
		return 0, err
	}

	// Check if the logged user has the same ID as the one to get
	if loggedUserID != userToGetID {
		return 0, entities.ErrUnauthorized
	}

	// Return user ID
	return userToGetID, nil
}
