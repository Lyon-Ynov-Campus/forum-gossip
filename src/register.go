package src

/*
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

	if req.Method == http.MethodPost {
	rv := &validators.RegisterValidator{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Avatar:   r.FormValue("avatar"),
		Password1: r.FormValue("password2"),
		Password2: r.FormValue("password1"),
	}
	}

	+
	if string.TrimSpace(register.Username) == "" {
		rv.Error["Username"] = "Le nom ne peut pas etre vide"
	}
	if string.TrimSpace(register.Email) == "" {
		rv.Error["Email"] = "L'email ne peut pas etre vide"
	}
	if string.TrimSpace(rv.Password) == "" {
		rv.Error["Password"] = "Le mot de passe ne peut pas etre vide"
	}

	_ = register
	tmpl, err := template.ParseFiles("templates/register.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)

}
*/
