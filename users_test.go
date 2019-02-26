package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//CreateUser

func TestCreateUserEndpoint(t *testing.T) {
	_, adminToken := setupTesting()

	var user User
	success := makeCreateUserTestRequest(adminToken, "gilperopiola", "ferra.main@gmail.com", "password", true, true)
	json.Unmarshal(success.Body.Bytes(), &user)
	assert.Equal(t, http.StatusOK, success.Code)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "gilperopiola", user.Username)
	assert.Equal(t, "ferra.main@gmail.com", user.Email)
	assert.True(t, user.Admin)
	assert.True(t, user.Active)
	assert.NotEmpty(t, user.DateCreated)
}

func TestCreateUserInvalid(t *testing.T) {
	_, adminToken := setupTesting()

	makeCreateUserTestRequest(adminToken, "gilperopiola", "ferra.main@gmail.com", "password", false, true)
	empty := makeCreateUserTestRequest(adminToken, "", "", "", false, false)
	duplicate := makeCreateUserTestRequest(adminToken, "gilperopiola", "ferra.main@gmail.com", "password", false, true)
	duplicate2 := makeCreateUserTestRequest(adminToken, "gilperopiola2", "ferra.main@gmail.com", "password", false, true)

	var user User
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

//ReadUser

func TestReadUserEndpoint(t *testing.T) {
	_, adminToken := setupTesting()

	var user User
	success := makeCreateUserTestRequest(adminToken, "gilperopiola", "ferra.main@gmail.com", "password", false, true)
	json.Unmarshal(success.Body.Bytes(), &user)

	success = makeReadUserTestRequest(adminToken, int(user.ID))
	json.Unmarshal(success.Body.Bytes(), &user)
	assert.Equal(t, http.StatusOK, success.Code)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "gilperopiola", user.Username)
	assert.Equal(t, "ferra.main@gmail.com", user.Email)
	assert.False(t, user.Admin)
	assert.True(t, user.Active)
	assert.NotEmpty(t, user.DateCreated)
}

func TestReadUserInvalid(t *testing.T) {
	_, adminToken := setupTesting()

	notFound := makeReadUserTestRequest(adminToken, 1)
	assert.Equal(t, http.StatusBadRequest, notFound.Code)
	assert.True(t, strings.Contains(notFound.Body.String(), "record not found"))
}

//ReadUsers

func TestReadUsersEndpoint(t *testing.T) {
	_, adminToken := setupTesting()

	makeCreateUserTestRequest(adminToken, "gilperopiola", "ferra.main@gmail.com", "password", false, true)
	makeCreateUserTestRequest(adminToken, "gilperopiola2", "ferra.main2@gmail.com", "password", false, true)
	makeCreateUserTestRequest(adminToken, "franco2", "franco@hotmail.com", "password", false, true)
	makeCreateUserTestRequest(adminToken, "asdqwe", "qweasd@gmail.com", "password", false, true)

	var users []User
	success := makeReadUsersTestRequest(adminToken, 0, "gilpero", "", 1, 0, "ID", "DESC")
	json.Unmarshal(success.Body.Bytes(), &users)
	assert.Equal(t, http.StatusOK, success.Code)
	assert.Equal(t, 1, len(users))
	assert.NotEmpty(t, users[0].ID)
	assert.Equal(t, "gilperopiola2", users[0].Username)
	assert.Equal(t, "ferra.main2@gmail.com", users[0].Email)

	success = makeReadUsersTestRequest(adminToken, int(users[0].ID), "", "", 100, 0, "", "")
	json.Unmarshal(success.Body.Bytes(), &users)
	assert.Equal(t, http.StatusOK, success.Code)
	assert.Equal(t, 1, len(users))
	assert.NotEmpty(t, users[0].ID)
	assert.Equal(t, "gilperopiola2", users[0].Username)
	assert.Equal(t, "ferra.main2@gmail.com", users[0].Email)

	success = makeReadUsersTestRequest(adminToken, 0, "", "", 100, 2, "", "")
	json.Unmarshal(success.Body.Bytes(), &users)
	assert.Equal(t, http.StatusOK, success.Code)
	assert.Equal(t, 2, len(users))
}

