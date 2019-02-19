package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignUpEndpoint(t *testing.T) {
	setupConfig()
	setupRouter()
	setupDatabase()
	deleteAllRecords()

	w := makeSignUpTestRequest()

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

func makeSignUpTestRequest() *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	body := `{
		"username": "gilperopiola",
		"email": "ferra.main@gmail.com",
		"password": "password"
	}`
	req, _ := http.NewRequest("POST", "/SignUp", bytes.NewReader([]byte(body)))
	router.ServeHTTP(w, req)
	return w
}

func TestLogInEndpoint(t *testing.T) {
	setupConfig()
	setupRouter()
	setupDatabase()
	deleteAllRecords()

	makeSignUpTestRequest()

	w := httptest.NewRecorder()
	body := `{
		"username": "gilperopiola",
		"password": "password"
	}`
	req, _ := http.NewRequest("POST", "/LogIn", bytes.NewReader([]byte(body)))
	router.ServeHTTP(w, req)

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
