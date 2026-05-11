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

	type Post struct {
		ID              int
		Title           string
		Content         string
		PublicationDate string
		NbLikes         int
		NbComments      int
	}

	rows, err := db.Query(`
		SELECT p.id, p.title, p.content, p.publication_date,
			COUNT(DISTINCT l.id) AS nb_likes,
			COUNT(DISTINCT c.id) AS nb_comments
		FROM posts p
		LEFT JOIN likes l ON l.post_id = p.id
		LEFT JOIN comments c ON c.post_id = p.id
		WHERE p.user_id = ?
		GROUP BY p.id
		ORDER BY p.publication_date DESC
	`, id)
	if err != nil {
		http.Error(w, "Erreur récupération posts", 500)
		return
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var p Post
		rows.Scan(&p.ID, &p.Title, &p.Content, &p.PublicationDate, &p.NbLikes, &p.NbComments)
		posts = append(posts, p)
	}

	data := map[string]interface{}{
		"Username": username,
		"Email":    email,
		"Avatar":   avatar,
		"Posts":    posts,
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
