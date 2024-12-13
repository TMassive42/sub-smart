package handlers

import (
	"html/template"
	"net/http"
)

type PageData struct {
	IsLoggedIn bool
	Username   string
}

func (app *App) HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		IsLoggedIn: false,
	}

	// Get current user if logged in
	sessionCookie, err := r.Cookie("session_token")
	if err == nil {
		usernameCookie, err := r.Cookie("username")
		if err == nil {
			if valid, err := app.Users.ValidateSession(sessionCookie.Value); valid && err == nil {
				data.IsLoggedIn = true
				data.Username = usernameCookie.Value
			}
		}
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
