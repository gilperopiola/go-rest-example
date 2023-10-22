package requests

import (
	"regexp"
	"strconv"
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	customErrors "github.com/gilperopiola/go-rest-example/pkg/common/errors"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"
)

type GinI interface {
	ShouldBindJSON(obj interface{}) error
	GetInt(key string) int
	Query(key string) (value string)
	DefaultQuery(key string, defaultValue string) string
}

const (
	usernameMinLength = 4
	usernameMaxLength = 32
	passwordMinLength = 8
	passwordMaxLength = 64
)

var (
	validEmailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
)

//-----------------------
//    REQUEST STRUCTS
//-----------------------

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`

	// User Detail
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type GetUserRequest struct {
	ID int `json:"id"`
}

type UpdateUserRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`

	// User Detail
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}

type DeleteUserRequest struct {
	ID int `json:"id"`
}

type SearchUsersRequest struct {
	Username string `json:"username"`
	Page     int    `json:"page"`
	PerPage  int    `json:"per_page"`
}

//-------------------------
//     REQUEST MAKERS
//-------------------------

func MakeCreateUserRequest(c GinI) (request CreateUserRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return CreateUserRequest{}, common.Wrap("makeCreateUserRequest", customErrors.ErrBindingRequest)
	}

	if err = request.Validate(); err != nil {
		return CreateUserRequest{}, common.Wrap("makeCreateUserRequest", err)
	}

	return request, nil
}

func MakeGetUserRequest(c GinI) (request GetUserRequest, err error) {
	userToGetID, err := getIntFromContext(c, "ID")
	if err != nil {
		return GetUserRequest{}, common.Wrap("makeGetUserRequest", err)
	}

	request.ID = userToGetID

	if err = request.Validate(); err != nil {
		return GetUserRequest{}, common.Wrap("makeGetUserRequest", err)
	}

	return request, nil
}

func MakeUpdateUserRequest(c GinI) (request UpdateUserRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return UpdateUserRequest{}, common.Wrap("makeUpdateUserRequest", customErrors.ErrBindingRequest)
	}

	userToUpdateID, err := getIntFromContext(c, "ID")
	if err != nil {
		return UpdateUserRequest{}, common.Wrap("makeUpdateUserRequest", err)
	}

	request.ID = userToUpdateID

	if err = request.Validate(); err != nil {
		return UpdateUserRequest{}, common.Wrap("makeUpdateUserRequest", err)
	}

	return request, nil
}

func MakeDeleteUserRequest(c GinI) (request DeleteUserRequest, err error) {
	userToDeleteID, err := getIntFromContext(c, "ID")
	if err != nil {
		return DeleteUserRequest{}, common.Wrap("makeDeleteUserRequest", err)
	}

	request.ID = userToDeleteID

	if err = request.Validate(); err != nil {
		return DeleteUserRequest{}, common.Wrap("makeDeleteUserRequest", err)
	}

	return request, nil
}

func MakeSearchUsersRequest(c GinI) (request SearchUsersRequest, err error) {
	request.Username = c.Query("username")
	request.Page, err = strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		return SearchUsersRequest{}, common.Wrap("makeSearchUsersRequest", customErrors.ErrInvalidValue)
	}
	request.PerPage, err = strconv.Atoi(c.DefaultQuery("per_page", "10"))
	if err != nil {
		return SearchUsersRequest{}, common.Wrap("makeSearchUsersRequest", customErrors.ErrInvalidValue)
	}

	if err = request.Validate(); err != nil {
		return SearchUsersRequest{}, common.Wrap("makeSearchUsersRequest", err)
	}

	return request, nil
}

//----------------------------
//     REQUEST TO MODEL
//----------------------------

func (r *CreateUserRequest) ToUserModel() models.User {
	return models.User{
		Email:    r.Email,
		Username: r.Username,
		Password: r.Password,
		Deleted:  false,
		Details: models.UserDetail{
			FirstName: r.FirstName,
			LastName:  r.LastName,
		},
		IsAdmin:   r.IsAdmin,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (r *GetUserRequest) ToUserModel() models.User {
	return models.User{ID: r.ID}
}

func (r *UpdateUserRequest) ToUserModel() models.User {
	var (
		firstName = ""
		lastName  = ""
	)
	if r.FirstName != nil {
		firstName = *r.FirstName
	}
	if r.LastName != nil {
		lastName = *r.LastName
	}

	return models.User{
		ID:       r.ID,
		Username: r.Username,
		Email:    r.Email,
		Details: models.UserDetail{
			FirstName: firstName,
			LastName:  lastName,
		},
	}
}

func (r *DeleteUserRequest) ToUserModel() models.User {
	return models.User{ID: r.ID}
}

func (r *SearchUsersRequest) ToUserModel() models.User {
	return models.User{Username: r.Username}
}

//--------------------------
//	 REQUEST VALIDATIONS
//--------------------------

func (req CreateUserRequest) Validate() error {
	return validateUsernameEmailAndPassword(req.Username, req.Email, req.Password)
}

func (req GetUserRequest) Validate() error {
	if req.ID == 0 {
		return customErrors.ErrAllFieldsRequired
	}
	return nil
}

func (req UpdateUserRequest) Validate() error {
	if req.ID == 0 || (req.Email == "" && req.Username == "") {
		return customErrors.ErrAllFieldsRequired
	}

	if req.Email != "" && !validEmailRegex.MatchString(req.Email) {
		return customErrors.ErrInvalidEmailFormat
	}

	if req.Username != "" {
		if len(req.Username) < usernameMinLength || len(req.Username) > usernameMaxLength {
			return customErrors.ErrInvalidUsernameLength
		}
	}

	return nil
}

func (req DeleteUserRequest) Validate() error {
	if req.ID == 0 {
		return customErrors.ErrAllFieldsRequired
	}

	return nil
}

func (req SearchUsersRequest) Validate() error {
	if req.Page < 0 || req.PerPage <= 0 {
		return customErrors.ErrAllFieldsRequired
	}

	return nil
}

//--------------------------
//	      HELPERS
//--------------------------

func getIntFromContext(c GinI, key string) (int, error) {
	value := c.GetInt(key)
	if value == 0 {
		return 0, customErrors.ErrReadingValueFromCtx
	}
	return value, nil
}

func validateUsernameEmailAndPassword(username, email, password string) error {
	if email == "" || username == "" || password == "" {
		return customErrors.ErrAllFieldsRequired
	}

	if !validEmailRegex.MatchString(email) {
		return customErrors.ErrInvalidEmailFormat
	}

	if len(username) < usernameMinLength || len(username) > usernameMaxLength {
		return customErrors.ErrInvalidUsernameLength
	}

	if len(password) < passwordMinLength || len(password) > passwordMaxLength {
		return customErrors.ErrInvalidPasswordLength
	}

	return nil
}
