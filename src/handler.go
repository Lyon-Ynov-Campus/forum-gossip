package src

import (
	"html/template"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Home(w http.ResponseWriter, r *http.Request) {
	id := getUser(r)
	msg := r.URL.Query().Get("msg")
	tmpl, err := template.ParseFiles("templates/pageAccueil.html")
	if err != nil {
		log.Fatal(err)
	}
	var myAvatar string
	if id != 0 {
		var avatar *string
		db.QueryRow("SELECT avatar FROM users WHERE id = ?", id).Scan(&avatar)
		if avatar != nil {
			myAvatar = *avatar
		}
	}

	rows, err := db.Query("SELECT username, avatar FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	type User struct {
		Username string
		Avatar   string
	}
	var users []User

	for rows.Next() {
		var u User
		var avatar *string
		rows.Scan(&u.Username, &avatar)
		if avatar != nil {
			u.Avatar = *avatar
		}
		users = append(users, u)
	}

	data := map[string]interface{}{
		"UserID":   id,
		"Message":  msg,
		"Users":    users,
		"MyAvatar": myAvatar,
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
	avatar := r.FormValue("avatar")

	if username != "" {
		db.Exec("UPDATE users SET username = ? WHERE id = ?", username, id)
	}
	if email != "" {
		db.Exec("UPDATE users SET email = ? WHERE id = ?", email, id)
	}
	if password != "" {
		ok, msg := isValidPassword(password)
		if !ok {
			http.Error(w, msg, 400)
			return
		}
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		db.Exec("UPDATE users SET password = ? WHERE id = ?", string(hash), id)
	}
	if avatar != "" {
		db.Exec("UPDATE users SET avatar = ? WHERE id = ?", avatar, id)
	}
	http.Redirect(w, r, "/profil", http.StatusSeeOther)
}
