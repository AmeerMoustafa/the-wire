package controllers

import (
	"net/http"
	"thewire/internal/auth"
)

func ServeIndex(w http.ResponseWriter, r *http.Request) {

	// Authenticate user
	isAuthenticated := auth.AuthenticateUser(w, r)

	if isAuthenticated == false {
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
