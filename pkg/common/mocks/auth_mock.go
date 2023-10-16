package mocks

import (
	"github.com/gilperopiola/go-rest-example/pkg/auth"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MockAuth struct {
	mock.Mock
}

func (m *MockAuth) GenerateToken(user auth.User, role auth.Role) (string, error) {
	args := m.Called(user, role)
	return args.String(0), args.Error(1)
}

func (m *MockAuth) ValidateToken(role auth.Role, shouldMatchUserID bool) gin.HandlerFunc {
	args := m.Called(role, shouldMatchUserID)
	return args.Get(0).(gin.HandlerFunc)
}
