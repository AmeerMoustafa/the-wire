package main

import (
	"net/http"
	websocketServer "thewire/controllers"

	"golang.org/x/net/websocket"
)

func servIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../templates/index.html")
}

func servLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../templates/login.html")

}

func main() {
	WSServer := websocketServer.NewServer()
	http.HandleFunc("/", servIndex)
	http.Handle("/ws", websocket.Handler(WSServer.HandleWS))
	http.HandleFunc("/login", servLogin)
	http.ListenAndServe(":1337", nil)
}
