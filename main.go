package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, _ := sql.Open("postgres", "postgres://user:password@localhost:5432/go_course?sslmode=disable")
	defer db.Close()

	res, _ := db.Exec("UPDATE users SET name = $1 WHERE id = $2", "Alice Updated", 1)
	n, _ := res.RowsAffected()
	log.Printf("Updated %d rows", n)

	res, _ = db.Exec("DELETE FROM users WHERE id = $1", 2)
	n, _ = res.RowsAffected()
	log.Printf("Deleted %d rows", n)
}
