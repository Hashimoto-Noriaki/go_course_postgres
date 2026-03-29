package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, _ := sql.Open("postgres", "postgres://user:password@localhost:5432/go_course?sslmode=disable")
	defer db.Close()

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Table created!")
}
