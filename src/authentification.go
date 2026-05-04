package src

import (
	"fmt"
	"math/rand"
	"net/http"
)

var sessions = map[string]int{}

func addSession(w http.ResponseWriter, userID int) {
	token := fmt.Sprint(rand.Int())
	sessions[token] = userID

	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: token,
		Path:  "/",
	})
}

func getUser(r *http.Request) int {
	c, err := r.Cookie("session")
	if err != nil {
		return 0
	}
	id, ok := sessions[c.Value]
	if !ok {
		return 0
	}

	return id
}

func removeSession(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err == nil {
		delete(sessions, c.Value)
	}

	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}
