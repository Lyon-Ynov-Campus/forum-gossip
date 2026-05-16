package src

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

func PostDetail(w http.ResponseWriter, r *http.Request) {
	userID := getUser(r)
	postID := strings.TrimPrefix(r.URL.Path, "/post/")

	if postID == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	type Post struct {
		ID       int
		Title    string
		Content  string
		Username string
		Avatar   string
		NbLikes  int
		Liked    bool
		Date     string
	}

	type Comment struct {
		ID       int
		Username string
		Content  string
		Date     string
		IsOwner  bool
	}

	var post Post
	var avatar *string
	err := db.QueryRow(`
		SELECT p.id, p.title, p.content, u.username, u.avatar, p.publication_date,
			COUNT(DISTINCT l.id)
		FROM posts p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN likes l ON l.post_id = p.id
		WHERE p.id = ?
		GROUP BY p.id
	`, postID).Scan(&post.ID, &post.Title, &post.Content, &post.Username, &avatar, &post.Date, &post.NbLikes)
	if err != nil {
		http.Error(w, "Post introuvable", 404)
		return
	}

	if avatar != nil {
		post.Avatar = *avatar
	} else {
		post.Avatar = "/static/default.png"
	}

	if userID != 0 {
		var liked int
		db.QueryRow("SELECT COUNT(*) FROM likes WHERE user_id = ? AND post_id = ?", userID, postID).Scan(&liked)
		post.Liked = liked > 0
	}

	rows, err := db.Query(`
	SELECT c.id, u.username, c.content, c.publication_date, c.user_id
	FROM comments c
	JOIN users u ON c.user_id = u.id
	WHERE c.post_id = ?
	ORDER BY c.publication_date ASC
	`, postID)
	if err != nil {
		http.Error(w, "Erreur commentaires", 500)
		return
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var c Comment
		var commentUserID int
		rows.Scan(&c.ID, &c.Username, &c.Content, &c.Date, &commentUserID)
		c.IsOwner = commentUserID == userID
		comments = append(comments, c)
	}

	tmpl, err := template.ParseFiles("templates/postDetail.html")
	if err != nil {
		http.Error(w, "Erreur template", 500)
		return
	}
	err = tmpl.Execute(w, map[string]interface{}{
		"Post":     post,
		"Comments": comments,
		"UserID":   userID,
		"PostID":   postID,
	})
	if err != nil {
		log.Println("Erreur Execute PostDetail:", err)
	}
}
