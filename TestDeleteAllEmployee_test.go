package main

// import (
//     "net/http"
//     "net/http/httptest"
//     "testing"
//     "github.com/gin-gonic/gin"
//     "github.com/stretchr/testify/assert"
// 	"database/sql"
// 	"log"
// )

// func TestDeleteALlEmployee (t *testing.T) {
//     r:= SetUpRouter()
//     r.DELETE("employee", deleteAllEmployee)

//     req, _ := http.NewRequest("DELETE", "/employee", nil)
//     w := httptest.NewRecorder()
//     r.ServeHTTP(w, req)

//     assert.Equal(t, http.StatusOK, w.Code)
// }
