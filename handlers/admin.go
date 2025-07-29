package handlers

import (
	"html/template"
	"net/http"
	"github.com/gorilla/sessions"
)

func AdminHandler(w http.ResponseWriter, r *http.Request, store *sessions.CookieStore) {
	if !IsAdmin(r, store) {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/admin.html"))
	tmpl.Execute(w, nil)
}
