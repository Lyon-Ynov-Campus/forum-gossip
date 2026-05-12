package src

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Posts(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		id := getUser(r)
		if id == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		tmpl, err := template.ParseFiles("templates/posts.html")
		if err != nil {
			http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
			return
		}

		data := map[string]interface{}{
			"UserID": id,
		}

		tmpl.Execute(w, data)
		return
	}
	if r.Method == "POST" {
		id := getUser(r)
		if id == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		title := r.FormValue("title")
		content := r.FormValue("content")
		//postType := r.FormValue("type")
		//post := title + content
		//Body, _ := os.ReadFile(post)
		_, err := db.Exec(
			"INSERT INTO posts (user_id, title, content) VALUES (?, ?, ?)",
			id, title, content,
		)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		fmt.Printf("Nouveau post reçu : id user: %d | Titre: %s | Contenu: %s\n", id, title, content)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func EditPost(w http.ResponseWriter, r *http.Request) {
	id := getUser(r)
	if id == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	postID := r.URL.Query().Get("id")

	type Post struct {
		ID      int
		Title   string
		Content string
	}

	var p Post
	db.QueryRow("SELECT id, title, content FROM posts WHERE id = ? AND user_id = ?", postID, id).Scan(&p.ID, &p.Title, &p.Content)

	tmpl, err := template.ParseFiles("templates/edit.html")
	if err != nil {
		http.Error(w, "Erreur template", 500)
		return
	}
	err = tmpl.Execute(w, p)
	if err != nil {
		log.Println("Erreur Execute edit:", err)
	}

}
