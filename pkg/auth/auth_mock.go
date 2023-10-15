package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MockAuth struct {
	mock.Mock
}

func (m *MockAuth) GenerateToken(user User, role Role) (string, error) {
	args := m.Called(user, role)
	return args.String(0), args.Error(1)
}

func (m *MockAuth) ValidateToken(role Role, shouldMatchUserID bool) gin.HandlerFunc {
	args := m.Called(role, shouldMatchUserID)
	return args.Get(0).(gin.HandlerFunc)
}
