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

func makeCreateUserTestRequest(username, email, password string, admin, active bool) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	body := `{
		"username": "` + username + `",
		"email": "` + email + `",
		"password": "` + password + `",
		"admin": ` + strconv.FormatBool(admin) + `,
		"active": ` + strconv.FormatBool(active) + `
	}`
	req, _ := http.NewRequest("POST", "/User", bytes.NewReader([]byte(body)))
	router.ServeHTTP(w, req)
	return w
}

func TestCreateUserEndpoint(t *testing.T) {
	setupTesting()

	w := makeCreateUserTestRequest("gilperopiola", "ferra.main@gmail.com", "password", true, true)
	var user User
	json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "gilperopiola", user.Username)
	assert.Equal(t, "ferra.main@gmail.com", user.Email)
	assert.True(t, user.Admin)
	assert.True(t, user.Active)
	assert.NotEmpty(t, user.DateCreated)
}

func TestCreateUserInvalid(t *testing.T) {
	setupTesting()

	w := makeCreateUserTestRequest("", "", "", false, false)
	var user User
	json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, 400, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "all fields required"))

	w = makeCreateUserTestRequest("gilperopiola", "ferra.main@gmail.com", "password", false, true)
	w = makeCreateUserTestRequest("gilperopiola", "ferra.main@gmail.com", "password", false, true)
	json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, 400, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "username already in use"))

	w = makeCreateUserTestRequest("gilperopiola2", "ferra.main@gmail.com", "password", false, true)
	json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, 400, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "email already in use"))
}

//ReadUser

func makeReadUserTestRequest(id int) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/User/"+strconv.Itoa(id), nil)
	router.ServeHTTP(w, req)
	return w
}

func TestReadUserEndpoint(t *testing.T) {
	setupTesting()

	w := makeCreateUserTestRequest("gilperopiola", "ferra.main@gmail.com", "password", true, true)
	var admin User
	json.Unmarshal(w.Body.Bytes(), &admin)

	w = makeReadUserTestRequest(int(admin.ID))
	var user User
	json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "gilperopiola", user.Username)
	assert.Equal(t, "ferra.main@gmail.com", user.Email)
	assert.True(t, user.Admin)
	assert.True(t, user.Active)
	assert.NotEmpty(t, user.DateCreated)
}

func TestReadUserInvalid(t *testing.T) {
	setupTesting()

	w := makeCreateUserTestRequest("gilperopiola", "ferra.main@gmail.com", "password", true, true)
	var admin User
	json.Unmarshal(w.Body.Bytes(), &admin)

	w = makeReadUserTestRequest(int(admin.ID + 99))
	assert.Equal(t, 400, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "record not found"))
}

//ReadUsers

func makeReadUsersTestRequest(id int, username, email string, limit, offset int, sortField, sortDir string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/Users?ID="+strconv.Itoa(id)+"&Username="+username+"&Email="+email+
		"&Limit="+strconv.Itoa(limit)+"&Offset="+strconv.Itoa(offset)+"&SortField="+sortField+"&SortDir="+sortDir, nil)
	router.ServeHTTP(w, req)
	return w
}

func TestReadUsersEndpoint(t *testing.T) {
	setupTesting()

	w := makeCreateUserTestRequest("gilperopiola", "ferra.main@gmail.com", "password", true, true)
	var admin User
	json.Unmarshal(w.Body.Bytes(), &admin)

	makeCreateUserTestRequest("gilperopiola2", "ferra.main2@gmail.com", "password", false, true)
	makeCreateUserTestRequest("franco2", "franco@hotmail.com", "password", false, true)
	makeCreateUserTestRequest("asdqwe", "qweasd@gmail.com", "password", false, true)

	w = makeReadUsersTestRequest(int(admin.ID), "", "", 100, 0, "", "")
	var users []User
	json.Unmarshal(w.Body.Bytes(), &users)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 1, len(users))
	assert.NotEmpty(t, users[0].ID)
	assert.Equal(t, "gilperopiola", users[0].Username)
	assert.Equal(t, "ferra.main@gmail.com", users[0].Email)

	w = makeReadUsersTestRequest(0, "gilpero", "", 1, 0, "ID", "DESC")
	json.Unmarshal(w.Body.Bytes(), &users)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 1, len(users))
	assert.NotEmpty(t, users[0].ID)
	assert.Equal(t, "gilperopiola2", users[0].Username)
	assert.Equal(t, "ferra.main2@gmail.com", users[0].Email)

	w = makeReadUsersTestRequest(0, "", "", 100, 2, "", "")
	json.Unmarshal(w.Body.Bytes(), &users)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 2, len(users))
}

//UpdateUser

func makeUpdateUserTestRequest(id int, username, email, password string, admin, active bool) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	body := `{
		"username": "` + username + `",
		"email": "` + email + `",
		"password": "` + password + `",
		"admin": ` + strconv.FormatBool(admin) + `,
		"active": ` + strconv.FormatBool(active) + `
	}`
	req, _ := http.NewRequest("PUT", "/User/"+strconv.Itoa(id), bytes.NewReader([]byte(body)))
	router.ServeHTTP(w, req)
	return w
}

func TestUpdateUserEndpoint(t *testing.T) {
	setupTesting()

	w := makeCreateUserTestRequest("gilperopiola", "ferra.main@gmail.com", "password", true, true)
	var admin User
	json.Unmarshal(w.Body.Bytes(), &admin)

	w = makeUpdateUserTestRequest(int(admin.ID), "gilperopiola2", "ferra.main2@gmail.com", "", false, false)
	var user User
	json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "gilperopiola2", user.Username)
	assert.Equal(t, "ferra.main2@gmail.com", user.Email)
	assert.False(t, user.Admin)
	assert.False(t, user.Active)
}

func TestUpdateUserInvalid(t *testing.T) {
	setupTesting()

	w := makeUpdateUserTestRequest(0, "", "", "", false, false)
	var user User
	json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, 400, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "record not found"))

	w = makeCreateUserTestRequest("gilperopiola", "ferra.main@gmail.com", "password", true, true)
	var admin User
	json.Unmarshal(w.Body.Bytes(), &admin)
	makeCreateUserTestRequest("gilperopiola2", "ferra.main2@gmail.com", "password", true, true)

	w = makeUpdateUserTestRequest(int(admin.ID), "gilperopiola2", "ferra.main2@gmail.com", "", true, true)
	json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, 400, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "username already in use"))

	w = makeUpdateUserTestRequest(int(admin.ID), "gilperopiola3", "ferra.main2@gmail.com", "", true, true)
	json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, 400, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "email already in use"))
}
