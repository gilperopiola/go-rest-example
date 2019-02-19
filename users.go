package main

import (
	"time"
)

type User struct {
	ID          uint      `json:"id" gorm:"auto_increment;unique;not null"`
	Username    string    `json:"username" gorm:"unique;not null"`
	Email       string    `json:"email" gorm:"unique;not null"`
	Password    string    `json:"password,omitempty" gorm:"not null"`
	Admin       bool      `json:"admin" gorm:"default: 0"`
	Active      bool      `json:"active" gorm:"default: 1"`
	DateCreated time.Time `json:"date_created" gorm:"default: current_timestamp"`

	Token string `json:"token" gorm:"-"`
}
