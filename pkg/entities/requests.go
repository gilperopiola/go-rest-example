package entities

import (
	"errors"
)

type SignupRequest struct {
	Email          string
	Username       string
	Password       string
	RepeatPassword string
}

type LoginRequest struct {
	EmailOrUsername string
	Password        string
}

/* ----------------------------- */

func (req *SignupRequest) Validate() error {
	if req.Email == "" || req.Password == "" || req.RepeatPassword == "" {
		return errors.New("all fields required")
	}

	if req.Password != req.RepeatPassword {
		return errors.New("passwords don't match")
	}

	return nil
}
