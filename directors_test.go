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

//CreateDirector

func TestCreateDirectorEndpoint(t *testing.T) {
	_, adminToken := setupTesting()

	var director Director
	success := makeCreateDirectorTestRequest(adminToken, "Wes Anderson", true)
	json.Unmarshal(success.Body.Bytes(), &director)
	assert.Equal(t, http.StatusOK, success.Code)
	assert.NotEmpty(t, director.ID)
	assert.Equal(t, "Wes Anderson", director.Name)
	assert.True(t, director.Active)
	assert.NotEmpty(t, director.DateCreated)
}

func TestCreateDirectorInvalid(t *testing.T) {
	_, adminToken := setupTesting()

	makeCreateDirectorTestRequest(adminToken, "Wes Anderson", true)
	empty := makeCreateDirectorTestRequest(adminToken, "", false)
	duplicate := makeCreateDirectorTestRequest(adminToken, "Wes Anderson", true)

	var director Director
	json.Unmarshal(empty.Body.Bytes(), &director)
	assert.Equal(t, http.StatusBadRequest, empty.Code)
	assert.True(t, strings.Contains(empty.Body.String(), "name required"))

	json.Unmarshal(duplicate.Body.Bytes(), &director)
	assert.Equal(t, http.StatusBadRequest, duplicate.Code)
	assert.True(t, strings.Contains(duplicate.Body.String(), "name already in use"))
}

//ReadDirector

func TestReadDirectorEndpoint(t *testing.T) {
	_, adminToken := setupTesting()

	var director Director
	success := makeCreateDirectorTestRequest(adminToken, "Wes Anderson", true)
	json.Unmarshal(success.Body.Bytes(), &director)

	success = makeReadDirectorTestRequest(adminToken, int(director.ID))
	json.Unmarshal(success.Body.Bytes(), &director)
	assert.Equal(t, http.StatusOK, success.Code)
	assert.NotEmpty(t, director.ID)
	assert.Equal(t, "Wes Anderson", director.Name)
	assert.True(t, director.Active)
	assert.NotEmpty(t, director.DateCreated)
}

func TestReadDirectorInvalid(t *testing.T) {
	_, adminToken := setupTesting()

	notFound := makeReadDirectorTestRequest(adminToken, 1)
	assert.Equal(t, http.StatusBadRequest, notFound.Code)
	assert.True(t, strings.Contains(notFound.Body.String(), "record not found"))
}

//ReadDirectors

func TestReadDirectorsEndpoint(t *testing.T) {
	_, adminToken := setupTesting()

	makeCreateDirectorTestRequest(adminToken, "Wes Anderson", true)
	makeCreateDirectorTestRequest(adminToken, "Wes Anderson 2", true)
	makeCreateDirectorTestRequest(adminToken, "Steven Spielberg", true)
	makeCreateDirectorTestRequest(adminToken, "Guillermo Del Toro", true)

	var directors []Director
	success := makeReadDirectorsTestRequest(adminToken, 0, "Wes A", 1, 0, "ID", "DESC")
	json.Unmarshal(success.Body.Bytes(), &directors)
	assert.Equal(t, http.StatusOK, success.Code)
	assert.Equal(t, 1, len(directors))
	assert.NotEmpty(t, directors[0].ID)
	assert.Equal(t, "Wes Anderson 2", directors[0].Name)

	success = makeReadDirectorsTestRequest(adminToken, int(directors[0].ID), "", 100, 0, "", "")
	json.Unmarshal(success.Body.Bytes(), &directors)
	assert.Equal(t, http.StatusOK, success.Code)
	assert.Equal(t, 1, len(directors))
	assert.NotEmpty(t, directors[0].ID)
	assert.Equal(t, "Wes Anderson 2", directors[0].Name)

	success = makeReadDirectorsTestRequest(adminToken, 0, "", 2, 2, "Name", "DESC")
	json.Unmarshal(success.Body.Bytes(), &directors)
	assert.Equal(t, http.StatusOK, success.Code)
	assert.Equal(t, 2, len(directors))
	assert.Equal(t, "Steven Spielberg", directors[0].Name)
	assert.Equal(t, "Guillermo Del Toro", directors[1].Name)
}

//UpdateDirector

func TestUpdateDirectorEndpoint(t *testing.T) {
	_, adminToken := setupTesting()

	var director Director
	success := makeCreateDirectorTestRequest(adminToken, "Wes Anderson", true)
	json.Unmarshal(success.Body.Bytes(), &director)

	success = makeUpdateDirectorTestRequest(adminToken, int(director.ID), "Wes Anderson 2", false)
	json.Unmarshal(success.Body.Bytes(), &director)
	assert.Equal(t, http.StatusOK, success.Code)
	assert.NotEmpty(t, director.ID)
	assert.Equal(t, "Wes Anderson 2", director.Name)
	assert.False(t, director.Active)
}

func TestUpdateDirectorInvalid(t *testing.T) {
	_, adminToken := setupTesting()

	var director Director
	notFound := makeUpdateDirectorTestRequest(adminToken, 0, "", false)
	json.Unmarshal(notFound.Body.Bytes(), &director)
	assert.Equal(t, http.StatusBadRequest, notFound.Code)
	assert.True(t, strings.Contains(notFound.Body.String(), "record not found"))

	success := makeCreateDirectorTestRequest(adminToken, "Wes Anderson", true)
	makeCreateDirectorTestRequest(adminToken, "Wes Anderson 2", true)
	json.Unmarshal(success.Body.Bytes(), &director)

	duplicate := makeUpdateDirectorTestRequest(adminToken, int(director.ID), "Wes Anderson 2", true)
	json.Unmarshal(duplicate.Body.Bytes(), &director)
	assert.Equal(t, http.StatusBadRequest, duplicate.Code)
	assert.True(t, strings.Contains(duplicate.Body.String(), "name already in use"))
}

//Helpers

func makeCreateDirectorTestRequest(token, name string, active bool) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	body := `{
		"name": "` + name + `",
		"active": ` + strconv.FormatBool(active) + `
	}`
	req, _ := http.NewRequest("POST", "/Admin/Director", bytes.NewReader([]byte(body)))
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)
	return w
}

func makeReadDirectorTestRequest(token string, id int) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/Admin/Director/"+strconv.Itoa(id), nil)
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)
	return w
}

func makeReadDirectorsTestRequest(token string, id int, name string, limit, offset int, sortField, sortDir string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/Admin/Directors?ID="+strconv.Itoa(id)+"&Name="+name+
		"&Limit="+strconv.Itoa(limit)+"&Offset="+strconv.Itoa(offset)+"&SortField="+sortField+"&SortDir="+sortDir, nil)
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)
	return w
}

func makeUpdateDirectorTestRequest(token string, id int, name string, active bool) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	body := `{
		"name": "` + name + `",
		"active": ` + strconv.FormatBool(active) + `
	}`
	req, _ := http.NewRequest("PUT", "/Admin/Director/"+strconv.Itoa(id), bytes.NewReader([]byte(body)))
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)
	return w
}
