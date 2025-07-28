package routes

import (
	"html/template" // Rendering HTML
	"log"
	"net/http"       // HTTP server functions
	"webpage/models" // Interacting with the DB via User model
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	// Loads the register.html from the templates folder
	tmpl := template.Must(template.ParseFiles("templates/register.html"))

	// If user is visiting the page
	if r.Method == "GET" {
		tmpl.Execute(w, nil)
		return
	}

	// If user wants to register
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		// Extract values needed
		FirstName := r.FormValue("FirstName")
		LastName := r.FormValue("LastName")
		Username := r.FormValue("Username")
		Password := r.FormValue("Password")

		// Register users using models layer
		err = models.CreateUser(FirstName, LastName, Username, Password)
		if err != nil {
			log.Println("Error in creating user:", err)
			tmpl.Execute(w, map[string]string{"Error": "Username might already exist or a different issue occoured."})
			return
		}

		// On Successful register, message would show
		tmpl.Execute(w, map[string]string{"Success": "Registered successfully, you may now login!"})
	}
}
