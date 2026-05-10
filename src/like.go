package src

import (
	"net/http"
)

func LikePost(w http.ResponseWriter, r *http.Request) {
	userID := getUser(r)
	if userID == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	postID := r.FormValue("post_id")

	var exists int
	db.QueryRow("SELECT COUNT(*) FROM likes WHERE user_id = ? AND post_id = ?", userID, postID).Scan(&exists)

	if exists > 0 {
		db.Exec("DELETE FROM likes WHERE user_id = ? AND post_id = ?", userID, postID)
	} else {
		db.Exec("INSERT INTO likes (user_id, post_id) VALUES (?, ?)", userID, postID)
	}

	http.Redirect(w, r, "/post/"+postID, http.StatusSeeOther)
}

func CommentPost(w http.ResponseWriter, r *http.Request) {
	userID := getUser(r)
	if userID == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	postID := r.FormValue("post_id")
	content := r.FormValue("content")

	if content != "" {
		db.Exec("INSERT INTO comments (user_id, post_id, content) VALUES (?, ?, ?)", userID, postID, content)
	}

	http.Redirect(w, r, "/post/"+postID, http.StatusSeeOther)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	userID := getUser(r)
	if userID == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	commentID := r.FormValue("comment_id")
	postID := r.FormValue("post_id")
	db.Exec("DELETE FROM comments WHERE id = ? AND user_id = ?", commentID, userID)
	http.Redirect(w, r, "/post/"+postID, http.StatusSeeOther)
}
