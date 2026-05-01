package src

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	id := getUser(r)
	tmpl, err := template.ParseFiles("templates/pageAccueil.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, id)
}

//func Login(w http.ResponseWriter, r *http.Request) {
//	tmpl, err := template.ParseFiles("templates/pageConnexion.html")
//	if err != nil {
//		log.Fatal(err)
//	}
//	tmpl.Execute(w, nil)

//}
