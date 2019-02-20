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

func TestCreateUserEndpoint(t *testing.T) {
	setupTesting()

	token := generateTestingAdminToken()

	w := makeCreateUserTestRequest(token, "gilperopiola", "ferra.main@gmail.com", "password", true, true)
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

	token := generateTestingAdminToken()

	w := makeCreateUserTestRequest(token, "", "", "", false, false)
	var user User
	json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, 400, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "all fields required"))

	w = makeCreateUserTestRequest(token, "gilperopiola", "ferra.main@gmail.com", "password", false, true)
	w = makeCreateUserTestRequest(token, "gilperopiola", "ferra.main@gmail.com", "password", false, true)
	json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, 400, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "username already in use"))

	w = makeCreateUserTestRequest(token, "gilperopiola2", "ferra.main@gmail.com", "password", false, true)
	json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, 400, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "email already in use"))
}

//ReadUser

func makeReadUserTestRequest(token string, id int) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/User/"+strconv.Itoa(id), nil)
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)
	return w
}

func TestReadUserEndpoint(t *testing.T) {
	setupTesting()

	token := generateTestingAdminToken()

	w := makeCreateUserTestRequest(token, "gilperopiola", "ferra.main@gmail.com", "password", true, true)
	var admin User
	json.Unmarshal(w.Body.Bytes(), &admin)

	w = makeReadUserTestRequest(token, int(admin.ID))
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

	token := generateTestingAdminToken()

	w := makeCreateUserTestRequest(token, "gilperopiola", "ferra.main@gmail.com", "password", true, true)
	var admin User
	json.Unmarshal(w.Body.Bytes(), &admin)

	w = makeReadUserTestRequest(token, int(admin.ID+99))
	assert.Equal(t, 400, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "record not found"))
}

//ReadUsers

func makeReadUsersTestRequest(token string, id int, username, email string, limit, offset int, sortField, sortDir string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/Users?ID="+strconv.Itoa(id)+"&Username="+username+"&Email="+email+
		"&Limit="+strconv.Itoa(limit)+"&Offset="+strconv.Itoa(offset)+"&SortField="+sortField+"&SortDir="+sortDir, nil)
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)
	return w
}

func TestReadUsersEndpoint(t *testing.T) {
	setupTesting()

	token := generateTestingAdminToken()

	w := makeCreateUserTestRequest(token, "gilperopiola", "ferra.main@gmail.com", "password", true, true)
	var admin User
	json.Unmarshal(w.Body.Bytes(), &admin)

	makeCreateUserTestRequest(token, "gilperopiola2", "ferra.main2@gmail.com", "password", false, true)
	makeCreateUserTestRequest(token, "franco2", "franco@hotmail.com", "password", false, true)
	makeCreateUserTestRequest(token, "asdqwe", "qweasd@gmail.com", "password", false, true)

	w = makeReadUsersTestRequest(token, int(admin.ID), "", "", 100, 0, "", "")
	var users []User
	json.Unmarshal(w.Body.Bytes(), &users)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 1, len(users))
	assert.NotEmpty(t, users[0].ID)
	assert.Equal(t, "gilperopiola", users[0].Username)
	assert.Equal(t, "ferra.main@gmail.com", users[0].Email)

	w = makeReadUsersTestRequest(token, 0, "gilpero", "", 1, 0, "ID", "DESC")
	json.Unmarshal(w.Body.Bytes(), &users)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 1, len(users))
	assert.NotEmpty(t, users[0].ID)
	assert.Equal(t, "gilperopiola2", users[0].Username)
	assert.Equal(t, "ferra.main2@gmail.com", users[0].Email)

	w = makeReadUsersTestRequest(token, 0, "", "", 100, 2, "", "")
	json.Unmarshal(w.Body.Bytes(), &users)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 2, len(users))
}

//UpdateUser

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

func TestUpdateUserEndpoint(t *testing.T) {
	setupTesting()

	token := generateTestingAdminToken()

	w := makeCreateUserTestRequest(token, "gilperopiola", "ferra.main@gmail.com", "password", true, true)
	var admin User
	json.Unmarshal(w.Body.Bytes(), &admin)

	w = makeUpdateUserTestRequest(token, int(admin.ID), "gilperopiola2", "ferra.main2@gmail.com", "", false, false)
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

	token := generateTestingAdminToken()

	w := makeUpdateUserTestRequest(token, 0, "", "", "", false, false)
	var user User
	json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, 400, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "record not found"))

	w = makeCreateUserTestRequest(token, "gilperopiola", "ferra.main@gmail.com", "password", true, true)
	var admin User
	json.Unmarshal(w.Body.Bytes(), &admin)
	makeCreateUserTestRequest(token, "gilperopiola2", "ferra.main2@gmail.com", "password", true, true)

	w = makeUpdateUserTestRequest(token, int(admin.ID), "gilperopiola2", "ferra.main2@gmail.com", "", true, true)
	json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, 400, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "username already in use"))

	w = makeUpdateUserTestRequest(token, int(admin.ID), "gilperopiola3", "ferra.main2@gmail.com", "", true, true)
	json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, 400, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "email already in use"))
}
