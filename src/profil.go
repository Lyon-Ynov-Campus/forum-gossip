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

	tab := r.URL.Query().Get("tab")
	if tab == "" {
		tab = "Publications"
	}

	type Comment struct {
		ID              int
		PostTitle       string
		Content         string
		PublicationDate string
	}

	var comments []Comment

	if tab == "Commentaires" {
		rows2, err := db.Query(`
		SELECT comments.id, posts.title, comments.content, comments.publication_date
		FROM comments
		JOIN posts ON comments.post_id = posts.id
		WHERE comments.user_id = ?
		ORDER BY comments.publication_date DESC
	`, id)
		if err != nil {
			http.Error(w, "Erreur récupération commentaires", 500)
			return
		}
		defer rows2.Close()
		for rows2.Next() {
			var c Comment
			rows2.Scan(&c.ID, &c.PostTitle, &c.Content, &c.PublicationDate)
			comments = append(comments, c)
		}
	}

	type Like struct {
		ID      int
		Title   string
		Content string
	}
	rows3, err := db.Query(`
		SELECT posts.id, posts.title, posts.content
		FROM likes
		JOIN posts ON likes.post_id = posts.id
		WHERE likes.user_id = ?
		ORDER BY posts.publication_date DESC
	`, id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Erreur récupération likes", 500)
		return
	}
	defer rows3.Close()

	var likes []Like
	for rows.Next() {
		var l Like
		rows.Scan(&l.ID, &l.Title, &l.Content)
		likes = append(likes, l)
	}

	data := map[string]interface{}{
		"Username":  username,
		"Email":     email,
		"Avatar":    avatar,
		"Posts":     posts,
		"ActiveTab": tab,
		"Comments":  comments,
		"Likes":     likes,
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
