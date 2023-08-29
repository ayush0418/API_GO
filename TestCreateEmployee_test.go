package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"mime/multipart"
	"fmt"
)

func TestCreateEmployee(t *testing.T) {
	r := SetUpRouter()
	r.POST("/emp", createEmployee)

	// Define form data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	
	// Add form fields
	formFields := map[string]string{
		"id":         "7",
		"emp_name":   "xenonstack",
		"team_name":  "CloudOps",
		"leave_from": "2023-09-24",
		"leave_to":   "2023-09-25",
		"leave_type": "Sick Leave",
		"reporter":   "Sahil Bansal",
	}

	for key, val := range formFields {
		_ = writer.WriteField(key, val)
	}

	// Add attachment
	if formFields["leave_type"] == "Sick Leave" {
		fmt.Println("Sick Leave")
		fileContents := []byte("file attachment content")
		part, _ := writer.CreateFormFile("attachment", "uploads/test.pdf")
		part.Write(fileContents)
	}

	err := writer.Close()
	if err != nil {
		t.Fatal(err)
	}

	req, _ := http.NewRequest("POST", "/emp", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
}
