package src

import (
	"html/template"
	"log"
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	id := getUser(r)
	q := r.URL.Query().Get("q")
	type UserResult struct {
		Username   string
		Avatar     string
		NbPosts    int
		NbLikes    int
		NbComments int
	}
	var users []UserResult
	rows, err := db.Query(`
		SELECT u.username, IFNULL(u.avatar, '/static/default.png'),
			COUNT(DISTINCT p.id) AS nb_posts,
			COUNT(DISTINCT l.id) AS nb_likes,
			COUNT(DISTINCT c.id) AS nb_comments
		FROM users u
		LEFT JOIN posts p ON p.user_id = u.id
		LEFT JOIN likes l ON l.post_id = p.id
		LEFT JOIN comments c ON c.post_id = p.id
		WHERE u.username LIKE ?
		GROUP BY u.id
	`, "%"+q+"%")
	if err != nil {
		http.Error(w, "Erreur recherche", 500)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u UserResult
		rows.Scan(&u.Username, &u.Avatar, &u.NbPosts, &u.NbLikes, &u.NbComments)
		users = append(users, u)
	}

	tmpl, err := template.ParseFiles("templates/search.html")
	if err != nil {
		http.Error(w, "Erreur template", 500)
		return
	}
	err = tmpl.Execute(w, map[string]interface{}{
		"Query":  q,
		"Users":  users,
		"UserID": id,
	})
	if err != nil {
		log.Println("Erreur Execute Search:", err)
	}
}
