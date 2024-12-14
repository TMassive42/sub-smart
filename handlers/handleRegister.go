package handlers

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"subsmart/internal"
)

func (app *App) HandleRegisterGet(w http.ResponseWriter, r *http.Request) {
	sessionCookie, err := r.Cookie("session_token")
	if err == nil {
		valid, validationErr := app.Users.ValidateSession(sessionCookie.Value)
		if validationErr == nil && valid {
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    "",
			Path:     "/",
			Expires:  time.Now().Add(-1 * time.Hour),
			HttpOnly: true,
		})
	}
	tmpl, err := template.ParseFiles("templates/register.page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func (app *App) HandleRegisterPost(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/register.page.html"))
	form_errors := map[string][]string{}
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirm_password := r.FormValue("confirm")

	if err := internal.ValidateUsername(username); err != nil {
		form_errors["username"] = append(form_errors["username"], err.Error())
	}

	if err := internal.ValidateEmail(email); !err {
		form_errors["email"] = append(form_errors["email"], "Invalid email address")
	}

	// validate passwords
	if password == "" || confirm_password == "" {
		form_errors["password"] = append(form_errors["password"], "Empty password not allowed")
	} else if password != confirm_password {
		form_errors["confirm"] = append(form_errors["confirm"], "Passwords do not match")
	} else if len(password) < 8 {
		form_errors["password"] = append(form_errors["password"], "Password must be at least 8 characters")
	}

	if len(form_errors) > 0 {
		tmpl.Execute(w, map[string]interface{}{
			"Errors": form_errors,
		})
		return
	}

	if exists, _ := app.Users.UserExists(email); exists {
		http.Error(w, "User already exists", http.StatusFound)
		return
	}

	id := internal.UUIDGen()
	password_hash, _ := internal.HashPassword(password)

	if err := app.Users.InsertUser(id, username, email, password_hash); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Headers before redirect: %+v", w.Header())
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
