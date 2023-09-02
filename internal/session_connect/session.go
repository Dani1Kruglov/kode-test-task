package session_connect

import "github.com/gorilla/sessions"

var Store = sessions.NewCookieStore([]byte("login-secret"))
