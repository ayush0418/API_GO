package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// CREATING EMPLOYEE ENTRY IN THE DATABASE EMPLOYEE, TABLE EMPLOYEE
func createEmployee(c *gin.Context) {
	fmt.Println("POSTING EMPLOYEE DATA")

	Name := c.PostForm("emp_name")
	TeamName := c.PostForm("team_name")
	LeaveFrom := c.PostForm("leave_from")
	LeaveTo := c.PostForm("leave_to")
	LeaveType := c.PostForm("leave_type")
	Reporter := c.PostForm("reporter")

	if LeaveType == "Sick Leave" {
		fmt.Println("With File Block")

		fileHeader, err := c.FormFile("attachment")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Validate file size 15MB in bytes
		if fileHeader.Size > 15*1024*1024 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds 15MB limit"})
			return
		}
	
		// Validate file extension
		ext := filepath.Ext(fileHeader.Filename)
		if ext != ".pdf" && ext != ".png" && ext != ".txt"{
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file extension. Only PDF and PNG are allowed."})
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer file.Close()

		attachment, err := ioutil.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Saving uploaded file in the uploads directory
		filename := filepath.Join("uploads", fileHeader.Filename)
		err = saveFile(attachment, filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// saving the Form Data in the database
		err = saveForm(Name, TeamName, LeaveFrom, LeaveTo, LeaveType, Reporter, attachment)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Sucessfull With file"})
	} else {
		fmt.Println("Without File Block")

		err := saveForm(Name, TeamName, LeaveFrom, LeaveTo, LeaveType, Reporter, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Sucessfull Without file"})
	}
}

// saving the Form Data in the database
func saveFile(data []byte, filename string) error {
	err := os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// saving the Form Data in the database
func saveForm(Name, TeamName string, LeaveFrom, LeaveTo, LeaveType, Reporter string, attachment []byte) error {
	var err error

	_, err = db.Exec("INSERT INTO employee (emp_name, team_name, leave_from, leave_to, leave_type, reporter, attachment) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		Name, TeamName, LeaveFrom, LeaveTo, LeaveType, Reporter, attachment)
	if err != nil {
		return err
	}

	return nil
}