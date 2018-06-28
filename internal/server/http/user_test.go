package http

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithRecorder(t *testing.T) {
	/*userService := user.NewService()

	router := NewHandler(userService)

	response := httptest.NewRecorder()
	endpoint := "/v1/users"
	body := []byte(`{"name": "Ricardo Longa", "age": 31}`)

	req, _ := http.NewRequest("POST", endpoint, bytes.NewReader(body))

	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusCreated, response.Code)*/
}

func TestUserWithRecorder(t *testing.T) {
	router := NewHandler(nil)

	response := httptest.NewRecorder()
	endpoint := "/v1/users/1"

	req, _ := http.NewRequest("GET", endpoint, nil)

	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)

	body, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	expectedBody := []byte(`{"id": "1", "name":"Fernando"}`)

	assert.JSONEq(t, string(expectedBody), string(body))
}
