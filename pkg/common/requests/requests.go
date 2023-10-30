package requests

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

type GinI interface {
	ShouldBindJSON(obj interface{}) error
	GetInt(key string) int
	Query(key string) (value string)
	DefaultQuery(key string, defaultValue string) string
}
