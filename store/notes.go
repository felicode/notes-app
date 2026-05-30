// Package store is the entry for notes app
package store

import (
	"database/sql"

	"notes-app/model"
)

type NoteStore struct {
	DB *sql.DB
}

func (s *NoteStore) Create(title, body string) (model.Note, error) {
	res, err := s.DB.Exec("INSERT INTO notes (title, body) VALUES (?, ?)", title, body)
	if err != nil {
		return model.Note{}, err
	}
	id, _ := res.LastInsertId()
	return model.Note{ID: id, Title: title, Body: body}, nil
}

func (s *NoteStore) List() ([]model.Note, error) {
	rows, err := s.DB.Query("SELECT id, title, body FROM notes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []model.Note
	for rows.Next() {
		var n model.Note
		rows.Scan(&n.ID, &n.Title, &n.Body)
		notes = append(notes, n)
	}
	return notes, nil
}

func (s *NoteStore) Update(id int64, title, body string) error {
	_, err := s.DB.Exec("UPDATE notes SET title=?, body=? WHERE id=?", title, body, id)
	return err
}

func (s *NoteStore) Delete(id int64) error {
	_, err := s.DB.Exec("DELETE FROM notes WHERE id=?", id)
	return err
}
