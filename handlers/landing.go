package handlers

import (
	"html/template" // Rendering HTML
	"log"
	"net/http"       // HTTP server functions
	"webpage/db"     // Database initialization
	"webpage/models" // Interacting with the DB via User model

	"golang.org/x/crypto/bcrypt" // Password hashing and comparison
)

// =========================
// RegisterHandler
// =========================
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize DB connection
	db.InitDB()

	// Load the register.html template from the templates folder
	tmpl := template.Must(template.ParseFiles("templates/register.html"))

	// If user is visiting the page (GET)
	if r.Method == "GET" {
		// Render the registration form
		tmpl.Execute(w, nil)
		return
	}

	// Count number of users in the DB to determine the first user as admin
	var UserCount int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM users").Scan(&UserCount)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Default role assignment
	role := "user"
	if UserCount == 0 {
		// First registered user is an admin
		role = "admin"
	}

	// If the form is submitted (POST)
	if r.Method == "POST" {
		// Parse form data from the request body
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		// Extract input values from the form
		FirstName := r.FormValue("FirstName")
		LastName := r.FormValue("LastName")
		Username := r.FormValue("Username")
		Password := r.FormValue("Password")

		// Register the user via the models layer
		err = models.CreateUser(FirstName, LastName, Username, Password, role)
		if err != nil {
			// Handle creation errors (e.g. duplicate username)
			log.Println("Error in creating user:", err)
			tmpl.Execute(w, map[string]string{"Error": "Username might already exist or a different issue occurred."})
			return
		}

		// On successful registration, show success message
		tmpl.Execute(w, map[string]string{"Success": "Registered successfully, you may now login!"})
	}
}

// =========================
// LoginHandler
// =========================
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Load the login.html template from the templates folder
	tmpl := template.Must(template.ParseFiles("templates/login.html"))

	// If user is visiting the login page via browser (GET request)
	if r.Method == "GET" {
		// Render the login form
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, "Template execution failed", http.StatusInternalServerError)
			return
		}
	}

	// If the login form is submitted (POST)
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		// Extract credentials
		Username := r.FormValue("username")
		Password := r.FormValue("password")

		log.Println("Username entered:", Username)
		log.Println("Entered password:", Password)

		// Fetch user from DB
		user, err := models.GetUserByUsername(Username)
		if err != nil {
			log.Println("Error fetching user:", err)
			tmpl.Execute(w, map[string]string{"Error": "Invalid username or password"})
			return
		}

		log.Println("User found in DB:", user.Username)
		log.Println("Stored hash:", user.PasswordHash)

		// Compare password with stored hash
		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(Password))
		if err != nil {
			log.Println("Password comparison failed:", err)
			tmpl.Execute(w, map[string]string{"Error": "Invalid username or password"})
			return
		}

		// TEMP SUCCESS (you'll add session logic here later)
		w.Write([]byte("Welcome, " + user.Username + "!"))
	}

}
