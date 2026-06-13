package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"notes-app/handler"
	"notes-app/store"

	_ "modernc.org/sqlite"
)

// statusRecorder wraps http.ResponseWriter to capture the status code.
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

// logging logs methods, path, status and duration for every request.
func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rec, r)
		log.Printf("%s %s %d %s", r.Method, r.URL.Path, rec.status, time.Since(start))
	})
}

func main() {
	db, err := sql.Open("sqlite", "./notes.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS notes (
									id INTEGER PRIMARY KEY AUTOINCREMENT,
									title TEXT NOT NULL,
									body TEXT NOT NULL
		)`)
	if err != nil {
		log.Fatal(err)
	}

	h := &handler.NoteHandler{Store: &store.NoteStore{DB: db}}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /notes", h.Create)
	mux.HandleFunc("GET /notes", h.List)
	mux.HandleFunc("PUT /notes/{id}", h.Update)
	mux.HandleFunc("DELETE /notes/{id}", h.Delete)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", logging(mux)))
}
