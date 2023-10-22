package responses

// These aren't the HTTP Responses that the API will return, but the responses that the Service Layer
// returns to the Transport Layer.

//-------------------
//      AUTH
//-------------------

type SignupResponse struct {
	User User `json:"user"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

//--------------------
//      USERS
//--------------------

type CreateUserResponse struct {
	User User `json:"user"`
}

type GetUserResponse struct {
	User User `json:"user"`
}

type UpdateUserResponse struct {
	User User `json:"user"`
}

type DeleteUserResponse struct {
	User User `json:"user"`
}

type SearchUsersResponse struct {
	Users   []User `json:"users"`
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
}

//-----------------------
//      USER POSTS
//-----------------------

type CreateUserPostResponse struct {
	UserPost UserPost `json:"user_post"`
}
