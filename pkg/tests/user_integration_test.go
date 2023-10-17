package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/common/config"
	"github.com/gilperopiola/go-rest-example/pkg/common/requests"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/service"
	"github.com/gilperopiola/go-rest-example/pkg/transport"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUsersCRUDIntegrationTest(t *testing.T) {

	// Prepare
	config := config.NewTestConfig()
	database := repository.NewDatabase(config.Database(), nopLogger{})
	repository := repository.NewRepository(database)
	service := service.NewService(repository, nopAuth{}, config)
	endpoints := transport.NewTransport(service, transport.NewErrorsMapper(nopLogger{}))

	// Signup
	c := makeTestContextWithHTTPRequest(map[string]string{
		"username":        "test",
		"email":           "test@email.com",
		"password":        "password",
		"repeat_password": "password",
	})
	transport.HandleRequest(endpoints, c, requests.MakeSignupRequest, endpoints.Service.Signup)
	statusCode := c.Writer.Status()
	assert.Equal(t, http.StatusOK, statusCode)

	// Login
	c = makeTestContextWithHTTPRequest(map[string]string{
		"username_or_email": "test",
		"password":          "password",
	})
	transport.HandleRequest(endpoints, c, requests.MakeLoginRequest, endpoints.Service.Login)
	statusCode = c.Writer.Status()
	assert.Equal(t, http.StatusOK, statusCode)

	// Get my user
	c = makeTestContextWithHTTPRequest(map[string]string{})
	addValueAndParamToContext(c, "ID", 1, "user_id", "1")
	transport.HandleRequest(endpoints, c, requests.MakeGetUserRequest, endpoints.Service.GetUser)
	statusCode = c.Writer.Status()
	assert.Equal(t, http.StatusOK, statusCode)

	// Update my user
	c = makeTestContextWithHTTPRequest(map[string]string{
		"username": "test2",
		"email":    "test2@email.com",
	})
	addValueAndParamToContext(c, "ID", 1, "user_id", "1")
	transport.HandleRequest(endpoints, c, requests.MakeUpdateUserRequest, endpoints.Service.UpdateUser)
	statusCode = c.Writer.Status()
	assert.Equal(t, http.StatusOK, statusCode)

	// Login again
	c = makeTestContextWithHTTPRequest(map[string]string{
		"username_or_email": "test2",
		"password":          "password",
	})
	transport.HandleRequest(endpoints, c, requests.MakeLoginRequest, endpoints.Service.Login)
	statusCode = c.Writer.Status()
	assert.Equal(t, http.StatusOK, statusCode)

	// Delete my user
	c = makeTestContextWithHTTPRequest(map[string]string{})
	addValueAndParamToContext(c, "ID", 1, "user_id", "1")
	transport.HandleRequest(endpoints, c, requests.MakeDeleteUserRequest, endpoints.Service.DeleteUser)
	statusCode = c.Writer.Status()
	assert.Equal(t, http.StatusOK, statusCode)

	// Get my deleted user
	c = makeTestContextWithHTTPRequest(map[string]string{})
	addValueAndParamToContext(c, "ID", 1, "user_id", "1")
	transport.HandleRequest(endpoints, c, requests.MakeGetUserRequest, endpoints.Service.GetUser)
	statusCode = c.Writer.Status()
	assert.Equal(t, http.StatusNotFound, statusCode)
}

/* Helpers */

func makeTestHTTPRequest(body []byte) *http.Request {
	req, _ := http.NewRequest("", "", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	return req
}

func addRequestToContext(request *http.Request) *gin.Context {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = request
	return c
}

func makeTestContextWithHTTPRequest(body any) *gin.Context {
	jsonData, _ := json.Marshal(body)
	httpReq := makeTestHTTPRequest(jsonData)
	return addRequestToContext(httpReq)
}

func addValueAndParamToContext(context *gin.Context, ctxKey string, ctxValue any, paramKey, paramValue string) {
	context.Set(ctxKey, ctxValue)
	context.Params = append(context.Params, gin.Param{Key: paramKey, Value: paramValue})
}

/* Nops */

type nopLogger struct{}

func (l nopLogger) Info(args ...interface{})                  {}
func (l nopLogger) Warn(args ...interface{})                  {}
func (l nopLogger) Error(args ...interface{})                 {}
func (l nopLogger) Fatalf(format string, args ...interface{}) {}

type nopAuth struct{}

func (a nopAuth) GenerateToken(user auth.User, role auth.Role) (string, error)         { return "", nil }
func (a nopAuth) ValidateToken(role auth.Role, shouldMatchUserID bool) gin.HandlerFunc { return nil }
