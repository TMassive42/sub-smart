// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"subsmart/handlers"
	database "subsmart/internal/db"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./subsmart.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := &handlers.App{
		Users: &database.UserModel{
			DB: db,
		},
	}

	app.Users.InitTables()

	server := &http.Server{
		Addr:    ":8000",
		Handler: app.Routes(),
	}

	fmt.Printf("Listening on port %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
