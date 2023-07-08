package entities

import "regexp"

func (req *LoginRequest) ToUserCredentials() UserCredentials {

	out := UserCredentials{Password: req.Password}

	// If it's an email, login with email. Otherwise login with username
	if matchesEmailFormat, _ := regexp.MatchString(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`, req.UsernameOrEmail); matchesEmailFormat {
		out.Email = req.UsernameOrEmail
	} else {
		out.Username = req.UsernameOrEmail
	}

	return out
}
