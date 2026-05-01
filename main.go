package main

import (
	"database/sql"
	forum "forum-gossip/src"
	"log"

	_ "modernc.org/sqlite"
)

//var db = *sql.DB

func main() {

	dbString := "./data/database.db"

	db, err := sql.Open("sqlite", dbString)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	queryUsers := `
	CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL,
    username TEXT NOT NULL,
    password TEXT NOT NULL
	);
	`
	if _, err := db.Exec(queryUsers); err != nil {
		log.Fatal(err)
	}

	forum.SetDb(db)
	forum.Server()
}
