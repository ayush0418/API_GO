package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// Delete data from the employee Table according to Id
func deleteEmployeeId(c *gin.Context) {
	fmt.Println("DELETING EMPLOYEE DATA BY ID")
	c.Header("Content-Type", "application/json")

	// Get the ID of the employee to be deleted
	id := c.Param("id")

	// Delete the employee from the database
	sqlStatement := `DELETE FROM employee WHERE id = $1;`
	stmt, err := db.Prepare(sqlStatement)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	result, err := stmt.Exec(id)

	if err != nil {
		log.Fatal(err)
	}

	// Get the number of rows affected
	n, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	// Return the success message
	if n == 1 {
		c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Employee not found"})
	}
}
