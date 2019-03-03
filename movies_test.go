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

//CreateMovie

func TestCreateMovieEndpoint(t *testing.T) {
	_, adminToken := setupTesting()

	var director Director
	success := makeCreateDirectorTestRequest(adminToken, "Wes Anderson", true)
	json.Unmarshal(success.Body.Bytes(), &director)

	var movie Movie
	success = makeCreateMovieTestRequest(adminToken, "Moonrise Kingdom", true, int(director.ID))
	json.Unmarshal(success.Body.Bytes(), &movie)
	assert.Equal(t, http.StatusOK, success.Code)
	assert.NotEmpty(t, movie.ID)
	assert.Equal(t, "Moonrise Kingdom", movie.Name)
	assert.True(t, movie.Active)
	assert.NotEmpty(t, movie.DateCreated)
}

func TestCreateMovieInvalid(t *testing.T) {
	_, adminToken := setupTesting()

	var director Director
	success := makeCreateDirectorTestRequest(adminToken, "Wes Anderson", true)
	json.Unmarshal(success.Body.Bytes(), &director)

	makeCreateMovieTestRequest(adminToken, "Moonrise Kingdom", true, int(director.ID))
	empty := makeCreateMovieTestRequest(adminToken, "", false, int(director.ID))
	noDirector := makeCreateMovieTestRequest(adminToken, "Movie", false, 0)

	assert.Equal(t, http.StatusBadRequest, empty.Code)
	assert.True(t, strings.Contains(empty.Body.String(), "name and director required"))

	assert.Equal(t, http.StatusBadRequest, noDirector.Code)
	assert.True(t, strings.Contains(noDirector.Body.String(), "name and director required"))
}

//ReadMovie

func TestReadMovieEndpoint(t *testing.T) {
	_, adminToken := setupTesting()

	var director Director
	success := makeCreateDirectorTestRequest(adminToken, "Wes Anderson", true)
	json.Unmarshal(success.Body.Bytes(), &director)

	var movie Movie
	success = makeCreateMovieTestRequest(adminToken, "Moonrise Kingdom", true, int(director.ID))
	json.Unmarshal(success.Body.Bytes(), &movie)

	success = makeReadMovieTestRequest(adminToken, int(movie.ID))
	json.Unmarshal(success.Body.Bytes(), &movie)
	assert.Equal(t, http.StatusOK, success.Code)
	assert.NotEmpty(t, movie.ID)
	assert.Equal(t, "Moonrise Kingdom", movie.Name)
	assert.True(t, movie.Active)
	assert.NotEmpty(t, movie.DateCreated)
}

func TestReadMovieInvalid(t *testing.T) {
	_, adminToken := setupTesting()

	notFound := makeReadMovieTestRequest(adminToken, 1)
	assert.Equal(t, http.StatusBadRequest, notFound.Code)
	assert.True(t, strings.Contains(notFound.Body.String(), "record not found"))
}

//ReadMovies

func TestReadMoviesEndpoint(t *testing.T) {
	_, adminToken := setupTesting()

	var director Director
	success := makeCreateDirectorTestRequest(adminToken, "Wes Anderson", true)
	json.Unmarshal(success.Body.Bytes(), &director)

	makeCreateMovieTestRequest(adminToken, "Moonrise Kingdom", true, int(director.ID))
	makeCreateMovieTestRequest(adminToken, "Moonrise Kingdom 2", true, int(director.ID))
	makeCreateMovieTestRequest(adminToken, "Bridge of Spies", true, int(director.ID))
	makeCreateMovieTestRequest(adminToken, "Alice in Wonderland", true, int(director.ID))

	var movies []Movie
	success = makeReadMoviesTestRequest(adminToken, 0, "Moonrise", 1, 0, "ID", "DESC")
	json.Unmarshal(success.Body.Bytes(), &movies)
	assert.Equal(t, http.StatusOK, success.Code)
	assert.Equal(t, 1, len(movies))
	assert.NotEmpty(t, movies[0].ID)
	assert.Equal(t, "Moonrise Kingdom 2", movies[0].Name)

	success = makeReadMoviesTestRequest(adminToken, int(movies[0].ID), "", 100, 0, "", "")
	json.Unmarshal(success.Body.Bytes(), &movies)
	assert.Equal(t, http.StatusOK, success.Code)
	assert.Equal(t, 1, len(movies))
	assert.NotEmpty(t, movies[0].ID)
	assert.Equal(t, "Moonrise Kingdom 2", movies[0].Name)

	success = makeReadMoviesTestRequest(adminToken, 0, "", 2, 2, "Name", "DESC")
	json.Unmarshal(success.Body.Bytes(), &movies)
	assert.Equal(t, http.StatusOK, success.Code)
	assert.Equal(t, 2, len(movies))
	assert.Equal(t, "Bridge of Spies", movies[0].Name)
	assert.Equal(t, "Alice in Wonderland", movies[1].Name)
}

