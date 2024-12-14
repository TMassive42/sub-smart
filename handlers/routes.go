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

	r.HandleFunc("/about", app.HandleAboutGet).Methods("GET")
	r.HandleFunc("/dashboard", app.HandleDashboard).Methods("GET")

	r.HandleFunc("/logout", app.HandleLogout).Methods("GET")

	r.HandleFunc("/recommendations", app.HandleRecommendations).Methods("GET")
	r.HandleFunc("/cancellations", app.HandleCancellations).Methods("GET")
	r.HandleFunc("/renewals", app.HandleRenewals).Methods("GET")



	return r
}
