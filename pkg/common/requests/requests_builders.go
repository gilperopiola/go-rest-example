package requests

import (
	"strconv"

	"github.com/gilperopiola/go-rest-example/pkg/common"
)

type GinI interface {
	ShouldBindJSON(obj interface{}) error
	GetInt(key string) int
	Query(key string) (value string)
	DefaultQuery(key string, defaultValue string) string
}

//--------------------------
//	    AUTH BUILDERS
//--------------------------

func (req *SignupRequest) Build(c GinI) error {
	return buildDefaultRequestBody(c, req)
}

func (req *LoginRequest) Build(c GinI) error {
	return buildDefaultRequestBody(c, req)
}

//--------------------------
//	    USER BUILDERS
//--------------------------

func (req *CreateUserRequest) Build(c GinI) error {
	return buildDefaultRequestBody(c, req)
}

func (req *GetUserRequest) Build(c GinI) error {
	userToGetID, err := getIntFromContext(c, "ID")
	if err != nil {
		return err
	}

	req.ID = userToGetID

	return nil
}

func (req *UpdateUserRequest) Build(c GinI) error {
	if err := c.ShouldBindJSON(&req); err != nil {
		return common.ErrBindingRequest
	}

	userToUpdateID, err := getIntFromContext(c, "ID")
	if err != nil {
		return err
	}

	req.ID = userToUpdateID
	return nil
}

func (req *DeleteUserRequest) Build(c GinI) error {
	userToDeleteID, err := getIntFromContext(c, "ID")
	if err != nil {
		return err
	}

	req.ID = userToDeleteID
	return nil
}

func (req *SearchUsersRequest) Build(c GinI) error {
	var err error

	req.Username = c.Query("username")

	req.Page, err = strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		return common.ErrInvalidValue
	}

	req.PerPage, err = strconv.Atoi(c.DefaultQuery("per_page", "10"))
	if err != nil {
		return common.ErrInvalidValue
	}

	return nil
}

//-----------------------------
//	    USER POST BUILDERS
//-----------------------------

func (req *CreateUserPostRequest) Build(c GinI) error {
	if err := c.ShouldBindJSON(&req); err != nil {
		return common.ErrBindingRequest
	}

	postOwnerID, err := getIntFromContext(c, "ID")
	if err != nil {
		return err
	}

	req.UserID = postOwnerID
	return nil
}

//--------------------
//	    HELPERS
//--------------------

// buildDefaultRequestBody just binds the request body to the request struct
func buildDefaultRequestBody(c GinI, request interface{}) error {
	if err := c.ShouldBindJSON(&request); err != nil {
		return common.ErrBindingRequest
	}
	return nil
}

func getIntFromContext(c GinI, key string) (int, error) {
	value := c.GetInt(key)
	if value == 0 {
		return 0, common.ErrReadingValueFromCtx
	}
	return value, nil
}
