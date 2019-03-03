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

//CreateActor

func TestCreateActorEndpoint(t *testing.T) {
	_, adminToken := setupTesting()

	var actor Actor
	success := makeCreateActorTestRequest(adminToken, "Emma Stone", true)
	json.Unmarshal(success.Body.Bytes(), &actor)
	assert.Equal(t, http.StatusOK, success.Code)
	assert.NotEmpty(t, actor.ID)
	assert.Equal(t, "Emma Stone", actor.Name)
	assert.True(t, actor.Active)
	assert.NotEmpty(t, actor.DateCreated)
}

func TestCreateActorInvalid(t *testing.T) {
	_, adminToken := setupTesting()

	makeCreateActorTestRequest(adminToken, "Emma Stone", true)
	empty := makeCreateActorTestRequest(adminToken, "", false)

	var actor Actor
	json.Unmarshal(empty.Body.Bytes(), &actor)
	assert.Equal(t, http.StatusBadRequest, empty.Code)
	assert.True(t, strings.Contains(empty.Body.String(), "name required"))
}

//ReadActor

func TestReadActorEndpoint(t *testing.T) {
	_, adminToken := setupTesting()

	var actor Actor
	success := makeCreateActorTestRequest(adminToken, "Emma Stone", true)
	json.Unmarshal(success.Body.Bytes(), &actor)

	success = makeReadActorTestRequest(adminToken, int(actor.ID))
	json.Unmarshal(success.Body.Bytes(), &actor)
	assert.Equal(t, http.StatusOK, success.Code)
	assert.NotEmpty(t, actor.ID)
	assert.Equal(t, "Emma Stone", actor.Name)
	assert.True(t, actor.Active)
	assert.NotEmpty(t, actor.DateCreated)
}

func TestReadActorInvalid(t *testing.T) {
	_, adminToken := setupTesting()

	notFound := makeReadActorTestRequest(adminToken, 1)
	assert.Equal(t, http.StatusBadRequest, notFound.Code)
	assert.True(t, strings.Contains(notFound.Body.String(), "record not found"))
}

//ReadActors

func TestReadActorsEndpoint(t *testing.T) {
	_, adminToken := setupTesting()

	makeCreateActorTestRequest(adminToken, "Emma Stone", true)
	makeCreateActorTestRequest(adminToken, "Emma Stone 2", true)
	makeCreateActorTestRequest(adminToken, "Elle Fanning", true)
	makeCreateActorTestRequest(adminToken, "Amanda Seyfried", true)

	var actors []Actor
	success := makeReadActorsTestRequest(adminToken, 0, "Emma S", 1, 0, "ID", "DESC")
	json.Unmarshal(success.Body.Bytes(), &actors)
	assert.Equal(t, http.StatusOK, success.Code)
	assert.Equal(t, 1, len(actors))
	assert.NotEmpty(t, actors[0].ID)
	assert.Equal(t, "Emma Stone 2", actors[0].Name)

	success = makeReadActorsTestRequest(adminToken, int(actors[0].ID), "", 100, 0, "", "")
	json.Unmarshal(success.Body.Bytes(), &actors)
	assert.Equal(t, http.StatusOK, success.Code)
	assert.Equal(t, 1, len(actors))
	assert.NotEmpty(t, actors[0].ID)
	assert.Equal(t, "Emma Stone 2", actors[0].Name)

	success = makeReadActorsTestRequest(adminToken, 0, "", 2, 2, "Name", "DESC")
	json.Unmarshal(success.Body.Bytes(), &actors)
	assert.Equal(t, http.StatusOK, success.Code)
	assert.Equal(t, 2, len(actors))
	assert.Equal(t, "Elle Fanning", actors[0].Name)
	assert.Equal(t, "Amanda Seyfried", actors[1].Name)
}

//UpdateActor

func TestUpdateActorEndpoint(t *testing.T) {
	_, adminToken := setupTesting()

	var actor Actor
	success := makeCreateActorTestRequest(adminToken, "Emma Stone", true)
	json.Unmarshal(success.Body.Bytes(), &actor)

	success = makeUpdateActorTestRequest(adminToken, int(actor.ID), "Emma Stone 2", false)
	json.Unmarshal(success.Body.Bytes(), &actor)
	assert.Equal(t, http.StatusOK, success.Code)
	assert.NotEmpty(t, actor.ID)
	assert.Equal(t, "Emma Stone 2", actor.Name)
	assert.False(t, actor.Active)
}

func TestUpdateActorInvalid(t *testing.T) {
	_, adminToken := setupTesting()

	var actor Actor
	notFound := makeUpdateActorTestRequest(adminToken, 0, "", false)
	json.Unmarshal(notFound.Body.Bytes(), &actor)
	assert.Equal(t, http.StatusBadRequest, notFound.Code)
	assert.True(t, strings.Contains(notFound.Body.String(), "record not found"))

	success := makeCreateActorTestRequest(adminToken, "Emma Stone", true)
	makeCreateActorTestRequest(adminToken, "Emma Stone 2", true)
	json.Unmarshal(success.Body.Bytes(), &actor)
}

//Helpers

func makeCreateActorTestRequest(token, name string, active bool) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	body := `{
		"name": "` + name + `",
		"active": ` + strconv.FormatBool(active) + `
	}`
	req, _ := http.NewRequest("POST", "/Admin/Actor", bytes.NewReader([]byte(body)))
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)
	return w
}

func makeReadActorTestRequest(token string, id int) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/Admin/Actor/"+strconv.Itoa(id), nil)
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)
	return w
}

func makeReadActorsTestRequest(token string, id int, name string, limit, offset int, sortField, sortDir string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/Admin/Actors?ID="+strconv.Itoa(id)+"&Name="+name+
		"&Limit="+strconv.Itoa(limit)+"&Offset="+strconv.Itoa(offset)+"&SortField="+sortField+"&SortDir="+sortDir, nil)
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)
	return w
}

func makeUpdateActorTestRequest(token string, id int, name string, active bool) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	body := `{
		"name": "` + name + `",
		"active": ` + strconv.FormatBool(active) + `
	}`
	req, _ := http.NewRequest("PUT", "/Admin/Actor/"+strconv.Itoa(id), bytes.NewReader([]byte(body)))
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)
	return w
}
