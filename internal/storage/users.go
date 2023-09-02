package storage

import (
	"database/sql"
	"kode-task/internal/model"
	"log"
)

type UsersPostgresStorage struct {
	Db *sql.DB
}

func (u *UsersPostgresStorage) LoginUser(userJSON model.User, conn *sql.DB) (int64, error) {

	rows, err := conn.Query(`SELECT * FROM users WHERE email = $1 AND password = $2`, userJSON.Email, userJSON.Password)
	if err != nil {
		return 0, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)

	var user model.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
			log.Fatal(err)
		}
	}
	if err := rows.Err(); err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (u *UsersPostgresStorage) StoreUser(user model.User, conn *sql.DB) (int64, error) {
	var userID int64
	err := conn.QueryRow(`INSERT INTO users (name, email, password) 
    SELECT $1, $2, $3 
    WHERE NOT EXISTS (SELECT * FROM users WHERE email = $2::varchar(255))
    RETURNING id`, user.Name, user.Email, user.Password).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