//UpdateMovie

func TestUpdateMovieEndpoint(t *testing.T) {
	_, adminToken := setupTesting()

	var director Director
	success := makeCreateDirectorTestRequest(adminToken, "Wes Anderson", true)
	json.Unmarshal(success.Body.Bytes(), &director)
	var director2 Director
	success = makeCreateDirectorTestRequest(adminToken, "Steven Spielberg", true)
	json.Unmarshal(success.Body.Bytes(), &director2)

	var movie Movie
	success = makeCreateMovieTestRequest(adminToken, "Moonrise Kingdom", true, int(director.ID))
	json.Unmarshal(success.Body.Bytes(), &movie)

	success = makeUpdateMovieTestRequest(adminToken, int(movie.ID), "Bridge of Spies", false, int(director2.ID))
	json.Unmarshal(success.Body.Bytes(), &movie)
	assert.Equal(t, http.StatusOK, success.Code)
	assert.NotEmpty(t, movie.ID)
	assert.Equal(t, "Bridge of Spies", movie.Name)
	assert.False(t, movie.Active)
	assert.Equal(t, "Steven Spielberg", movie.Director.Name)
}

func TestUpdateMovieInvalid(t *testing.T) {
	_, adminToken := setupTesting()

	var director Director
	success := makeCreateDirectorTestRequest(adminToken, "Wes Anderson", true)
	json.Unmarshal(success.Body.Bytes(), &director)

	var movie Movie
	notFound := makeUpdateMovieTestRequest(adminToken, 0, "", false, int(director.ID))
	json.Unmarshal(notFound.Body.Bytes(), &movie)
	assert.Equal(t, http.StatusBadRequest, notFound.Code)
	assert.True(t, strings.Contains(notFound.Body.String(), "record not found"))
}

//Helpers

func makeCreateMovieTestRequest(token, name string, active bool, directorID int) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	body := `{
		"name": "` + name + `",
		"active": ` + strconv.FormatBool(active) + `,
		"director_id": ` + strconv.Itoa(directorID) + `
	}`
	req, _ := http.NewRequest("POST", "/Admin/Movie", bytes.NewReader([]byte(body)))
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)
	return w
}

func makeReadMovieTestRequest(token string, id int) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/Admin/Movie/"+strconv.Itoa(id), nil)
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)
	return w
}

func makeReadMoviesTestRequest(token string, id int, name string, limit, offset int, sortField, sortDir string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/Admin/Movies?ID="+strconv.Itoa(id)+"&Name="+name+
		"&Limit="+strconv.Itoa(limit)+"&Offset="+strconv.Itoa(offset)+"&SortField="+sortField+"&SortDir="+sortDir, nil)
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)
	return w
}

func makeUpdateMovieTestRequest(token string, id int, name string, active bool, directorID int) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	body := `{
		"name": "` + name + `",
		"active": ` + strconv.FormatBool(active) + `,
		"director_id": ` + strconv.Itoa(directorID) + `
	}`
	req, _ := http.NewRequest("PUT", "/Admin/Movie/"+strconv.Itoa(id), bytes.NewReader([]byte(body)))
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)
	return w
}
