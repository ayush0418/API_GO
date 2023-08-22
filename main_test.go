package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
	"database/sql"
	"log"
)


func SetUpRouter() *gin.Engine{
    var err error
	db, err = sql.Open("postgres", "postgres://postgres:mysecretpassword@localhost:5400/api_database?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
    router := gin.Default()
    return router
}

/***EMPLOYEE******EMPLOYEE******EMPLOYEE******EMPLOYEE******EMPLOYEE******EMPLOYEE******EMPLOYEE******EMPLOYEE******EMPLOYEE******EMPLOYEE******EMPLOYEE***/
/***EMPLOYEE******EMPLOYEE******EMPLOYEE******EMPLOYEE******EMPLOYEE******EMPLOYEE******EMPLOYEE******EMPLOYEE******EMPLOYEE******EMPLOYEE******EMPLOYEE***/
func TestCreateEmployee (t *testing.T) {
    r:= SetUpRouter()
    r.POST("/emp", createEmployee)

    emp := employee{
        Id:3,
        Name:"xenonstack",
        Gender:"Male",
        LeaveDate:"2022-01-24",
        LeaveDuration:1,
        LeaveType:"SICK LEAVE",
    }
    jsonValue, _ := json.Marshal(emp)
    req, _ := http.NewRequest("POST", "/emp", bytes.NewBuffer(jsonValue))
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
}

func TestGetEmployee (t *testing.T) {
    r := SetUpRouter()
    r.GET("/employee", getEmployee)

    req, _ := http.NewRequest("GET", "/employee", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    
    var employees []employee
    json.Unmarshal(w.Body.Bytes(), &employees)

    // Add assertions to check the content of the response data
    assert.Equal(t, http.StatusOK, w.Code)
    assert.NotEmpty(t, employees)
}

func TestGetEmployeeId (t *testing.T) {
    r := SetUpRouter()
    r.GET("/employee/:id", getEmployeeId)
    
    req, _ := http.NewRequest("GET", "/employee/3", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)
}


func TestDeleteEmployeeId (t *testing.T) {
    r:= SetUpRouter()
    r.DELETE("/employee/:id", deleteEmployeeId)
    req, _ := http.NewRequest("DELETE", "/employee/3", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteALlEmployee (t *testing.T) {
    r:= SetUpRouter()
    r.DELETE("employee", deleteAllEmployee)

    req, _ := http.NewRequest("DELETE", "/employee", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)
}

/***MANAGER******MANAGER******MANAGER******MANAGER******MANAGER******MANAGER******MANAGER******MANAGER******MANAGER******MANAGER******MANAGER******MANAGER***/
/***MANAGER******MANAGER******MANAGER******MANAGER******MANAGER******MANAGER******MANAGER******MANAGER******MANAGER******MANAGER******MANAGER******MANAGER***/
func TestCreateManager (t *testing.T) {
    r:= SetUpRouter()
    r.POST("/man", createManager)

    man := manager{
        Id:3,
        Name:"xenonstack",
        ManName:"xenonstack manager",
        TeamName:"CloudOps",
    }
    jsonValue, _ := json.Marshal(man)
    req, _ := http.NewRequest("POST", "/man", bytes.NewBuffer(jsonValue))
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
}

func TestGetManager (t *testing.T) {
    r := SetUpRouter()
    r.GET("/manager", getManager)

    req, _ := http.NewRequest("GET", "/manager", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    
    var managers []manager
    json.Unmarshal(w.Body.Bytes(), &managers)
    
    // Add assertions to check the content of the response data
    assert.Equal(t, http.StatusOK, w.Code)
    assert.NotEmpty(t, managers)
}

func TestGetManagerId (t *testing.T) {
    r := SetUpRouter()
    r.GET("/manager/:id", getManagerId)
    
    req, _ := http.NewRequest("GET", "/manager/3", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)
}


func TestDeleteManagerId (t *testing.T) {
    r:= SetUpRouter()
    r.DELETE("/manager/:id", deleteManagerId)
    req, _ := http.NewRequest("DELETE", "/manager/3", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteAllManager (t *testing.T) {
    r:= SetUpRouter()
    r.DELETE("manager", deleteAllManager)

    req, _ := http.NewRequest("DELETE", "/manager", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)
}