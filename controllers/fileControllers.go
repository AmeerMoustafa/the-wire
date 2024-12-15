package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"thewire/internal/auth"
	templates "thewire/view"
)

func ServeIndex(w http.ResponseWriter, r *http.Request) {

	index_file, err := template.ParseFS(templates.Templates, "templates/index.html")

	if err != nil {
		fmt.Println(err)
	}

	// Authenticate user
	isAuthenticated := auth.AuthenticateUser(w, r)

	if !isAuthenticated {
		return
	}

	if err != nil {
		fmt.Println(err)
	}
	index_file.Execute(w, "")
}

func ServeLogin(w http.ResponseWriter, r *http.Request) {
	login_file, err := template.ParseFS(templates.Templates, "templates/login.html")
	if err != nil {
		log.Fatal(err)
	}

	login_file.Execute(w, "")
}

func ServeRegister(w http.ResponseWriter, r *http.Request) {
	register_file, err := template.ParseFS(templates.Templates, "templates/register.html")
	if err != nil {
		log.Fatal(err)
	}

	register_file.Execute(w, "")
}
