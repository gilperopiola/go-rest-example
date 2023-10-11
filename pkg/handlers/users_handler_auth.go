package handlers

import (
	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/entities"
)

func (h *UserHandler) GetAuthRole() entities.Role {
	if h.User.IsAdmin {
		return entities.AdminRole
	}
	return entities.UserRole
}

func (h *UserHandler) GenerateTokenString(a auth.AuthI) (string, error) {
	return a.GenerateToken(h.ToEntity(), h.GetAuthRole())
}
