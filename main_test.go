package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:mysecretpassword@localhost:5400/employee?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	return router
}
