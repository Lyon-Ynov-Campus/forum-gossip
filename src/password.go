package src

import (
	"fmt"
	"html/template"
	"net/http"
	"net/smtp"
	"golang.org/x/crypto/bcrypt"
)

func sendMail(to string, link string) {
	from := "avril.brn.gnzz@gmail.com"
	password := "irkl ngil xdzq vbbl"
	msg := "Subject: Reset password\n\nClique ici pour changer ton mot de passe : " + link
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		smtp.PlainAuth("", from, password, "smtp.gmail.com"),
		from,
		[]string{to},
		[]byte(msg),
	)

	if err != nil {
		fmt.Println("Erreur mail:", err)
	} else {
		fmt.Println("Mail envoyé !")
	}
}

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/forgot.html")
		t.Execute(w, nil)
		return
	}

	email := r.FormValue("email")
	link := "http://localhost:8080/reset?email=" + email

	sendMail(email, link)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		email := r.URL.Query().Get("email")

		t, _ := template.ParseFiles("templates/reset.html")
		t.Execute(w, map[string]string{"Email": email})
		return
	}

	email := r.FormValue("email")
	newPass := r.FormValue("password")
	hash, _ := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost)
	db.Exec("UPDATE users SET password = ? WHERE email = ?", string(hash), email)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
