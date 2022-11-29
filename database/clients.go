package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {

	DB, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=2022 dbname=db1 sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	return DB

}
