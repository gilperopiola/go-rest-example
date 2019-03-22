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

//CreateDirector
func TestCreateDirector(t *testing.T) {
	_, adminToken := setupTesting()

	directorInput := Director{Name: "name", Active: true}
	directorOutput := Director{}

	response := makeDirectorTestRequest(adminToken, "POST", "/Admin/Director", &directorInput)
	json.Unmarshal(response.Body.Bytes(), &directorOutput)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.Equal(t, directorInput.Name, directorOutput.Name)
	assert.Equal(t, directorInput.Active, directorOutput.Active)
}

func TestCreateDirectorEmpty(t *testing.T) {
	_, adminToken := setupTesting()

	directorInput := Director{}

	response := makeDirectorTestRequest(adminToken, "POST", "/Admin/Director", &directorInput)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

//ReadDirector
func TestReadDirector(t *testing.T) {
	_, adminToken := setupTesting()

	director := database.GetTestingDirectors()[0]

	directorInput := Director{}
	directorOutput := Director{}

	response := makeDirectorTestRequest(adminToken, "GET", "/Admin/Director/"+strconv.Itoa(int(director.ID)), &directorInput)
	json.Unmarshal(response.Body.Bytes(), &directorOutput)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.Equal(t, director.Name, directorOutput.Name)
	assert.Equal(t, director.Active, directorOutput.Active)
}

func TestReadDirectorInvalid(t *testing.T) {
	_, adminToken := setupTesting()

	directorInput := Director{}

	response := makeDirectorTestRequest(adminToken, "GET", "/Admin/Director/0", &directorInput)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

//ReadDirectors
func TestReadDirectors(t *testing.T) {
	_, adminToken := setupTesting()

	directors := database.GetTestingDirectors()
	directorInput := Director{Name: "name"}
	directorsOutput := []Director{}

	response := makeReadDirectorsTestRequest(adminToken, &directorInput, 2, 1, "name", "DESC")
	json.Unmarshal(response.Body.Bytes(), &directorsOutput)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.Equal(t, 2, len(directorsOutput))
	assert.Equal(t, directors[0].Name, directorsOutput[1].Name)
	assert.Equal(t, directors[1].Name, directorsOutput[0].Name)
}

//UpdateDirector
func TestUpdateDirectorEndpoint(t *testing.T) {
	_, adminToken := setupTesting()

	director := database.GetTestingDirectors()[0]
	directorInput := Director{Name: "name", Active: false}
	directorOutput := Director{}

	response := makeDirectorTestRequest(adminToken, "PUT", "/Admin/Director/"+strconv.Itoa(int(director.ID)), &directorInput)
	json.Unmarshal(response.Body.Bytes(), &directorOutput)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.Equal(t, directorInput.Name, directorOutput.Name)
	assert.Equal(t, directorInput.Active, directorOutput.Active)
}

func TestUpdateDirectorInvalid(t *testing.T) {
	_, adminToken := setupTesting()

	directors := database.GetTestingDirectors()
	directorInput := Director{ID: directors[2].ID + 1}

	response := makeDirectorTestRequest(adminToken, "PUT", "/Admin/Director/"+strconv.Itoa(int(directorInput.ID)), &directorInput)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	directorInput = Director{ID: directors[0].ID, Name: directors[1].Name}
	response = makeDirectorTestRequest(adminToken, "PUT", "/Admin/Director/"+strconv.Itoa(int(directorInput.ID)), &directorInput)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

//Helpers
func makeDirectorTestRequest(token, method, url string, director *Director) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	body := director.GetJSONBody()
	req, _ := http.NewRequest(method, url, bytes.NewReader([]byte(body)))
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)
	return w
}

func makeReadDirectorsTestRequest(token string, director *Director, limit, offset int, sortField, sortDir string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()

	url := strconv.Itoa(int(director.ID)) + "&Name=" + director.Name + "&Limit=" + strconv.Itoa(limit) + "&Offset=" + strconv.Itoa(offset) +
		"&SortField=" + sortField + "&SortDir=" + sortDir

	req, _ := http.NewRequest("GET", "/Admin/Directors?ID="+url, nil)
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)
	return w
}
