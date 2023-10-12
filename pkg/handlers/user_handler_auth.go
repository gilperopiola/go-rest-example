package handlers

import (
	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/mock"
)

type mockAuth struct {
	mock.Mock
}

func (m *mockAuth) GenerateToken(user entities.User, role entities.Role) (string, error) {
	args := m.Called(user, role)
	return args.String(0), args.Error(1)
}

func (m *mockAuth) ValidateToken(role entities.Role, shouldMatchUserID bool) gin.HandlerFunc {
	args := m.Called(role, shouldMatchUserID)
	return args.Get(0).(gin.HandlerFunc)
}

// -

func (h *UserHandler) GetAuthRole() entities.Role {
	if h.User.IsAdmin {
		return entities.AdminRole
	}
	return entities.UserRole
}

func (h *UserHandler) GenerateTokenString(a auth.AuthI) (string, error) {
	return a.GenerateToken(h.ToEntity(), h.GetAuthRole())
}
