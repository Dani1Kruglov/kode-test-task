package notes

import (
	"encoding/json"
	_ "encoding/json"
	"io"
	"kode-task/internal/database"
	"kode-task/internal/middleware"
	"kode-task/internal/model"
	"kode-task/internal/speller"
	"kode-task/internal/storage"
	"log"
	"net/http"
)

func GetNotesHandler(_ http.ResponseWriter, request *http.Request) {
	log.Println("GetNotesHandler work")

	userID := middleware.AuthorizedUsersOnly(request)

	db := database.ConnectToDatabase()

	notes := storage.NotesPostgresStorage{Db: db}
	notesJSON, err := notes.GetNotesByUserId(userID, db)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(notesJSON)
	log.Println("GetNotesHandler end")
}

func StoreNoteHandler(w http.ResponseWriter, request *http.Request) {
	log.Println("StoreNoteHandler work")

	var noteData model.Note

	noteData.UserID = middleware.AuthorizedUsersOnly(request)

	db := database.ConnectToDatabase()

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&noteData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalln(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(request.Body)

	noteData.Title, err = speller.GetCorrectedTextUsingSpeller(noteData.Title)
	if err != nil {
		log.Fatalln(err)
	}
	noteData.Content, err = speller.GetCorrectedTextUsingSpeller(noteData.Content)
	if err != nil {
		log.Fatalln(err)
	}

	note := storage.NotesPostgresStorage{Db: db}
	err = note.StoreNote(noteData, db)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("StoreNoteHandler end")
}

func DeleteNoteHandler(w http.ResponseWriter, request *http.Request) {
	log.Println("DeleteNoteHandler work")

	userID := middleware.AuthorizedUsersOnly(request)

	db := database.ConnectToDatabase()

	var noteIDData model.NoteIdJSON

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&noteIDData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalln(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(request.Body)

	note := storage.NotesPostgresStorage{Db: db}
	err = note.DeleteNote(noteIDData.ID, userID, db)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("DeleteNoteHandler end")
}
