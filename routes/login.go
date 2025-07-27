package routes

import (
	"html/template"  // Rendering HTML
	"net/http"       // HTTP server functions
	"webpage/models" // Interacting with the DB via User model

	"golang.org/x/crypto/bcrypt"
)

// 'LoginHandler' handles both GET (display) and POST (handle) requests
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Loads login.html template from the template folder:
	tmpl := template.Must(template.ParseFiles("templates/login.html"))

	// If user is visitng the login page via browser (GET request)
	if r.Method == "GET" {
		//Render the form when accessed via browser
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, "Template execution failed", http.StatusInternalServerError)
			return
		}
	}

	// If the form is submitted (POST)
	if r.Method == "POST" {
		//Parse form data from the request body
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		// Get the input values from the form
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Find user in the database by username
		user, err := models.GetUserByUsername(username)
		if err != nil {
			tmpl.Execute(w, map[string]string{"Error": "Invalid username or password"})
			return
		}

		// Compare passwords whether hashed password from the database matches the plaintext entered by the user during login.
		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

		if err != nil {
			// If the comparison is wrong (wrong passowrd), shows an error message
			tmpl.Execute(w, map[string]string{"Error": "Invalid username or password"})
			return
		}

		//If there isn't any errors on the username and password, the user is logged in successfully

		// Successful login displays a welcome message:
		w.Write([]byte("Welcome," + user.FirstName + "!"))
	}
}
