package main

import "time"

type User struct {
	ID          int
	Email       string
	Password    string
	Enabled     bool
	DateCreated time.Time
}
