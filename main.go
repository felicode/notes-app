package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./notes.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// TODO: create my tables here (db.Exec with a CREATE TABLE IF NOT EXISTS)

	mux := http.NewServerMux()

	// TODO: register my routes here, e.g.:
	// mux.HandlerFunc("POST /notes", handler.CreateNote)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
