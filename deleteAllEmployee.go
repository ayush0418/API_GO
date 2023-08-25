package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// Delete all the data from the employee Table
func deleteAllEmployee(c *gin.Context) {
	fmt.Println("DELETING ALL EMPLOYEE DATA")
	c.Header("Content-Type", "application/json")

	// Delete the employee from the database
	sqlStatement := `DELETE FROM employee;`
	stmt, err := db.Prepare(sqlStatement)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	result, err := stmt.Exec()

	if err != nil {
		log.Fatal(err)
	}

	// Get the number of rows affected
	n, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	// Return the success message
	if n > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "All Employee deleted successfully"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Employee not found"})
	}
}
