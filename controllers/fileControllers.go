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
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sessionToken := cookie.Value

	userSession, exists := auth.Sessions[sessionToken]
	if !exists {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if userSession.IsExpired() {
		delete(auth.Sessions, sessionToken)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
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
