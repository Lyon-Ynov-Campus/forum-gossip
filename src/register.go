package src

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

type RegisterData struct {
	Error  map[string]string
	Values map[string]string
}

func Register(w http.ResponseWriter, r *http.Request) {

	data := RegisterData{
		Error:  make(map[string]string),
		Values: make(map[string]string),
	}

	var Username, Email, Avatar, password, password_confirmation string

	if r.Method == http.MethodPost {
		Username = r.FormValue("username")
		Email = r.FormValue("email")
		Avatar = r.FormValue("avatar")
		password = r.FormValue("password")
		password_confirmation = r.FormValue("password_confirmation")

		data.Values["Username"] = Username
		data.Values["Email"] = Email

		if strings.TrimSpace(Username) == "" {
			data.Error["Username"] = "Le nom ne peut pas etre vide"
		}
		if strings.TrimSpace(password) == "" || strings.TrimSpace(password_confirmation) == "" {
			data.Error["Password"] = "Les mots de passes ne doivent pas etre vide"
		}
		if password != password_confirmation {
			data.Error["PasswordVerification"] = "Les mots de passe ne correspondent pas"
		}
		if strings.TrimSpace(Email) == "" {
			data.Error["Email"] = "L'email ne peut pas etre vide"
		} else {
			re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
			if !re.MatchString(strings.ToLower(Email)) {
				data.Error["Email"] = "L'email n'est pas valide"
			}
		}
		if Avatar != "" {
			data.Values["Avatar"] = Avatar
		} else {
			data.Values["Avatar"] = "/static/default.png"
		}
		if len(password) < 8 {
			data.Error["PasswordLength"] = "Le mot de passe est trop court"
		} else {
			var Lower, Upper, Digit, Special bool
			for _, char := range password {
				switch {
				case unicode.IsLower(char):
					Lower = true
				case unicode.IsUpper(char):
					Upper = true
				case unicode.IsDigit(char):
					Digit = true
				default:
					Special = unicode.IsPunct(char) || unicode.IsSymbol(char)
				}
			}
			if !Lower || !Upper || !Digit || !Special {
				data.Error["PasswordLength"] = "Le mot de passe doit contenir au moins 8 caractères dont 1 Majuscule, 1 minuscule, 1 chiffre et un caractère spécial"
			}
		}
		if len(data.Error) == 0 {
			var ExistUser string
			err := db.QueryRow("SELECT username FROM users WHERE username = ?", Username).Scan(&ExistUser)
			if err == nil {
				data.Error["SameUsername"] = "Ce nom d'utilisateur est deja pris"
			}
			var ExistEmail string
			err = db.QueryRow("SELECT email FROM users WHERE email = ?", Email).Scan(&ExistEmail)
			if err == nil {
				data.Error["SameEmail"] = "Cet email est deja pris"
			}
			if len(data.Error) == 0 {
				hash, _ := HashPassword(password)
				_, insertErr := db.Exec("INSERT INTO users (username, email, avatar, password) values (?, ?, ?, ?)",
					Username, Email, Avatar, hash)

				if insertErr != nil {
					data.Error["Main"] = "Impossible d'ajouter l'utilisateur dans la db"
					fmt.Println("pas possible d'ajouter ")
					fmt.Println("Erreur insertion:", insertErr)
				} else {
					http.Redirect(w, r, "/login", http.StatusFound)
					return
				}
			}
		}
	}
	tmpl, err := template.ParseFiles("templates/register.html")
	if err != nil {
		fmt.Println("Erreur template:", err)
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

/*
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}*/
