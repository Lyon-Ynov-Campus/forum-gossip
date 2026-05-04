package src

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
=======
	id := getUser(r)
	msg := r.URL.Query().Get("msg")
>>>>>>> 293899e21614191b6f56257d5c90588570820b9e
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
