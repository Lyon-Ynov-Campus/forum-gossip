package src

import (
	"html/template"
	"net/http"
)

func Profil(w http.ResponseWriter, r *http.Request) {
	id := getUser(r)

	if id == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	var user struct {
		Username string
		Email    string
	}
	err := db.QueryRow(
		"SELECT username, email FROM users WHERE id = ?", id).Scan(&user.Username, &user.Email)
	if err != nil {
		http.Error(w, "Erreur récupération user", 500)
		return
	}
	tmpl, _ := template.ParseFiles("templates/pageUtilisateur.html")
	tmpl.Execute(w, user)
}
