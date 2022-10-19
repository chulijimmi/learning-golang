package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAlbums(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "[]", w.Body.String())
}

func TestCreateAlbums(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	var jsonData = []byte(`{
		"id": 1,
		"title": "Lucid Dream",
		"artist": "Juice Wrld",
		"price": 100.5
	}`)
	var expectedData = "{\"id\":1,\"title\":\"Lucid Dream\",\"artist\":\"Juice Wrld\",\"price\":100.5}"
	req, _ := http.NewRequest("POST", "/albums/create", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, expectedData, w.Body.String())
}
