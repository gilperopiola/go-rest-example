package requests

//-------------------
//      STRUCTS
//-------------------

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
	UserID int `json:"user_id"`
}

type UpdateUserRequest struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`

	// User Detail
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}

type DeleteUserRequest struct {
	UserID int `json:"user_id"`
}

type SearchUsersRequest struct {
	Username string `json:"username"`
	Page     int    `json:"page"`
	PerPage  int    `json:"per_page"`
}

type ChangePasswordRequest struct {
	UserID         int    `json:"user_id"`
	OldPassword    string `json:"old_password"`
	NewPassword    string `json:"new_password"`
	RepeatPassword string `json:"repeat_password"`
}

type CreateUserPostRequest struct {
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

//-------------------
//      MAKERS
//-------------------

func MakeSignupRequest(c GinI) (SignupRequest, error) {
	var request SignupRequest
	if err := makeRequest(c, &request); err != nil {
		return SignupRequest{}, err
	}
	return request, nil
}

func MakeLoginRequest(c GinI) (LoginRequest, error) {
	var request LoginRequest
	if err := makeRequest(c, &request); err != nil {
		return LoginRequest{}, err
	}
	return request, nil
}

func MakeCreateUserRequest(c GinI) (CreateUserRequest, error) {
	var request CreateUserRequest
	if err := makeRequest(c, &request); err != nil {
		return CreateUserRequest{}, err
	}
	return request, nil
}

func MakeGetUserRequest(c GinI) (GetUserRequest, error) {
	var request GetUserRequest
	if err := makeRequest(c, &request); err != nil {
		return GetUserRequest{}, err
	}
	return request, nil
}

func MakeUpdateUserRequest(c GinI) (UpdateUserRequest, error) {
	var request UpdateUserRequest
	if err := makeRequest(c, &request); err != nil {
		return UpdateUserRequest{}, err
	}
	return request, nil
}

func MakeDeleteUserRequest(c GinI) (DeleteUserRequest, error) {
	var request DeleteUserRequest
	if err := makeRequest(c, &request); err != nil {
		return DeleteUserRequest{}, err
	}
	return request, nil
}

func MakeSearchUsersRequest(c GinI) (SearchUsersRequest, error) {
	var request SearchUsersRequest
	if err := makeRequest(c, &request); err != nil {
		return SearchUsersRequest{}, err
	}
	return request, nil
}

func MakeChangePasswordRequest(c GinI) (ChangePasswordRequest, error) {
	var request ChangePasswordRequest
	if err := makeRequest(c, &request); err != nil {
		return ChangePasswordRequest{}, err
	}
	return request, nil
}

func MakeCreateUserPostRequest(c GinI) (CreateUserPostRequest, error) {
	var request CreateUserPostRequest
	if err := makeRequest(c, &request); err != nil {
		return CreateUserPostRequest{}, err
	}
	return request, nil
}
