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

//CreateActor
func TestCreateActor(t *testing.T) {
	_, adminToken := setupTesting()

	actorInput := Actor{Name: "name", Active: true}
	actorOutput := Actor{}

	response := makeActorTestRequest(adminToken, "POST", "/Admin/Actor", &actorInput)
	json.Unmarshal(response.Body.Bytes(), &actorOutput)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.Equal(t, actorInput.Name, actorOutput.Name)
	assert.Equal(t, actorInput.Active, actorOutput.Active)

	//

	actorInput = Actor{}
	response = makeActorTestRequest(adminToken, "POST", "/Admin/Actor", &actorInput)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

//ReadActor
func TestReadActor(t *testing.T) {
	_, adminToken := setupTesting()

	actor := database.GetTestingActors()[0]

	actorInput := Actor{}
	actorOutput := Actor{}

	response := makeActorTestRequest(adminToken, "GET", "/Admin/Actor/"+strconv.Itoa(int(actor.ID)), &actorInput)
	json.Unmarshal(response.Body.Bytes(), &actorOutput)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.Equal(t, actor.Name, actorOutput.Name)
	assert.Equal(t, actor.Active, actorOutput.Active)

	//

	actorInput = Actor{}
	response = makeActorTestRequest(adminToken, "GET", "/Admin/Actor/0", &actorInput)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

//ReadActors
func TestReadActors(t *testing.T) {
	_, adminToken := setupTesting()

	actors := database.GetTestingActors()
	actorInput := Actor{Name: "name"}
	actorsOutput := []Actor{}

	response := makeReadActorsTestRequest(adminToken, &actorInput, 2, 1, "name", "DESC")
	json.Unmarshal(response.Body.Bytes(), &actorsOutput)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.Equal(t, 2, len(actorsOutput))
	assert.Equal(t, actors[0].Name, actorsOutput[1].Name)
	assert.Equal(t, actors[1].Name, actorsOutput[0].Name)
}

//UpdateActor
func TestUpdateActorEndpoint(t *testing.T) {
	_, adminToken := setupTesting()

	actor := database.GetTestingActors()[0]
	actorInput := Actor{Name: "name", Active: false}
	actorOutput := Actor{}

	response := makeActorTestRequest(adminToken, "PUT", "/Admin/Actor/"+strconv.Itoa(int(actor.ID)), &actorInput)
	json.Unmarshal(response.Body.Bytes(), &actorOutput)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.Equal(t, actorInput.Name, actorOutput.Name)
	assert.Equal(t, actorInput.Active, actorOutput.Active)

	//

	actorInput = Actor{ID: actor.ID + 10}
	response = makeActorTestRequest(adminToken, "PUT", "/Admin/Actor/"+strconv.Itoa(int(actorInput.ID)), &actorInput)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

//Helpers
func makeActorTestRequest(token, method, url string, actor *Actor) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	body := actor.GetJSONBody()
	req, _ := http.NewRequest(method, url, bytes.NewReader([]byte(body)))
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)
	return w
}

func makeReadActorsTestRequest(token string, actor *Actor, limit, offset int, sortField, sortDir string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()

	url := strconv.Itoa(int(actor.ID)) + "&Name=" + actor.Name + "&Limit=" + strconv.Itoa(limit) + "&Offset=" + strconv.Itoa(offset) +
		"&SortField=" + sortField + "&SortDir=" + sortDir

	req, _ := http.NewRequest("GET", "/Admin/Actors?ID="+url, nil)
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)
	return w
}
