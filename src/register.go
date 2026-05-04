package src

import (
	"html/template"
	"log"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	type UserRegister struct {
		Username string
		Email    string
		Avatar   string
		Password string
	}

	register := UserRegister{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Avatar:   r.FormValue("avatar"),
		Password: r.FormValue("password"),
	}

	if string.TrimSpace(register.Username) == "" {
		rv.Error["Username"] = "Le nom ne peut pas etre vide"
	}
	if string.TrimSpace(register.Email) == "" {
		rv.Error["Email"] = "L'email ne peut pas etre vide"
	}
	if string.TrimSpace(register.Password) == "" {
		rv.Error["Password"] = "Le mot de passe ne peut pas etre vide"
	}

	_ = register
	tmpl, err := template.ParseFiles("templates/register.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)

}
