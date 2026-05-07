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
	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		log.Fatal(err)
	}

	queryUsers := `
	CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    avatar TEXT,
    password TEXT NOT NULL
	);`
	if _, err := db.Exec(queryUsers); err != nil {
		log.Fatal(err)
	}
	queryPosts := `
	CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    publication_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`
	if _, err := db.Exec(queryPosts); err != nil {
		log.Fatal(err)
	}
	queryLikes := `
	CREATE TABLE IF NOT EXISTS likes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    UNIQUE(user_id, post_id),
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`

	if _, err := db.Exec(queryLikes); err != nil {
		log.Fatal(err)
	}

	queryComments := `
	CREATE TABLE IF NOT EXISTS comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    publication_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`
	if _, err := db.Exec(queryComments); err != nil {
		log.Fatal(err)
	}
	//`SELECT username FROM users WHERE username = ?);`

	forum.SetDb(db)
	forum.Server()
}
