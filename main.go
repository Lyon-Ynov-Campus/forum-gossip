package main

import (
	"log"
	"net/http"
)

func main() {
	home := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})

	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
