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
	Router.HandleFunc("/login", controllers.ServeLogin)
	Router.Handle("/ws", websocket.Handler(WSServer.HandleWS))

}
