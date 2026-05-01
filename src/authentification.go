package src

import (
	"net/http"
)

var connectedUserID int = 0

func addSession(w http.ResponseWriter, userID int) {
	connectedUserID = userID

	http.SetCookie(w, &http.Cookie{
		Name:  "login",
		Value: "true",
		Path:  "/",
	})
}

func getUser(r *http.Request) int {
	c, err := r.Cookie("login")
	if err != nil || c.Value != "true" {
		return 0
	}

	return connectedUserID
}

func removeSession(w http.ResponseWriter, r *http.Request) {
	connectedUserID = 0

	http.SetCookie(w, &http.Cookie{
		Name:   "login",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}
