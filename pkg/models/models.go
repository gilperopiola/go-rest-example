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
