package responses

import "time"

type User struct {
	ID        int        `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	IsAdmin   bool       `json:"is_admin,omitempty"`
	Details   UserDetail `json:"details"`
	Posts     []UserPost `json:"posts"`
	Deleted   bool       `json:"deleted,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
}

type UserDetail struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserPost struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
