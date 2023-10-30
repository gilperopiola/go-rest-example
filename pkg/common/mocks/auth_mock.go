package mocks

import (
	"github.com/gilperopiola/go-rest-example/pkg/common/auth"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type AuthMock struct {
	*mock.Mock
}

func NewAuthMock() *AuthMock {
	return &AuthMock{Mock: &mock.Mock{}}
}

func (m *AuthMock) GenerateToken(id int, username, email string, role auth.Role) (string, error) {
	args := m.Called(id, username, email, role)
	return args.String(0), args.Error(1)
}

func (m *AuthMock) ValidateToken(role auth.Role, shouldMatchUserID bool) gin.HandlerFunc {
	args := m.Called(role, shouldMatchUserID)
	return args.Get(0).(gin.HandlerFunc)
}
