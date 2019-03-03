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

func setupTesting() (string, string) {
	setupConfig()
	setupRouter()
	setupDatabase()
	purgeDatabase()

	return generateTestingToken("User"), generateTestingToken("Admin")
}

//SignUp

func TestSignUpEndpoint(t *testing.T) {
	setupTesting()

	var user User
	success := makeSignUpTestRequest("gilperopiola", "ferra.main@gmail.com", "password")
	json.Unmarshal(success.Body.Bytes(), &user)
	assert.Equal(t, http.StatusOK, success.Code)
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

	var user User
	empty := makeSignUpTestRequest("", "", "")
	makeSignUpTestRequest("gilperopiola", "ferra.main@gmail.com", "password")
	duplicate := makeSignUpTestRequest("gilperopiola", "ferra.main@gmail.com", "password")
	duplicate2 := makeSignUpTestRequest("gilperopiola2", "ferra.main@gmail.com", "password")

	json.Unmarshal(empty.Body.Bytes(), &user)
	assert.Equal(t, http.StatusBadRequest, empty.Code)
	assert.True(t, strings.Contains(empty.Body.String(), "all fields required"))

	json.Unmarshal(duplicate.Body.Bytes(), &user)
	assert.Equal(t, http.StatusBadRequest, duplicate.Code)
	assert.True(t, strings.Contains(duplicate.Body.String(), "username already in use"))

	json.Unmarshal(duplicate2.Body.Bytes(), &user)
	assert.Equal(t, http.StatusBadRequest, duplicate2.Code)
	assert.True(t, strings.Contains(duplicate2.Body.String(), "email already in use"))
}

//LogIn

func TestLogInEndpoint(t *testing.T) {
	setupTesting()

	var user User
	makeSignUpTestRequest("gilperopiola", "ferra.main@gmail.com", "password")
	success := makeLogInTestRequest("gilperopiola", "password")

	json.Unmarshal(success.Body.Bytes(), &user)
	assert.Equal(t, http.StatusOK, success.Code)
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

	var user User
	makeSignUpTestRequest("gilperopiola", "ferra.main@gmail.com", "password")
	empty := makeLogInTestRequest("", "")
	wrong := makeLogInTestRequest("gilperopiola2", "password")
	wrong2 := makeLogInTestRequest("gilperopiola", "password2")

	json.Unmarshal(empty.Body.Bytes(), &user)
	assert.Equal(t, http.StatusBadRequest, empty.Code)
	assert.True(t, strings.Contains(empty.Body.String(), "both fields required"))

	json.Unmarshal(wrong.Body.Bytes(), &user)
	assert.Equal(t, http.StatusBadRequest, wrong.Code)
	assert.True(t, strings.Contains(wrong.Body.String(), "wrong username"))

	json.Unmarshal(wrong2.Body.Bytes(), &user)
	assert.Equal(t, http.StatusBadRequest, wrong2.Code)
	assert.True(t, strings.Contains(wrong2.Body.String(), "wrong password"))
}

//Helpers

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
