package src

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

func Profil(w http.ResponseWriter, r *http.Request) {
	id := getUser(r)
	if id == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	var avatarNull sql.NullString
	var username, email string

	err := db.QueryRow(
		"SELECT username, email, avatar FROM users WHERE id = ?", id).Scan(&username, &email, &avatarNull)
	if err != nil {
		http.Error(w, "Erreur récupération utilisateur", 500)
		return
	}

	avatar := "/static/default.png"
	if avatarNull.Valid && avatarNull.String != "" {
		avatar = avatarNull.String
	}

	data := map[string]interface{}{
		"Username": username,
		"Email":    email,
		"Avatar":   avatar,
	}

	tmpl, err := template.ParseFiles("templates/pageUtilisateur.html")
	if err != nil {
		http.Error(w, "Erreur template", 500)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Erreur Execute Profil:", err)
	}
}
