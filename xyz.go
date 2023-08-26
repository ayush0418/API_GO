package main

// import (
// 	"database/sql"
// 	"encoding/base64"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// 	"strconv"

// 	"github.com/gin-contrib/cors"

// 	"github.com/gin-gonic/gin"
// 	"github.com/joho/godotenv"
// 	_ "github.com/lib/pq"
// )

// type LeaveApplication struct {
// 	Name           string `json:"name"`
// 	Email          string `json:"email"`
// 	LeaveType      string `json:"leave_type"`
// 	FromDate       string `json:"from_date"`
// 	ToDate         string `json:"to_date"`
// 	Manager        string `json:"manager"`
// 	Team           string `json:"team"`
// 	Attachment     []byte `json:"attachment"`
// 	AttachmentLink string `json:"attachment_link"`
// }


// func createLeaveAppTableIfNotExists(db *sql.DB) error {
// 	createTableQuery := `
//     CREATE TABLE IF NOT EXISTS leave_app (
//         name VARCHAR(255) NOT NULL,
// 		email VARCHAR(255) DEFAULT NULL,
//         leave_type VARCHAR(255) NOT NULL,
//         from_date DATE NOT NULL,
//         to_date DATE NOT NULL,
//         manager VARCHAR(255) NOT NULL,
//         team VARCHAR(255) NOT NULL,
//         attachment BYTEA 
//     );
//     `

// 	_, err := db.Exec(createTableQuery)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
// func main() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatal("Error loading .env file:", err)
// 	}

// 	// Connect to PostgreSQL
// 	connStr := os.Getenv("DB_CONN_STRING")
// 	if connStr == "" {
// 		log.Fatal("DB_CONN_STRING environment variable is not set")
// 	}
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal("Failed to connect to PostgreSQL:", err)
// 	}
// 	defer db.Close()

// 	// Create enum types and leave_app table if they don't exist
// 	err = createLeaveAppTableIfNotExists(db)
// 	if err != nil {
// 		log.Fatal("Failed to create leave_app table:", err)
// 	}
// 	r := gin.Default()
// 	r.Use(cors.Default())
// 	r.POST("/post", handleLeaveApplication)
// 	r.GET("/get", getAllLeaveApplications)
// 	r.Run(":8080")
// }

// func handleLeaveApplication(c *gin.Context) {

// 	name := c.PostForm("name")
// 	Email := c.PostForm("email")
// 	LeaveType := c.PostForm("leave_type")
// 	FromDate := c.PostForm("from_date")
// 	ToDate := c.PostForm("to_date")
// 	Manager := c.PostForm("manager")
// 	Team := c.PostForm("team")

// 	if LeaveType == "sickLeave" {

// 		fileHeader, err := c.FormFile("attachment")
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		file, err := fileHeader.Open()
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		defer file.Close()

// 		attachment, err := ioutil.ReadAll(file)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// Saving file attachment to local directory
// 		filename := filepath.Join("uploads", fileHeader.Filename)
// 		err = saveAttachmentToFile(attachment, filename)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		err = saveLeaveApplicationToDB(name, Email, LeaveType, FromDate, ToDate, Manager, Team, attachment)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"message": "Leave application submitted successfully with file"})
// 	} else {

// 		err := saveLeaveApplicationToDB(name, Email, LeaveType, FromDate, ToDate, Manager, Team, nil)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"message": "Leave application submitted successfully without file"})
// 	}
// }

// func saveAttachmentToFile(data []byte, filename string) error {
// 	err := os.MkdirAll(filepath.Dir(filename), os.ModePerm)
// 	if err != nil {
// 		return err
// 	}

// 	err = ioutil.WriteFile(filename, data, os.ModePerm)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func saveLeaveApplicationToDB(name, Email, LeaveType, FromDate, ToDate, Manager string, Team string, attachment []byte) error {

// 	// Connect to PostgreSQL
// 	connStr := os.Getenv("DB_CONN_STRING")
// 	if connStr == "" {
// 		return fmt.Errorf("DB_CONN_STRING environment variable is not set")
// 	}
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()

// 	_, err = db.Exec("INSERT INTO leave_app (name, email, leave_type, from_date, to_date, manager, team, attachment) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
// 		name, Email, LeaveType, FromDate, ToDate, Manager, Team, attachment)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func getAllLeaveApplications(c *gin.Context) {
// 	// Connect to PostgreSQL
// 	connStr := os.Getenv("DB_CONN_STRING")
// 	if connStr == "" {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB_CONN_STRING environment variable is not set"})
// 		return
// 	}
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, email, leave_type, from_date, to_date, manager, team, attachment FROM leave_app")
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	defer rows.Close()

// 	// Iterate over the rows and build the response
// 	var leaveApplications []LeaveApplication
// 	for rows.Next() {
// 		var application LeaveApplication
// 		err := rows.Scan(&application.Name, &application.Email, &application.LeaveType, &application.FromDate, &application.ToDate, &application.Manager, &application.Team, &application.Attachment)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// Generate download link for attachment
// 		attachmentLink, err := generateDownloadLink(application.Attachment)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		application.AttachmentLink = attachmentLink

// 		leaveApplications = append(leaveApplications, application)
// 	}

// 	c.JSON(http.StatusOK, leaveApplications)
// }