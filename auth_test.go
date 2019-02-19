package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTesting() {
	setupConfig()
	setupRouter()
	setupDatabase()
	deleteAllRecords()
}

//SignUp

func makeSignUpTestRequest(username, email, password string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	body := `{
		"username": "` + username + `",
		"email": "` + email + `",
		"password": "` + password + `"
	}`
	req, _ := http.NewRequest("POST", "/SignUp", bytes.NewReader([]byte(body)))
	router.ServeHTTP(w, req)
	return w
}

func TestSignUpEndpoint(t *testing.T) {
	setupTesting()

	w := makeSignUpTestRequest("gilperopiola", "ferra.main@gmail.com", "password")
	var user User
	json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "gilperopiola", user.Username)
	assert.Equal(t, "ferra.main@gmail.com", user.Email)
	assert.False(t, user.Admin)
	assert.True(t, user.Active)
	assert.NotEmpty(t, user.DateCreated)
	assert.NotEmpty(t, user.Token)
}

func TestSignUpInvalid(t *testing.T) {
	setupTesting()

	w := makeSignUpTestRequest("", "", "")
	var user User
	json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, 400, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "all fields required"))

	w = makeSignUpTestRequest("gilperopiola", "ferra.main@gmail.com", "password")
	w = makeSignUpTestRequest("gilperopiola", "ferra.main@gmail.com", "password")
	json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, 400, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "username already in use"))

	w = makeSignUpTestRequest("gilperopiola2", "ferra.main@gmail.com", "password")
	json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, 400, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "email already in use"))
}

//LogIn

func makeLogInTestRequest(username, password string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	body := `{
		"username": "` + username + `",
		"password": "` + password + `"
	}`
	req, _ := http.NewRequest("POST", "/LogIn", bytes.NewReader([]byte(body)))
	router.ServeHTTP(w, req)
	return w
}

func TestLogInEndpoint(t *testing.T) {
	setupTesting()

	makeSignUpTestRequest("gilperopiola", "ferra.main@gmail.com", "password")
	w := makeLogInTestRequest("gilperopiola", "password")
	var user User
	json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "gilperopiola", user.Username)
	assert.Equal(t, "ferra.main@gmail.com", user.Email)
	assert.False(t, user.Admin)
	assert.True(t, user.Active)
	assert.NotEmpty(t, user.DateCreated)
	assert.NotEmpty(t, user.Token)
}

func TestLogInInvalid(t *testing.T) {
	setupTesting()

	makeSignUpTestRequest("gilperopiola", "ferra.main@gmail.com", "password")

	w := makeLogInTestRequest("", "")
	var user User
	json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, 400, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "both fields required"))

	w = makeLogInTestRequest("gilperopiola2", "password")
	json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, 400, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "wrong username"))

	w = makeLogInTestRequest("gilperopiola", "password2")
	json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, 400, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "wrong password"))
}
