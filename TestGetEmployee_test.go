package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEmployee(t *testing.T) {
	r := SetUpRouter()
	r.GET("/employee", getEmployee)

	req, _ := http.NewRequest("GET", "/employee", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Add assertions to check the content of the response data
	assert.Equal(t, http.StatusOK, w.Code)
}
