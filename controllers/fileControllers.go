package controllers

import "net/http"

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../templates/index.html")
}

func ServeLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../templates/login.html")

}
