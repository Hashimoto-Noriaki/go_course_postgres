package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, _ := sql.Open("postgres", "postgres://user:password@localhost:5432/go_course?sslmode=disable")
	defer db.Close()

	var id int
	var name, email string
	err := db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", 1).Scan(&id, &name, &email)
	if err == sql.ErrNoRows {
		log.Println("not found")
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("id=%d name=%s email=%s", id, name, email)
}
