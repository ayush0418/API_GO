package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// returns a row of EMPLOYEE from the database according to the id
func getEmployeeId(c *gin.Context) {
	fmt.Println("GETTING EMPLOYEE ID DATA")
	c.Header("Content-Type", "application/json")

	// Get the ID of the employee to be fetched
	id := c.Param("id")

	// Query the database for the employee with the given ID
	sqlStatement := `SELECT * FROM employee WHERE id=$1;`

	row := db.QueryRow(sqlStatement, id)
	var e employee
	err := row.Scan(&e.Id, &e.Name, &e.TeamName, &e.LeaveFrom, &e.LeaveTo, &e.LeaveType, &e.Reporter, &e.Attachment, &e.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			log.Fatal(err)
		}
	}

	c.IndentedJSON(http.StatusOK, e)
}
