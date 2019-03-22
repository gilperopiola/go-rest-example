package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

//CreateUser
func TestCreateUser(t *testing.T) {
	_, adminToken := setupTesting()

	userInput := User{Username: "username", Email: "email", Password: "password", Admin: true, Active: true}
	userOutput := User{}

	response := makeUserTestRequest(adminToken, "POST", "/Admin/User", &userInput)
	json.Unmarshal(response.Body.Bytes(), &userOutput)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.Equal(t, userInput.Username, userOutput.Username)
	assert.Equal(t, userInput.Email, userOutput.Email)
	assert.Equal(t, userInput.Admin, userOutput.Admin)
	assert.Equal(t, userInput.Active, userOutput.Active)
}

func TestCreateUserEmpty(t *testing.T) {
	_, adminToken := setupTesting()

	userInput := User{}

	response := makeUserTestRequest(adminToken, "POST", "/Admin/User", &userInput)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

//ReadUser
func TestReadUser(t *testing.T) {
	_, adminToken := setupTesting()

	user := database.GetTestingUsers()[0]

	userInput := User{}
	userOutput := User{}

	response := makeUserTestRequest(adminToken, "GET", "/Admin/User/"+strconv.Itoa(int(user.ID)), &userInput)
	json.Unmarshal(response.Body.Bytes(), &userOutput)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.Equal(t, user.Username, userOutput.Username)
	assert.Equal(t, user.Email, userOutput.Email)
	assert.Equal(t, user.Admin, userOutput.Admin)
	assert.Equal(t, user.Active, userOutput.Active)
}

func TestReadUserInvalid(t *testing.T) {
	_, adminToken := setupTesting()

	userInput := User{}

	response := makeUserTestRequest(adminToken, "GET", "/Admin/User/0", &userInput)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

//ReadUsers
func TestReadUsers(t *testing.T) {
	_, adminToken := setupTesting()

	users := database.GetTestingUsers()
	userInput := User{Username: "name"}
	usersOutput := []User{}

	response := makeReadUsersTestRequest(adminToken, &userInput, 2, 1, "email", "DESC")
	json.Unmarshal(response.Body.Bytes(), &usersOutput)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.Equal(t, 2, len(usersOutput))
	assert.Equal(t, users[0].Username, usersOutput[1].Username)
	assert.Equal(t, users[1].Username, usersOutput[0].Username)
}

//UpdateUser
func TestUpdateUserEndpoint(t *testing.T) {
	_, adminToken := setupTesting()

	user := database.GetTestingUsers()[0]
	userInput := User{Username: "username", Email: "email", Password: "password", Admin: false, Active: false}
	userOutput := User{}

	response := makeUserTestRequest(adminToken, "PUT", "/Admin/User/"+strconv.Itoa(int(user.ID)), &userInput)
	json.Unmarshal(response.Body.Bytes(), &userOutput)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.Equal(t, userInput.Username, userOutput.Username)
	assert.Equal(t, userInput.Email, userOutput.Email)
	assert.Equal(t, userInput.Admin, userOutput.Admin)
	assert.Equal(t, userInput.Active, userOutput.Active)
}

func TestUpdateUserInvalid(t *testing.T) {
	_, adminToken := setupTesting()

	users := database.GetTestingUsers()
	userInput := User{ID: users[2].ID + 1}

	response := makeUserTestRequest(adminToken, "PUT", "/Admin/User/"+strconv.Itoa(int(userInput.ID)), &userInput)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	userInput = User{ID: users[0].ID, Username: users[1].Username}
	response = makeUserTestRequest(adminToken, "PUT", "/Admin/User/"+strconv.Itoa(int(userInput.ID)), &userInput)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

//Helpers
func makeUserTestRequest(token, method, url string, user *User) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	body := user.GetJSONBody()
	req, _ := http.NewRequest(method, url, bytes.NewReader([]byte(body)))
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)
	return w
}

func makeReadUsersTestRequest(token string, user *User, limit, offset int, sortField, sortDir string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()

	url := strconv.Itoa(int(user.ID)) + "&Username=" + user.Username + "&Limit=" + strconv.Itoa(limit) + "&Offset=" + strconv.Itoa(offset) +
		"&SortField=" + sortField + "&SortDir=" + sortDir

	req, _ := http.NewRequest("GET", "/Admin/Users?ID="+url, nil)
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)
	return w
}
