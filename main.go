package main

import (
	"fmt"
	"html"
	"net/http"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func servIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func newServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("New incoming connection from client:", ws.RemoteAddr())
	s.conns[ws] = true
	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {

	for {
		var packet map[string]interface{}
		websocket.JSON.Receive(ws, &packet)
		message := string(packet["message_input"].(string))
		formatted_message := fmt.Sprintf(`<div id="message" hx-swap-oob="beforeend">
    	<p>%s</p></div>`, html.EscapeString(message))
		go s.broadcast([]byte(formatted_message))
	}

}

func (s *Server) broadcast(b []byte) {
	for ws := range s.conns {
		func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("Write Error", err)
			}
		}(ws)
	}

}

func main() {
	server := newServer()
	http.HandleFunc("/", servIndex)
	http.Handle("/ws", websocket.Handler(server.handleWS))

	http.ListenAndServe(":1337", nil)
}
