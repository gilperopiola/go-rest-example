package main

import (
	"errors"
	"strings"
	"time"
)

type UserActions interface {
	ValidateSignUp() error
	ValidateLogIn() error
}

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

func (user *User) ValidateSignUp() error {
	if len(user.Username) == 0 || len(user.Email) == 0 || len(user.Password) == 0 {
		return errors.New("all fields required")
	}

	if len(user.Username) < config.USERS.USERNAME_MIN_CHARACTERS || len(user.Username) > config.USERS.USERNAME_MAX_CHARACTERS {
		return errors.New("username must have between " + string(config.USERS.USERNAME_MIN_CHARACTERS) + " and " + string(config.USERS.USERNAME_MAX_CHARACTERS) + " characters")
	}

	if !strings.Contains(user.Email, "@") {
		return errors.New("email format invalid")
	}

	return nil
}

func (user *User) ValidateLogIn() error {
	if len(user.Username) == 0 || len(user.Password) == 0 {
		return errors.New("both fields required")
	}
	return nil
}
