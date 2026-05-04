package src

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	id := getUser(r)
	msg := r.URL.Query().Get("msg")
	tmpl, err := template.ParseFiles("templates/pageAccueil.html")
	if err != nil {
		log.Fatal(err)
	}

	data := map[string]interface{}{
		"UserID":  id,
		"Message": msg,
	}
	tmpl.Execute(w, data)
}

//func Login(w http.ResponseWriter, r *http.Request) {
//	tmpl, err := template.ParseFiles("templates/pageConnexion.html")
//	if err != nil {
//		log.Fatal(err)
//	}
//	tmpl.Execute(w, nil)

//}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	id := getUser(r)
	if id == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	db.Exec("DELETE FROM users WHERE id = ?", id)
	removeSession(w, r)
	http.Redirect(w, r, "/?msg=deleted", http.StatusSeeOther)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := getUser(r)
	if id == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	if username != "" {
		db.Exec("UPDATE users SET username = ? WHERE id = ?", username, id)
	}
	if email != "" {
		db.Exec("UPDATE users SET email = ? WHERE id = ?", email, id)
	}
	if password != "" {
		db.Exec("UPDATE users SET password = ? WHERE id = ?", password, id)
	}
	http.Redirect(w, r, "/profil", http.StatusSeeOther)
}
