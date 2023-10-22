package requests

import (
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"
)

//---------------------------
//    AUTH REQUEST STRUCTS
//---------------------------

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

//----------------------------
//     AUTH REQUEST MAKERS
//----------------------------

func MakeSignupRequest(c GinI) (request SignupRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return SignupRequest{}, common.Wrap("makeSignupRequest", common.ErrBindingRequest)
	}

	if err = request.Validate(); err != nil {
		return SignupRequest{}, common.Wrap("makeSignupRequest", err)
	}

	return request, nil
}

func MakeLoginRequest(c GinI) (request LoginRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return LoginRequest{}, common.Wrap("makeLoginRequest", common.ErrBindingRequest)
	}

	if err = request.Validate(); err != nil {
		return LoginRequest{}, common.Wrap("makeLoginRequest", err)
	}

	return request, nil
}

//-------------------------------
//     REQUEST TO USER MODEL
//-------------------------------

func (r *SignupRequest) ToUserModel() models.User {
	return models.User{
		Username: r.Username,
		Email:    r.Email,
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
		return common.ErrPasswordsDontMatch
	}

	return nil
}

func (req LoginRequest) Validate() error {
	if req.UsernameOrEmail == "" || req.Password == "" {
		return common.ErrAllFieldsRequired
	}

	return nil
}
