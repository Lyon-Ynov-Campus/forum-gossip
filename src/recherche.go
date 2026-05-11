package src

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

func Search(w http.ResponseWriter, r *http.Request) {
	id := getUser(r)
	q := r.URL.Query().Get("q")
	searchType := r.URL.Query().Get("type")
	if searchType == "" {
		searchType = "users"
	}
	sort := r.URL.Query().Get("sort")
	order := r.URL.Query().Get("order")
	type UserResult struct {
		Username   string
		Avatar     string
		NbPosts    int
		NbLikes    int
		NbComments int
	}

	type PostResult struct {
		ID              int
		Title           string
		Content         string
		Username        string
		Avatar          string
		PublicationDate string
		NbLikes         int
		NbComments      int
	}

	var users []UserResult
	var posts []PostResult

	if searchType == "users" {

		rows, err := db.Query(`
			SELECT 
				u.username,
				IFNULL(u.avatar, '/static/default.png'),
				COUNT(DISTINCT p.id),
				COUNT(DISTINCT l.id),
				COUNT(DISTINCT c.id)
			FROM users u
			LEFT JOIN posts p ON p.user_id = u.id
			LEFT JOIN likes l ON l.post_id = p.id
			LEFT JOIN comments c ON c.post_id = p.id
			WHERE u.username LIKE ?
			GROUP BY u.id
		`, "%"+q+"%")

		if err != nil {
			http.Error(w, "Erreur recherche users", 500)
			return
		}
		defer rows.Close()
		for rows.Next() {
			var u UserResult
			rows.Scan(
				&u.Username,
				&u.Avatar,
				&u.NbPosts,
				&u.NbLikes,
				&u.NbComments,
			)
			users = append(users, u)
		}
	}
	if searchType == "posts" {
		query := `
			SELECT 
				p.id,
				p.title,
				p.content,
				u.username,
				IFNULL(u.avatar, '/static/default.png'),
				p.publication_date,
				COUNT(DISTINCT l.id),
				COUNT(DISTINCT c.id)
			FROM posts p
			JOIN users u ON p.user_id = u.id
			LEFT JOIN likes l ON l.post_id = p.id
			LEFT JOIN comments c ON c.post_id = p.id
			WHERE p.title LIKE ?
			GROUP BY p.id
		`
		if sort == "likes" {
			query += " ORDER BY COUNT(DISTINCT l.id) "
		} else if sort == "comments" {
			query += " ORDER BY COUNT(DISTINCT c.id) "
		} else {
			query += " ORDER BY p.publication_date "
		}
		if order == "asc" {
			query += "ASC"
		} else {
			query += "DESC"
		}

		rows, err := db.Query(query, "%"+q+"%")

		if err != nil {
			http.Error(w, "Erreur recherche posts", 500)
			return
		}

		defer rows.Close()
		for rows.Next() {
			var p PostResult
			rows.Scan(
				&p.ID,
				&p.Title,
				&p.Content,
				&p.Username,
				&p.Avatar,
				&p.PublicationDate,
				&p.NbLikes,
				&p.NbComments,
			)
			posts = append(posts, p)
		}
	}

	tmpl, err := template.ParseFiles("templates/search.html")
	if err != nil {
		http.Error(w, "Erreur template", 500)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{
		"Query":  q,
		"Users":  users,
		"Posts":  posts,
		"Type":   searchType,
		"Sort":   sort,
		"Order":  order,
		"UserID": id,
	})

	if err != nil {
		log.Println("Erreur Execute Search:", err)
	}
}

func UserProfil(w http.ResponseWriter, r *http.Request) {
	username := strings.TrimPrefix(r.URL.Path, "/user/")
	if username == "" {
		http.Redirect(w, r, "/search", http.StatusSeeOther)
		return
	}

	type Post struct {
		ID              int
		Title           string
		Content         string
		PublicationDate string
		NbLikes         int
		NbComments      int
	}

	var avatar string
	err := db.QueryRow(
		"SELECT IFNULL(avatar, '/static/default.png') FROM users WHERE username = ?",
		username,
	).Scan(&avatar)
	if err != nil {
		http.Error(w, "Utilisateur introuvable", 404)
		return
	}

	rows, err := db.Query(`
		SELECT 
			p.id,
			p.title,
			p.content,
			p.publication_date,
			COUNT(DISTINCT l.id),
			COUNT(DISTINCT c.id)
		FROM posts p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN likes l ON l.post_id = p.id
		LEFT JOIN comments c ON c.post_id = p.id
		WHERE u.username = ?
		GROUP BY p.id
		ORDER BY p.publication_date DESC
	`, username)
	if err != nil {
		http.Error(w, "Erreur récupération posts", 500)
		return
	}
	defer rows.Close()
	var posts []Post
	for rows.Next() {
		var p Post
		rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.PublicationDate,
			&p.NbLikes,
			&p.NbComments,
		)

		posts = append(posts, p)
	}

	tmpl, err := template.ParseFiles("templates/userProfil.html")
	if err != nil {
		http.Error(w, "Erreur template", 500)
		return
	}
	err = tmpl.Execute(w, map[string]interface{}{
		"Username": username,
		"Avatar":   avatar,
		"Posts":    posts,
	})
	if err != nil {
		log.Println("Erreur Execute UserProfil:", err)
	}
}
