package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// returns a list of EMPLOYEE from the database
func getEmployee(c *gin.Context) {
	fmt.Println("GETTING EMPLOYEE DATA")
	c.Header("Content-Type", "application/json")

	rows, err := db.Query("SELECT * FROM employee")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var employees []employee

	for rows.Next() {
		var e employee
		err := rows.Scan(&e.Id, &e.Name, &e.TeamName, &e.LeaveFrom, &e.LeaveTo, &e.LeaveType, &e.Reporter, &e.Attachment)
		if err != nil {
			log.Fatal(err)
		}
		employees = append(employees, e)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, employees)
}
