package models

import "time"

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

func (user *User) FillFields(username, email string) {
	if username != "" {
		user.Username = username
	}
	if email != "" {
		user.Email = email
	}
}
