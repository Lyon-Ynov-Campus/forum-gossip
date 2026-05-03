package src

import (
	"html/template"
	"net/http"
)

func Profil(w http.ResponseWriter, r *http.Request) {
	id := getUser(r)
	if id == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	tmpl, _ := template.ParseFiles("templates/pageUtilisateur.html")
	tmpl.Execute(w, nil)
}
