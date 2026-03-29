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
		INSERT INTO users (name, email)
		VALUES ($1, $2)
	`, "Alice", "alice@example.com")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted!")
}
