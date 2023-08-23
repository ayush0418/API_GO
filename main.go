package main

import (
	"fmt"
	"database/sql"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	// var err error
	err := godotenv.Load(".env")

	if err != nil {
        log.Fatal("Error loading .env file:", err)
    }
    connStr := os.Getenv("DB_CONN_STRING")
    if connStr == "" {
        log.Fatal("DB_CONN_STRING environment variable is not set")
	}
	db, err := sql.Open("postgres", connStr)

    if err != nil {

        log.Fatal("Failed to connect to PostgreSQL:", err)

    }

    defer db.Close()
	router := gin.Default()

	router.POST("/emp", createEmployee)
	router.DELETE("/employee", deleteAllEmployee)
	router.DELETE("/employee/:id", deleteEmployeeId)
	router.GET("/employee", getEmployee)
	router.GET("/employee/:id", getEmployeeId)

	router.POST("/man", createManager)
	router.DELETE("/manager", deleteAllManager)
	router.DELETE("/manager/:id", deleteManagerId)
	router.GET("/manager", getManager)
	router.GET("/manager/:id", getManagerId)

	router.Run("localhost:8086")
}


type employee struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Gender string `json:"gender"`
	LeaveDate string `json:"leavedate"`
	LeaveDuration int `json:"leaveduration"`
	LeaveType string `json:"leavetype"`
}

//returns a list of EMPLOYEE from the database
func getEmployee(c *gin.Context) {
	connStr := os.Getenv("DB_CONN_STRING")
    if connStr == "" {
        log.Fatal("DB_CONN_STRING environment variable is not set")
	}
	db, err := sql.Open("postgres", connStr)

    if err != nil {

        log.Fatal("Failed to connect to PostgreSQL:", err)

    }

    defer db.Close()
	fmt.Println("GETTING EMPLOYEE DATA")
	c.Header("Content-Type", "application/json")
	fmt.Println("h1")
	rows, err := db.Query("SELECT * FROM employee")
	fmt.Println("h2")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("h3")

	defer rows.Close()
	var employees []employee

	for rows.Next() {
		var e employee
		err := rows.Scan(&e.Id, &e.Name, &e.Gender, &e.LeaveDate, &e.LeaveDuration, &e.LeaveType)
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

// returns a row of EMPLOYEE from the database according to the id 
func getEmployeeId(c *gin.Context) {
	fmt.Println("GETTING EMPLOYEE ID DATA")
	c.Header("Content-Type", "application/json")

	// Get the ID of the employee to be fetched
	id := c.Param("id")
	// Query the database for the employee with the given ID
	sqlStatement := `SELECT * FROM employee WHERE emp_id=$1;`
	row := db.QueryRow(sqlStatement,id)
	
	var e employee
	err := row.Scan(&e.Id, &e.Name, &e.Gender, &e.LeaveDate, &e.LeaveDuration, &e.LeaveType)
	
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			log.Fatal(err)
		}
	}
	

	c.IndentedJSON(http.StatusOK, e)
}


// CREATING EMPLOYEE ENTRY IN THE DATABASE API_DATABASE, TABLE EMPLOYEE
func createEmployee(c *gin.Context) {
	fmt.Println("POSTING EMPLOYEE DATA")
	var awesomeAlbum employee

	if err := c.BindJSON(&awesomeAlbum); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	stmt, err := db.Prepare("INSERT INTO employee ( emp_id, emp_name, emp_gender, emp_leavedate, emp_leaveduration, emp_leavetype) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	if _, err := stmt.Exec(awesomeAlbum.Id, awesomeAlbum.Name, awesomeAlbum.Gender, awesomeAlbum.LeaveDate, awesomeAlbum.LeaveDuration, awesomeAlbum.LeaveType); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, awesomeAlbum)
}

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

// Delete data from the employee Table according to Id
func deleteEmployeeId(c *gin.Context) {
	fmt.Println("DELETING EMPLOYEE DATA BY ID")
	c.Header("Content-Type", "application/json")
	
	// Get the ID of the employee to be deleted
	id := c.Param("id")
	
	// Delete the employee from the database
	sqlStatement := `DELETE FROM employee WHERE emp_id = $1;`
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
	
	


type manager struct {
	Id int `json:id"`
	Name string `json:"name"`
	ManName string `json:"man_name"`
	TeamName string `json:"teamname"`
}

//returns a list of MANAGER from the database
func getManager(c *gin.Context) {
	fmt.Println("GETTING MANAGER DATA")
	c.Header("Content-Type", "application/json")

	rows, err := db.Query("SELECT * FROM manager")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var managers []manager

	for rows.Next() {
		var m manager
		err := rows.Scan(&m.Id, &m.Name, &m.ManName, &m.TeamName)
		if err != nil {
			log.Fatal(err)
		}
		managers = append(managers, m)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, managers)
}

// returns a row of MANAGER from the database according to the id
func getManagerId(c *gin.Context) {
	fmt.Println("GETTING MANAGER ID DATA")
	c.Header("Content-Type", "application/json")

	// Get the ID of the employee to be fetched
	id := c.Param("id")

	// Query the database for the manager with the given ID
	sqlStatement := `SELECT * FROM manager WHERE id=$1;`
	row := db.QueryRow(sqlStatement,id)

	var m manager
	err := row.Scan(&m.Id, &m.Name, &m.ManName, &m.TeamName)

	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			log.Fatal(err)
		}
	}

	c.IndentedJSON(http.StatusOK, m)
}

// CREATING MANAGER ENTRY IN THE DATABASE API_DATABASE, TABLE MANAGER
func createManager(c *gin.Context) {
	fmt.Println("POSTING EMPLOYEE DATA")
	var awesomeAlbum manager

	if err := c.BindJSON(&awesomeAlbum); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	stmt, err := db.Prepare("INSERT INTO manager ( id, emp_name, man_name, teamname) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	if _, err := stmt.Exec(awesomeAlbum.Id, awesomeAlbum.Name, awesomeAlbum.ManName, awesomeAlbum.TeamName); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, awesomeAlbum)
}

// Delete all the data from the Manager Table
func deleteAllManager(c *gin.Context) {
	fmt.Println("DELETING ALL MANAGER DATA")
	c.Header("Content-Type", "application/json")

	// Delete the employee from the database
	sqlStatement := `DELETE FROM manager;`
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
		c.JSON(http.StatusOK, gin.H{"message": "All Manager deleted successfully"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Manager not found"})
	}
}

// Delete data from the Manager Table according to Id
func deleteManagerId(c *gin.Context) {
	fmt.Println("DELETING MANAGER DATA BY ID")
	c.Header("Content-Type", "application/json")
	
	// Get the ID of the employee to be deleted
	id := c.Param("id")
	
	// Delete the employee from the database
	sqlStatement := `DELETE FROM manager WHERE id = $1;`
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
		c.JSON(http.StatusOK, gin.H{"message": "Manager deleted successfully"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Manager not found"})
	}
}