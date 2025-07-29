package db

import (
	"database/sql" // Go's standard SQL package
	"log"          // For logging errors or status messages
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver (blank import means init() is called)
)

var DB *sql.DB // Global variable to hold the DB connection

func InitDB() {

	var err error

	// Get connection string from environment variable
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL not set")
	}

	// Open conection to PostgreSQL
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Cannot reach database:", err)
	}

	log.Println("Connected to the database successfully.")
}
