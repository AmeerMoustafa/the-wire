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
	Router.HandleFunc("GET /login", controllers.ServeLogin)
	Router.HandleFunc("POST /login", controllers.LoginUser)
	Router.Handle("/ws", websocket.Handler(WSServer.HandleWS))

}
