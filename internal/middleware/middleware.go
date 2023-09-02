package middleware

import (
	"kode-task/internal/session_connect"
	"log"
	"net/http"
)

func AuthorizedUsersOnly(request *http.Request) int64 {
	session, _ := session_connect.Store.Get(request, "User-loggedIn")
	if session.Values["user_id"] == nil {
		log.Fatalln("You do not have access to notes because you are not logged in to your account")
	}
	return session.Values["user_id"].(int64)
}
