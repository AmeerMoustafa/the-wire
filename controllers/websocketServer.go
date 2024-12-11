package controllers

import (
	"fmt"
	"html"
	"thewire/internal/auth"

	"golang.org/x/net/websocket"
)

type WSServer struct {
	conns map[*websocket.Conn]bool
	users map[string]string
}

func NewServer() *WSServer {
	return &WSServer{
		conns: make(map[*websocket.Conn]bool),
		users: make(map[string]string),
	}
}

func (s *WSServer) HandleWS(ws *websocket.Conn) {
	cookie, err := ws.Request().Cookie("session_token")
	if err != nil {
		fmt.Println("Cookie not found")
		ws.Request().Header.Set("HX-Redirect", "/")
		return

	}
	fmt.Println("New incoming connection from client:", ws.RemoteAddr())
	username := auth.Sessions[cookie.Value].Username
	s.users[cookie.Value] = username
	s.conns[ws] = true
	s.readLoop(ws)

}

func (s *WSServer) readLoop(ws *websocket.Conn) {
	for {
		var packet map[string]interface{}
		websocket.JSON.Receive(ws, &packet)
		if len(packet) > 0 {
			cookie, err := ws.Request().Cookie("session_token")
			if err != nil {
				fmt.Println("Cookie not found, closing")
				ws.Request().Header.Set("HX-Redirect", "/")
				return
			}
			username := s.users[cookie.Value]
			message := string(packet["message_input"].(string))
			formatted_message := fmt.Sprintf(`<div id="message" hx-swap-oob="beforeend">
    	<p><span>%s: </span>%s</p></div>`, html.EscapeString(username), html.EscapeString(message))
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
