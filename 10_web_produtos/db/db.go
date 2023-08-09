package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func DbConnect() *sql.DB {
	connStr := "user=postgres dbname=go_loja password=153624 host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
