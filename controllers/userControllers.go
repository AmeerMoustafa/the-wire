package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"thewire/internal/auth"
	"thewire/internal/database"
	"time"
	"unicode/utf8"

	"github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id       int
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&user)
	password := []byte(user.Password)

	hashed_password, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	if user.Username == "" && user.Password == "" {
		form_error := fmt.Sprintf(`
		<div
                class="mb-4 p-2 border border-red-500 bg-red-500 bg-opacity-10 text-red-500 flex items-center"
              >
                Are you drunk?, Access denied!
              </div>`)
		w.Write([]byte(form_error))
		return
	}

	// Password length check
	if utf8.RuneCountInString(user.Password) < 8 {
		form_error := fmt.Sprintf(`
		<div
                class="mb-4 p-2 border border-red-500 bg-red-500 bg-opacity-10 text-red-500 flex items-center"
              >
                Password must be at least 8 characters, Access denied!
              </div>`)
		w.Write([]byte(form_error))
		return

	}

	db := database.Connect()

	results, err := db.Exec("INSERT into users (username, password_hash) VALUES (?, ?)", user.Username, hashed_password)

	if err != nil {
		if sqlite3Err, ok := err.(sqlite3.Error); ok {
			if sqlite3Err.Code == sqlite3.ErrConstraint {
				form_error := fmt.Sprintf(`
		<div
                class="mb-4 p-2 border border-red-500 bg-red-500 bg-opacity-10 text-red-500 flex items-center"
              >
                Username already exists, Access denied!
              </div>`)
				w.Write([]byte(form_error))
			}
		}
		return
	}

	id, _ := results.LastInsertId()

	w.WriteHeader(http.StatusCreated)

	fmt.Println("Created user with the ID: ", id)

	db.Close()

}

func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	var returned_user User
	json.NewDecoder(r.Body).Decode(&user)

	if user.Username == "" && user.Password == "" {
		form_error := fmt.Sprintf(`
		<div
                class="mb-4 p-2 border border-red-500 bg-red-500 bg-opacity-10 text-red-500 flex items-center"
              >
                Are you drunk?, Access denied!
              </div>`)
		w.Write([]byte(form_error))
		return
	}

	db := database.Connect()

	row := db.QueryRow("SELECT * from users WHERE username = ?", user.Username)

	row.Scan(&returned_user.id, &returned_user.Username, &returned_user.Password)

	result := bcrypt.CompareHashAndPassword([]byte(returned_user.Password), []byte(user.Password))

	if result == nil {
		// Generating and sending a session cookie
		sessionCookie := auth.GenerateSession(returned_user.Username)
		http.SetCookie(w, &sessionCookie)
		w.Header().Set("HX-Redirect", "/")
		return

	} else {
		form_error := fmt.Sprintf(`
		<div
                class="mb-4 p-2 border border-red-500 bg-red-500 bg-opacity-10 text-red-500 flex items-center"
              >
                Invalid credentials. Access denied.
              </div>`)
		w.Write([]byte(form_error))
		return
	}

}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")

	if err != nil {
		if err == http.ErrNoCookie {
			w.Header().Set("HX-Redirect", "/login")
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sessionToken := cookie.Value

	delete(auth.Sessions, sessionToken)

	// We need to let the client know that the cookie is expired
	// In the response, we set the session token to an empty
	// value and set its expiry as the current time

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now(),
	})

	w.Header().Set("HX-Redirect", "/login")

}
