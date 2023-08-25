package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// CREATING EMPLOYEE ENTRY IN THE DATABASE API_DATABASE, TABLE EMPLOYEE
func createEmployee(c *gin.Context) {
	fmt.Println("POSTING EMPLOYEE DATA")
	var e employee

	if err := c.BindJSON(&e); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Handle the uploaded file
	file, err := c.FormFile("attachment")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Attachment not provided"})
		return
	}

	// Save the file to a location (e.g., local directory or cloud storage)
	// For example, you can save the file in a directory named "uploads"
	filepath := "uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving file"})
		return
	}

	e.Attachment = filepath // Store the file path in the database

	stmt, err := db.Prepare("INSERT INTO employee (id, emp_name, team_name, leave_from, leave_to, leave_type, reporter, attachment) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	if _, err := stmt.Exec(e.Id, e.Name, e.TeamName, e.LeaveFrom, e.LeaveTo, e.LeaveType, e.Reporter, e.Attachment); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, e)
}
