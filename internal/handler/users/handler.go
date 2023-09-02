package users

import (
	"encoding/json"
	"io"
	"kode-task/internal/database"
	"kode-task/internal/model"
	"kode-task/internal/session_connect"
	"kode-task/internal/storage"
	"log"
	"net/http"
)

func LoginUserHandler(w http.ResponseWriter, request *http.Request) {
	log.Println("LoginUserHandler work")
	var userData model.User

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&userData)
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

	db := database.ConnectToDatabase()

	user := storage.UsersPostgresStorage{Db: db}

	userID, err := user.LoginUser(userData, db)
	if err != nil {
		log.Fatalln(err)
	}
	if userID == 0 {
		log.Fatalln("There is no such user, register first.")
	}

	err = setUserIDInSession(w, request, userID)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Login success")
	log.Println("LoginUserHandler end")
}

func setUserIDInSession(w http.ResponseWriter, request *http.Request, userID int64) error {
	session, _ := session_connect.Store.Get(request, "User-loggedIn")
	session.Values["user_id"] = userID
	err := session.Save(request, w)
	return err
}

func LogoutUserHandler(w http.ResponseWriter, request *http.Request) {
	log.Println("LogoutUserHandler work")
	session, _ := session_connect.Store.Get(request, "User-loggedIn")
	delete(session.Values, "user_id")
	err := session.Save(request, w)
	if err != nil {
		log.Fatalln("Error save session_connect by logout")
	}
	log.Println("Logout success")
	log.Println("LogoutUserHandler end")
}

func StoreUserHandler(w http.ResponseWriter, request *http.Request) {
	log.Println("StoreUserHandler work")
	var userData model.User

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&userData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(request.Body)

	db := database.ConnectToDatabase()

	user := storage.UsersPostgresStorage{Db: db}

	userID, err := user.StoreUser(userData, db)
	if err != nil {
		log.Fatalln(err)
	}
	err = setUserIDInSession(w, request, userID)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("StoreUserHandler end")
}
