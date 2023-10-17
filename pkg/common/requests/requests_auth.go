package requests

import (
	"fmt"
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	customErrors "github.com/gilperopiola/go-rest-example/pkg/common/errors"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"
)

//-----------------------
//    REQUEST STRUCTS
//-----------------------

type SignupRequest struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`

	// User Detail
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type LoginRequest struct {
	UsernameOrEmail string `json:"username_or_email"`
	Password        string `json:"password"`
}

//-------------------------
//     REQUEST MAKERS
//-------------------------

func MakeSignupRequest(c GinI) (request SignupRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return SignupRequest{}, common.Wrap(fmt.Errorf("makeSignupRequest"), customErrors.ErrBindingRequest)
	}

	if err = request.Validate(); err != nil {
		return SignupRequest{}, common.Wrap(fmt.Errorf("makeSignupRequest"), err)
	}

	return request, nil
}

func MakeLoginRequest(c GinI) (request LoginRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return LoginRequest{}, common.Wrap(fmt.Errorf("makeLoginRequest"), customErrors.ErrBindingRequest)
	}

	if err = request.Validate(); err != nil {
		return LoginRequest{}, common.Wrap(fmt.Errorf("makeLoginRequest"), err)
	}

	return request, nil
}

//----------------------------
//     REQUEST TO MODEL
//----------------------------

func (r *SignupRequest) ToUserModel() models.User {
	return models.User{
		Email:    r.Email,
		Username: r.Username,
		Password: r.Password,
		Deleted:  false,
		Details: models.UserDetail{
			FirstName: r.FirstName,
			LastName:  r.LastName,
		},
		IsAdmin:   false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (r *LoginRequest) ToUserModel() models.User {
	user := models.User{
		Password: r.Password,
	}

	if validEmailRegex.MatchString(r.UsernameOrEmail) {
		user.Email = r.UsernameOrEmail
	} else {
		user.Username = r.UsernameOrEmail
	}

	return user
}

//--------------------------
//	 REQUEST VALIDATIONS
//--------------------------

func (req SignupRequest) Validate() error {
	if err := validateUsernameEmailAndPassword(req.Username, req.Email, req.Password); err != nil {
		return err
	}

	if req.Password != req.RepeatPassword {
		return customErrors.ErrPasswordsDontMatch
	}

	return nil
}

func (req LoginRequest) Validate() error {
	if req.UsernameOrEmail == "" || req.Password == "" {
		return customErrors.ErrAllFieldsRequired
	}

	return nil
}
