package src

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	//tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl, err := template.ParseFiles("templates/pageAccueil.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, nil)
}
