package storage

import (
	"database/sql"
	"encoding/json"
	"kode-task/internal/model"
	"log"
)

type NotesPostgresStorage struct {
	Db *sql.DB
}

func (n *NotesPostgresStorage) GetNotesByUserId(userID int64, conn *sql.DB) ([]byte, error) {
	rows, err := conn.Query(`SELECT * FROM notes WHERE user_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)

	var notes []model.Note
	for rows.Next() {
		var note model.Note

		if err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.UserID); err != nil {
			log.Fatal(err)
		}
		notes = append(notes, note)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	notesJSON, err := json.Marshal(notes)
	if err != nil {
		return nil, err
	}

	return notesJSON, nil
}

func (n *NotesPostgresStorage) StoreNote(note model.Note, conn *sql.DB) error {
	_, err := conn.Exec(`INSERT INTO notes (title, content, user_id) SELECT $1, $2, $3 WHERE NOT EXISTS (SELECT * FROM notes WHERE title = $1::varchar(255) AND user_id = $3::int)`, note.Title, note.Content, note.UserID)
	return err
}

func (n *NotesPostgresStorage) DeleteNote(noteID int64, userID int64, conn *sql.DB) error {
	result, err := conn.Exec(`DELETE FROM notes WHERE id = $1 AND user_id = $2`, noteID, userID)
	answer, _ := result.RowsAffected()
	if answer == 0 {
		log.Println("You didn't find any notes with this id.")
	}
	return err
}
