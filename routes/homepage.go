package routes

import (
	"fmt"      // Used to format and write strings to the reponse
	"net/http" //Used to handle web requests and responses
)

// HomeHandler handles requests to root path "/"
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// We usse fmt.Fprintf to send a simple HTML page as a response
	// This page inclueds a heading and a link to the login page

	fmt.Fprintf(w, `
	<!DOCTYPE html>
		<html>
		<head>
			<title>Home</title>
		</head>
		<body>
			<h1>Welcome to My Personal Website</h1>
			<p><a href="/login">Login </a></p>
			<p><a href="/register">Register</a></p>
		</body>
		</html>
	`)
}
