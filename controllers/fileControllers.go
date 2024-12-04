package controllers

import (
	"net/http"
	"thewire/internal/auth"
)

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	// Delegate all this to middleware
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sessionToken := cookie.Value

	userSession, exists := auth.Sessions[sessionToken]
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if userSession.IsExpired() {
		delete(auth.Sessions, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	http.ServeFile(w, r, "../templates/index.html")
}

func ServeLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../templates/login.html")

}

func ServeRegister(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../templates/register.html")
}
