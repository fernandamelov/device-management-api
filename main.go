package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fernandamelov/device-management-api/app/repository"
	"github.com/fernandamelov/device-management-api/app/router"
)

func main() {
	waitForDatabase()
	repository.InitializeDatabase()
	r := router.InitializeRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}

func waitForDatabase() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	for {
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Println("Waiting for database to be ready...")
			time.Sleep(5 * time.Second)
			continue
		}

		err = db.Ping()
		if err != nil {
			log.Println("Waiting for database to be ready...")
			time.Sleep(5 * time.Second)
		} else {
			log.Println("Database is ready!")
			db.Close()
			break
		}
	}
}
