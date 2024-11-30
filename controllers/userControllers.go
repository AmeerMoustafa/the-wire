package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"thewire/internal/database"

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

	db := database.Connect()

	results, err := db.Exec("INSERT into users (username, password_hash) VALUES (?, ?)", user.Username, hashed_password)

	if err != nil {
		log.Fatal(err)
	}

	id, _ := results.LastInsertId()

	w.Header().Set("content-type", "application/json")

	// Return real JSON here
	fmt.Println("Created user with the ID: ", id)

	db.Close()

}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user User
	var returned_user User
	decoder := json.NewDecoder(r.Body)

	decoder.Decode(&user)

	db := database.Connect()

	row, err := db.Query("SELECT * from users WHERE username = ?", user.Username)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	for row.Next() {
		row.Scan(&returned_user.id, &returned_user.Username, &returned_user.Password)
	}

	result := bcrypt.CompareHashAndPassword([]byte(returned_user.Password), []byte(user.Password))

	fmt.Println(returned_user.Password)
	fmt.Println(returned_user)
	fmt.Println(result)

}
