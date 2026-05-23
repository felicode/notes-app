package main

import (
	"database/sql"
	"log"
	"net/http"

	"notes-app/handler"
	"notes-app/store"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./notes.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec(`CREATE TABLE IF NOT EXISTS notes
									id INTEGER PRIMARY KEY AUTOINCREMENT,
									title TEXT NOT NULL,
									body TEXT NOT NULL
		)`)

	h := &handler.NoteHandler{Store: &store.NoteStore{DB: db}}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /notes", h.Create)
	mux.HandleFunc("GET /notes", h.List)
	mux.HandleFunc("PUT /notes/{id}", h.Update)
	mux.HandleFunc("DELETE /notes/{id}", h.Delete)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
