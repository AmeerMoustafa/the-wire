package routes

import (
	"net/http"
	"thewire/controllers"

	"golang.org/x/net/websocket"
)

var Router = http.NewServeMux()

func init() {
	WSServer := controllers.NewServer()

	Router.HandleFunc("/", controllers.ServeIndex)
	// Login Routes
	Router.HandleFunc("GET /login", controllers.ServeLogin)
	Router.HandleFunc("POST /login", controllers.Login)

	// Registration Routes
	Router.HandleFunc("GET /register", controllers.ServeRegister)
	Router.HandleFunc("POST /register", controllers.CreateUser)

	// Logout Route
	Router.HandleFunc("POST /logout", controllers.Logout)

	Router.Handle("/ws", websocket.Handler(WSServer.HandleWS))

}
