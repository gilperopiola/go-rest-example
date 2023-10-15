package entities

import "time"

// - Auth Roles
type Role string

const (
	AnyRole   Role = "any"
	UserRole  Role = "user"
	AdminRole Role = "admin"
)

// - User entity
type User struct {
	ID        int       `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	IsAdmin   bool      `json:"is_admin,omitempty"`
	Deleted   bool      `json:"deleted,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
