package requests

import (
	"fmt"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	customErrors "github.com/gilperopiola/go-rest-example/pkg/common/errors"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"
)

//-----------------------
//    REQUEST STRUCTS
//-----------------------

type CreateUserPostRequest struct {
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

//-------------------------
//     REQUEST MAKERS
//-------------------------

func MakeCreateUserPostRequest(c GinI) (request CreateUserPostRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return CreateUserPostRequest{}, common.Wrap(fmt.Errorf("makeCreateUserPostRequest"), customErrors.ErrBindingRequest)
	}

	postOwnerID, err := getIntFromContext(c, "ID")
	if err != nil {
		return CreateUserPostRequest{}, common.Wrap(fmt.Errorf("makeCreateUserPostRequest"), err)
	}

	request.UserID = postOwnerID

	if err = request.Validate(); err != nil {
		return CreateUserPostRequest{}, common.Wrap(fmt.Errorf("makeCreateUserPostRequest"), err)
	}

	return request, nil
}

//----------------------------
//     REQUEST TO MODEL
//----------------------------

func (r *CreateUserPostRequest) ToUserPostModel() models.UserPost {
	return models.UserPost{
		UserID: r.UserID,
		Title:  r.Title,
		Body:   r.Body,
	}
}

//--------------------------
//	 REQUEST VALIDATIONS
//--------------------------

func (req CreateUserPostRequest) Validate() error {
	if req.UserID == 0 || req.Title == "" {
		return customErrors.ErrAllFieldsRequired
	}
	return nil
}
