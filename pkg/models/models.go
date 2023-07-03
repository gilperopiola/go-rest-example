package models

import "time"

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	Deleted   bool
	IsAdmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
