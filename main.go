package main

import (
	"log"      // For logging errors or status messages
	"net/http" // For handling HTTP routs and servers

	"webpage/db"     // Custom package fr DB connection
	"webpage/routes" // Custom package that handles routing logic like login
)

func main() {
	// Initializing the database connection
	db.InitDB()

	// Set up routes and assign handler functions
	http.HandleFunc("/", routes.LoginHandler)
	http.HandleFunc("/login", routes.LoginHandler)

	log.Println("Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error", err)
	}
}

