package src

import (
	"fmt"
	"html/template"
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

		fmt.Printf("Nouveau post reçu : [%v\n], %s,", "| Contenu : %s\n" /*postType, title,*/, content, id, title)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func LoadPost(w http.ResponseWriter, r *http.Request) {
	id := getUser(r)
	if id == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

}

func ChangePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := getUser(r)
		if id == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
	}

}

func DeletePost(w http.ResponseWriter, r *http.Request) {

}
