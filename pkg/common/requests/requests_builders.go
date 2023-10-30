package requests

import (
	"strconv"

	"github.com/gilperopiola/go-rest-example/pkg/common"
)

var (
	contextUserIDKey = "UserID"
	pathUserIDKey    = "user_id"
)

//--------------------------
//	    AUTH BUILDERS
//--------------------------

func (req *SignupRequest) Build(c GinI) error {
	return bindRequestBody(c, req)
}

func (req *LoginRequest) Build(c GinI) error {
	return bindRequestBody(c, req)
}

//--------------------------
//	    USER BUILDERS
//--------------------------

func (req *CreateUserRequest) Build(c GinI) error {
	return bindRequestBody(c, req)
}

func (req *GetUserRequest) Build(c GinI) error {
	req.UserID = c.GetInt(contextUserIDKey)
	return nil
}

func (req *UpdateUserRequest) Build(c GinI) error {
	if err := bindRequestBody(c, req); err != nil {
		return err
	}

	req.UserID = c.GetInt(contextUserIDKey)

	return nil
}

func (req *DeleteUserRequest) Build(c GinI) error {
	req.UserID = c.GetInt(contextUserIDKey)
	return nil
}

func (req *SearchUsersRequest) Build(c GinI) error {
	var (
		err            = error(nil)
		defaultPage    = "0"
		defaultPerPage = "10"
	)

	req.Username = c.Query("username")

	req.Page, err = strconv.Atoi(c.DefaultQuery("page", defaultPage))
	if err != nil {
		return common.ErrInvalidValue
	}

	req.PerPage, err = strconv.Atoi(c.DefaultQuery("per_page", defaultPerPage))
	if err != nil {
		return common.ErrInvalidValue
	}

	return nil
}

func (req *ChangePasswordRequest) Build(c GinI) error {
	err := bindRequestBody(c, req)
	if err != nil {
		return err
	}

	req.UserID = c.GetInt(contextUserIDKey)

	return nil
}

//-----------------------------
//	    USER POST BUILDERS
//-----------------------------

func (req *CreateUserPostRequest) Build(c GinI) error {
	err := bindRequestBody(c, req)
	if err != nil {
		return err
	}

	req.UserID = c.GetInt(contextUserIDKey)

	return nil
}

//--------------------
//	    HELPERS
//--------------------

// bindRequestBody just binds the request body to the request struct
func bindRequestBody(c GinI, request interface{}) error {
	if err := c.ShouldBindJSON(&request); err != nil {
		return common.ErrBindingRequest
	}
	return nil
}
