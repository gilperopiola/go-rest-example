package models

import (
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/utils"
)

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	IsAdmin   bool
	Deleted   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) OverwriteFields(username, email string) {
	if username != "" {
		user.Username = username
	}
	if email != "" {
		user.Email = email
	}
}

func (user *User) PasswordMatches(password string) bool {
	return user.Password == utils.Hash(user.Email, password)
}
