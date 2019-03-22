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

//CreateMovie
func TestCreateMovie(t *testing.T) {
	_, adminToken := setupTesting()

	director := database.GetTestingDirectors()[0]
	movieInput := Movie{Name: "name", Year: 1, Rating: 1, Active: true, DirectorID: int(director.ID)}
	movieOutput := Movie{}

	response := makeMovieTestRequest(adminToken, "POST", "/Admin/Movie", &movieInput)
	json.Unmarshal(response.Body.Bytes(), &movieOutput)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.Equal(t, movieInput.Name, movieOutput.Name)
	assert.Equal(t, movieInput.Year, movieOutput.Year)
	assert.Equal(t, movieInput.Rating, movieOutput.Rating)
	assert.Equal(t, movieInput.Active, movieOutput.Active)
	assert.Equal(t, director, movieOutput.Director)

	//errors

	movieInput = Movie{}
	response = makeMovieTestRequest(adminToken, "POST", "/Admin/Movie", &movieInput)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

//ReadMovie
func TestReadMovie(t *testing.T) {
	_, adminToken := setupTesting()

	movie := database.GetTestingMovies()[0]

	movieInput := Movie{}
	movieOutput := Movie{}

	response := makeMovieTestRequest(adminToken, "GET", "/Admin/Movie/"+strconv.Itoa(int(movie.ID)), &movieInput)
	json.Unmarshal(response.Body.Bytes(), &movieOutput)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.Equal(t, movie.Name, movieOutput.Name)
	assert.Equal(t, movie.Year, movieOutput.Year)
	assert.Equal(t, movie.Rating, movieOutput.Rating)
	assert.Equal(t, movie.Active, movieOutput.Active)
	assert.Equal(t, movie.GetDirector(), movieOutput.Director)

	//errors

	movieInput = Movie{}
	response = makeMovieTestRequest(adminToken, "GET", "/Admin/Movie/0", &movieInput)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

//ReadMovies
func TestReadMovies(t *testing.T) {
	_, adminToken := setupTesting()

	movies := database.GetTestingMovies()
	movieInput := Movie{Name: "name"}
	moviesOutput := []Movie{}

	response := makeReadMoviesTestRequest(adminToken, &movieInput, 2, 1, "year", "DESC")
	json.Unmarshal(response.Body.Bytes(), &moviesOutput)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.Equal(t, 2, len(moviesOutput))
	assert.Equal(t, movies[0].Name, moviesOutput[1].Name)
	assert.Equal(t, movies[1].Name, moviesOutput[0].Name)
}

//UpdateMovie
func TestUpdateMovieEndpoint(t *testing.T) {
	_, adminToken := setupTesting()

	movies := database.GetTestingMovies()
	movieInput := Movie{Name: "name", Year: 2, Rating: 2, Active: false, DirectorID: int(movies[0].GetDirector().ID + 1)}
	movieOutput := Movie{}

	response := makeMovieTestRequest(adminToken, "PUT", "/Admin/Movie/"+strconv.Itoa(int(movies[0].ID)), &movieInput)
	json.Unmarshal(response.Body.Bytes(), &movieOutput)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.Equal(t, movieInput.Name, movieOutput.Name)
	assert.Equal(t, movieInput.Year, movieOutput.Year)
	assert.Equal(t, movieInput.Rating, movieOutput.Rating)
	assert.Equal(t, movieInput.Active, movieOutput.Active)
	assert.Equal(t, movieInput.GetDirector(), movieOutput.Director)

	//errors

	movieInput = Movie{ID: movies[2].ID + 1}
	response = makeMovieTestRequest(adminToken, "PUT", "/Admin/Movie/"+strconv.Itoa(int(movieInput.ID)), &movieInput)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	movieInput = Movie{ID: movies[0].ID, Name: movies[1].Name}
	response = makeMovieTestRequest(adminToken, "PUT", "/Admin/Movie/"+strconv.Itoa(int(movieInput.ID)), &movieInput)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

//Helpers
func makeMovieTestRequest(token, method, url string, movie *Movie) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	body := movie.GetJSONBody()
	req, _ := http.NewRequest(method, url, bytes.NewReader([]byte(body)))
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)
	return w
}

func makeReadMoviesTestRequest(token string, movie *Movie, limit, offset int, sortField, sortDir string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()

	url := strconv.Itoa(int(movie.ID)) + "&Name=" + movie.Name + "&Limit=" + strconv.Itoa(limit) + "&Offset=" + strconv.Itoa(offset) +
		"&SortField=" + sortField + "&SortDir=" + sortDir

	req, _ := http.NewRequest("GET", "/Admin/Movies?ID="+url, nil)
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)
	return w
}