//UpdateUser

func TestUpdateUserEndpoint(t *testing.T) {
	_, adminToken := setupTesting()

	var user User
	success := makeCreateUserTestRequest(adminToken, "gilperopiola", "ferra.main@gmail.com", "password", false, true)
	json.Unmarshal(success.Body.Bytes(), &user)

	success = makeUpdateUserTestRequest(adminToken, int(user.ID), "gilperopiola2", "ferra.main2@gmail.com", "", false, false)
	json.Unmarshal(success.Body.Bytes(), &user)
	assert.Equal(t, http.StatusOK, success.Code)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "gilperopiola2", user.Username)
	assert.Equal(t, "ferra.main2@gmail.com", user.Email)
	assert.False(t, user.Admin)
	assert.False(t, user.Active)
}

func TestUpdateUserInvalid(t *testing.T) {
	_, adminToken := setupTesting()

	var user User
	notFound := makeUpdateUserTestRequest(adminToken, 0, "", "", "", false, false)
	json.Unmarshal(notFound.Body.Bytes(), &user)
	assert.Equal(t, http.StatusBadRequest, notFound.Code)
	assert.True(t, strings.Contains(notFound.Body.String(), "record not found"))

	success := makeCreateUserTestRequest(adminToken, "gilperopiola", "ferra.main@gmail.com", "password", false, true)
	makeCreateUserTestRequest(adminToken, "gilperopiola2", "ferra.main2@gmail.com", "password", true, true)
	json.Unmarshal(success.Body.Bytes(), &user)

	duplicate := makeUpdateUserTestRequest(adminToken, int(user.ID), "gilperopiola2", "ferra.main2@gmail.com", "", true, true)
	json.Unmarshal(duplicate.Body.Bytes(), &user)
	assert.Equal(t, http.StatusBadRequest, duplicate.Code)
	assert.True(t, strings.Contains(duplicate.Body.String(), "username already in use"))

	duplicate2 := makeUpdateUserTestRequest(adminToken, int(user.ID), "gilperopiola3", "ferra.main2@gmail.com", "", true, true)
	json.Unmarshal(duplicate2.Body.Bytes(), &user)
	assert.Equal(t, http.StatusBadRequest, duplicate2.Code)
	assert.True(t, strings.Contains(duplicate2.Body.String(), "email already in use"))
}

//Helpers

func makeCreateUserTestRequest(token, username, email, password string, admin, active bool) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	body := `{
		"username": "` + username + `",
		"email": "` + email + `",
		"password": "` + password + `",
		"admin": ` + strconv.FormatBool(admin) + `,
		"active": ` + strconv.FormatBool(active) + `
	}`
	req, _ := http.NewRequest("POST", "/User", bytes.NewReader([]byte(body)))
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)
	return w
}

func makeReadUserTestRequest(token string, id int) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/User/"+strconv.Itoa(id), nil)
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)
	return w
}

func makeReadUsersTestRequest(token string, id int, username, email string, limit, offset int, sortField, sortDir string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/Users?ID="+strconv.Itoa(id)+"&Username="+username+"&Email="+email+
		"&Limit="+strconv.Itoa(limit)+"&Offset="+strconv.Itoa(offset)+"&SortField="+sortField+"&SortDir="+sortDir, nil)
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)
	return w
}

func makeUpdateUserTestRequest(token string, id int, username, email, password string, admin, active bool) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	body := `{
		"username": "` + username + `",
		"email": "` + email + `",
		"password": "` + password + `",
		"admin": ` + strconv.FormatBool(admin) + `,
		"active": ` + strconv.FormatBool(active) + `
	}`
	req, _ := http.NewRequest("PUT", "/User/"+strconv.Itoa(id), bytes.NewReader([]byte(body)))
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)
	return w
}
