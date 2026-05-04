package src

import (
	"database/sql"
	"fmt"
	"net/http"
)

var db *sql.DB

func SetDb(database *sql.DB) {
	db = database
}

func Server() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/logout", Logout)
	http.HandleFunc("/delete-account", DeleteAccount)
	http.HandleFunc("/profil", Profil)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Serveur lancé sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
