package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEmployeeId(t *testing.T) {
	r := SetUpRouter()
	r.GET("/employee/:id", getEmployeeId)

	req, _ := http.NewRequest("GET", "/employee/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
