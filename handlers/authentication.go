package handlers

import (
	"net/http"

	"github.com/gorilla/sessions"
)

func IsAdmin(r *http.Request, store *sessions.CookieStore) bool {
	session, err := store.Get(r, "auth-session")
	if err != nil {
		return false
	}

	role, ok := session.Values["role"].(string)
	return ok && role == "admin"
}