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
	http.HandleFunc("/register", Register)
	http.HandleFunc("/delete-account", DeleteAccount)
	http.HandleFunc("/profil", Profil)
	http.HandleFunc("/posts", Posts)
	http.HandleFunc("/update-user", UpdateUser)
	http.HandleFunc("/forgot", ForgotPassword)
	http.HandleFunc("/reset", ResetPassword)
	http.HandleFunc("/search", Search)
	http.HandleFunc("/api/like", LikeAPIHandler)
	http.HandleFunc("/user/", UserProfil)
	http.HandleFunc("/post/", PostDetail)
	http.HandleFunc("/like", LikePost)
	http.HandleFunc("/comment", CommentPost)
	http.HandleFunc("/comment/delete", DeleteComment)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Serveur lancé sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
