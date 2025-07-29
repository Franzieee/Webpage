package main

import (
	"log"      // For logging errors or status messages
	"net/http" // For handling HTTP routs and servers
	"os"       // Environment variables

	"webpage/db"       // Custom package fr DB connection
	"webpage/handlers" // Custom package that handles routing logic like login

	"github.com/gorilla/sessions" // For handing sessions with Gorilla Mux framework
	"github.com/joho/godotenv"    // For hashing
)

var store *sessions.CookieStore

func main() {
	// Initializing the database connection
	db.InitDB()

	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found or not loaded - relying on actual environment variables")
	}

	sessionSecret := os.Getenv("SESSION_SECRET")
	if sessionSecret == "" {
		log.Fatal("SESSION_SECRET not set in .env")
	}

	store = sessions.NewCookieStore([]byte(sessionSecret))

	// Set up routes and assign handler functions
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		handlers.AdminHandler(w, r, store)
	})

	log.Println("Server started on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error", err)
	}
}
