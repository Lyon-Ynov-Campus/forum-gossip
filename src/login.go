package src

import (
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/login.html")
		t.Execute(w, nil)
		return
	}

	mail := r.FormValue("email")
	pass := r.FormValue("password")
	var id int
	var dbPass string

	err := db.QueryRow(`
		SELECT id, password 
		FROM users 
		WHERE email = ? OR username = ?
	`, mail, mail).Scan(&id, &dbPass)

	if err != nil {
		t, _ := template.ParseFiles("templates/login.html")
		t.Execute(w, map[string]string{"Error": "Email ou mot de passe incorrect"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass))
	if err != nil {
		t, _ := template.ParseFiles("templates/login.html")
		t.Execute(w, map[string]string{"Error": "Email ou mot de passe incorrect"})
		return
	}

	addSession(w, id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	removeSession(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
