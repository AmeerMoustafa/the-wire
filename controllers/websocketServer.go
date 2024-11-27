package controllers

import (
	"fmt"
	"html"

	"golang.org/x/net/websocket"
)

type WSServer struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *WSServer {
	return &WSServer{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *WSServer) HandleWS(ws *websocket.Conn) {
	fmt.Println("New incoming connection from client:", ws.RemoteAddr())
	s.conns[ws] = true
	s.readLoop(ws)
}

func (s *WSServer) readLoop(ws *websocket.Conn) {
	for {
		var packet map[string]interface{}
		websocket.JSON.Receive(ws, &packet)
		if len(packet) > 0 {
			message := string(packet["message_input"].(string))
			formatted_message := fmt.Sprintf(`<div id="message" hx-swap-oob="beforeend">
    	<p>%s</p></div>`, html.EscapeString(message))
			go s.broadcast([]byte(formatted_message))
		}

	}

}

func (s *WSServer) broadcast(b []byte) {
	for ws := range s.conns {
		func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("Write Error", err)
			}
		}(ws)
	}

}
