func LikeAPIHandler(w http.ResponseWriter, r *http.Request) {
	userID := getUser(r)
	if userID == 0 {
		http.Error(w, "Login requis", http.StatusUnauthorized)
		return
	}

	postID := r.URL.Query().Get("id")

	var newCount int
	db.QueryRow("SELECT COUNT(*) FROM likes WHERE post_id = ?", postID).Scan(&newCount)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"likesCount": %d}`, newCount)
}