package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"kode-task/internal/config"
	"log"
)

func ConnectToDatabase() *sql.DB {
	conn := config.Get().DatabaseDSN
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
