package main

import (
	"github.com/gorilla/mux"
	"kode-task/internal/handler/notes"
	"kode-task/internal/handler/users"
	"log"
	"net/http"
)

const PORT = ":8080"

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user/notes", notes.GetNotesHandler).Methods("GET")
	router.HandleFunc("/user/note/store", notes.StoreNoteHandler).Methods("POST")
	router.HandleFunc("/user/note/delete", notes.DeleteNoteHandler).Methods("DELETE")
	router.HandleFunc("/user/register", users.StoreUserHandler).Methods("POST")
	router.HandleFunc("/user/login", users.LoginUserHandler).Methods("POST")
	router.HandleFunc("/user/logout", users.LogoutUserHandler).Methods("POST")

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
