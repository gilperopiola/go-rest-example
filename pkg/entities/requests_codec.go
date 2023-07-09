package entities

import "regexp"

func (req *LoginRequest) ToUserCredentials() UserCredentials {

	out := UserCredentials{Password: req.Password}

	// If it's an email, login with email. Otherwise login with username
	if matchesEmailFormat, _ := regexp.MatchString(VALID_EMAIL_REGEX, req.UsernameOrEmail); matchesEmailFormat {
		out.Email = req.UsernameOrEmail
	} else {
		out.Username = req.UsernameOrEmail
	}

	return out
}
