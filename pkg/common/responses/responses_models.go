package responses

import "time"

type User struct {
	ID        int        `json:"id,omitempty"`
	Username  string     `json:"username,omitempty"`
	Email     string     `json:"email,omitempty"`
	Password  string     `json:"password,omitempty"`
	IsAdmin   bool       `json:"is_admin,omitempty"`
	Details   UserDetail `json:"details,omitempty"`
	Deleted   bool       `json:"deleted,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
}

type UserDetail struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}
