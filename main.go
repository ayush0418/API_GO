package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:mysecretpassword@localhost:5400/employee?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	
	router := gin.Default()

	router.POST("/emp", createEmployee)
	router.DELETE("/employee", deleteAllEmployee)
	router.DELETE("/employee/:id", deleteEmployeeId)
	router.GET("/employee", getEmployee)
	router.GET("/employee/:id", getEmployeeId)

	router.Run("localhost:8084")
}

type employee struct {
	Id int `json:"id"`
	Name string `json:"name"`
	TeamName string `json:"teamname"`
	LeaveFrom string `json:"leavefrom"`
	LeaveTo string `json:"leaveto"`
	LeaveType string `json:"leavetype"`
	Reporter string `json:"reporter"`
	Attachment	[]byte `json:"attachment"`
}