package handlers

import (
	"net/http"

	database "subsmart/internal/db"

	"github.com/gorilla/mux"
)

type App struct {
	Users *database.UserModel
}

func (app *App) Routes() http.Handler {
	r := mux.NewRouter()

	staticDir := "./static"
	fs := http.FileServer(http.Dir(staticDir))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	//
	r.HandleFunc("/", app.HomeHandler).Methods("GET")
	r.HandleFunc("/home", app.HomeHandler).Methods("GET")

	//
	r.HandleFunc("/register", app.HandleRegisterPost).Methods("POST")
	r.HandleFunc("/register", app.HandleRegisterGet).Methods("GET")
	//
	r.HandleFunc("/login", app.HandleLoginPost).Methods("POST")
	r.HandleFunc("/login", app.HandleLoginGet).Methods("GET")

	return r
}
