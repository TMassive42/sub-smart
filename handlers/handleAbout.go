package handlers

import (
	"html/template"
	"net/http"
)

func (app *App) HandleAboutGet(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/aboutus.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
