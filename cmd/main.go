package main

import (
	"net/http"
	"thewire/routes"
)

var server = http.Server{
	Addr:    ":1337",
	Handler: routes.Router,
}

func main() {

	server.ListenAndServe()
}
