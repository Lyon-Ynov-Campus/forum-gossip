package src

import (
	"fmt"
	"html/template"
	"net/http"
)

func Posts(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("templates/share.html")
		if err != nil {
			http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
			return
		}

		data := map[string]interface{}{
			"UserID": 1,
		}

		tmpl.Execute(w, data)
		return
	}
	if r.Method == "POST" {

		title := r.FormValue("title")
		content := r.FormValue("content")
		postType := r.FormValue("type")

		fmt.Printf("Nouveau post reçu : [%s] %s | Contenu : %s\n", postType, title, content)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func ChangePost(w http.ResponseWriter, r *http.Request) {

}

func DeletePost(w http.ResponseWriter, r *http.Request) {

}
