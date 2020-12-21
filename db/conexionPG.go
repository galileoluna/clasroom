package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // here
)

var db *sql.DB

func init() {
	tmpDB, err := sql.Open("postgres", "dbname=avalith user=postgres password=root host=localhost sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db = tmpDB

	log.Println("Base de datos conectada...")
}
