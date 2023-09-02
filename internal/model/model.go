package model

type Note struct {
	ID      int64  `json:"note_id"`
	Title   string `json:"note_title"`
	Content string `json:"note_content"`
	UserID  int64  `json:"user_id"`
}

type User struct {
	ID       int64  `json:"user_id"`
	Name     string `json:"user_name"`
	Email    string `json:"user_email"`
	Password string `json:"user_password"`
}

type NoteIdJSON struct {
	ID int64 `json:"note_id"`
}
