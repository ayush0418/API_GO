package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateEmployee(t *testing.T) {
	r := SetUpRouter()
	r.POST("/emp", createEmployee)

	emp := employee{
		Id:        7,
		Name:      "xenonstack",
		TeamName:  "CloudOps",
		LeaveFrom: "2023-09-24",
		LeaveTo:   "2023-09-25",
		LeaveType: "Sick Leave",
		Reporter:  "Sahil Bansal",
	}

	jsonValue, _ := json.Marshal(emp)
	req, _ := http.NewRequest("POST", "/emp", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
}
