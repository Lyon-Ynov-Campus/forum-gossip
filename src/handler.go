package src

import (
	"database/sql"
	"fmt"
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
	type Post struct {
		Id               int
		Title            string
		Content          string
		Username         string
		Avatar           string
		NbComments       int
		NbLikes          int
		Publication_date string
	}

	var posts []Post
	rows2, err := db.Query(`
	SELECT posts.id, posts.title, posts.content, posts.publication_date, users.username, users.avatar,
        COUNT(DISTINCT comments.id) AS nb_comments,
        COUNT(DISTINCT likes.id) AS nb_likes
	FROM posts
	JOIN users ON posts.user_id = users.id
	LEFT JOIN comments ON posts.id = comments.post_id
	LEFT JOIN likes ON posts.id = likes.post_id
	GROUP BY posts.id
	ORDER BY posts.publication_date DESC
	`)
	if err != nil {
		fmt.Println("erreur ajout post", err)
	}
	if err == nil {
		defer rows2.Close()
		for rows2.Next() {
			var p Post
			var avatar *string
			rows2.Scan(&p.Id, &p.Title, &p.Content, &p.Publication_date, &p.Username, &avatar, &p.NbComments, &p.NbLikes)

			if avatar != nil {
				p.Avatar = *avatar
			} /*else {
			p.avatar = "static.png"}*/

			posts = append(posts, p)
		}

	}
	data := map[string]interface{}{
		"UserID":   id,
		"Message":  msg,
		"Users":    users,
		"Posts":    posts,
		"MyAvatar": myAvatar,
	}
	tmpl.Execute(w, data)
}

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
	data := map[string]interface{}{}

	if username != "" {
		db.Exec("UPDATE users SET username = ? WHERE id = ?", username, id)
	}
	if email != "" {
		db.Exec("UPDATE users SET email = ? WHERE id = ?", email, id)
	}
	if password != "" {
		ok, msg := isValidPassword(password)
		if !ok {
			data["Error"] = msg
		} else {
			hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			db.Exec("UPDATE users SET password = ? WHERE id = ?", string(hash), id)
			data["Success"] = "Mot de passe modifié"
		}
	}
	if avatar != "" {
		db.Exec("UPDATE users SET avatar = ? WHERE id = ?", avatar, id)
	}

	var user struct {
		Username string
		Email    string
		Avatar   string
	}
	var avatarNull sql.NullString
	db.QueryRow("SELECT username, email, avatar FROM users WHERE id = ?", id).Scan(&user.Username, &user.Email, &avatarNull)
	if avatarNull.Valid && avatarNull.String != "" {
		user.Avatar = avatarNull.String
	} else {
		user.Avatar = "/static/default.png"
	}
	data["Username"] = user.Username
	data["Email"] = user.Email
	data["Avatar"] = user.Avatar

	tmpl, err := template.ParseFiles("templates/pageUtilisateur.html")
	if err != nil {
		http.Error(w, "Erreur template", 500)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Erreur Execute UpdateUser:", err)
	}
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	id := getUser(r)
	if id == 0 {
		http.Redirect(w, r, "login", http.StatusSeeOther)
		return
	}
	title := r.FormValue("title")
	content := r.FormValue("content")
	publication_date := r.FormValue("publication_date")

	if title != "" {
	}
	if content != "" {
	}
	if publication_date != "" {
	}
}
