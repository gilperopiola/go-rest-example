package requests

import "github.com/gilperopiola/go-rest-example/pkg/common"

var (
	contextUserIDKey = "UserID"
	pathUserIDKey    = "user_id"
)

type All interface {
	Build(c GinI) error
	Validate() error

	*SignupRequest |
		*LoginRequest |
		*CreateUserRequest |
		*GetUserRequest |
		*UpdateUserRequest |
		*DeleteUserRequest |
		*SearchUsersRequest |
		*ChangePasswordRequest |
		*CreateUserPostRequest
}

func MakeRequest[req All](c GinI, request req) (req, error) {
	if err := makeRequest(c, request); err != nil {
		return req(nil), err
	}
	return request, nil
}

func makeRequest[req All](c GinI, request req) error {
	if err := request.Build(c); err != nil {
		return common.Wrap("request.Build", err)
	}
	if err := request.Validate(); err != nil {
		return common.Wrap("request.Validate", err)
	}
	return nil
}

type GinI interface {
	ShouldBindJSON(obj interface{}) error
	GetInt(key string) int
	Query(key string) (value string)
	DefaultQuery(key string, defaultValue string) string
}

// bindRequestBody just binds the request body to the request struct
func bindRequestBody(c GinI, request interface{}) error {
	if err := c.ShouldBindJSON(&request); err != nil {
		return common.ErrBindingRequest
	}
	return nil
}
