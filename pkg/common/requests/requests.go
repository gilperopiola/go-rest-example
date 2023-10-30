package requests

import "github.com/gilperopiola/go-rest-example/pkg/common"

type All interface {
	SignupRequest |
		LoginRequest |
		CreateUserRequest |
		GetUserRequest |
		UpdateUserRequest |
		DeleteUserRequest |
		SearchUsersRequest |
		ChangePasswordRequest |
		CreateUserPostRequest
}

type RequestMaker interface {
	Build(c GinI) error
	Validate() error
}

func makeRequest[req RequestMaker](c GinI, request req) (err error) {
	if err := request.Build(c); err != nil {
		return common.Wrap("request.Build", err)
	}
	if err = request.Validate(); err != nil {
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
