package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestDeleteEmployeeId (t *testing.T) {
    r:= SetUpRouter()
    r.DELETE("/employee/:id", deleteEmployeeId)

    req, _ := http.NewRequest("DELETE", "/employee/7", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    
    assert.Equal(t, http.StatusOK, w.Code)
}