package requests

import "github.com/gilperopiola/go-rest-example/pkg/common"

var (
	contextUserIDKey = "UserID"
	pathUserIDKey    = "user_id"
)

func MakeRequest[req All](c common.GinI, request req) (req, error) {
	if err := makeRequest(c, request); err != nil {
		return req(nil), err
	}
	return request, nil
}

type All interface {
	Build(c common.GinI) error
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

func makeRequest[req All](c common.GinI, request req) error {
	if err := request.Build(c); err != nil {
		return common.Wrap("request.Build", err)
	}
	if err := request.Validate(); err != nil {
		return common.Wrap("request.Validate", err)
	}
	return nil
}
