package main

import (
	"net/http"

	"github.com/ascribner/sustainabilityapp/services/user"
)

type Server struct {
	host    string
	port    string
	handler *http.ServeMux
}

// Creates a new instance of the server struct
func NewServer(host string, port string, handler *http.ServeMux) *Server {
	if host == "" {
		host = "127.0.0.1"
	}

	if port == "" {
		port = "8080"
	}

	if handler == nil {
		handler = http.NewServeMux()
	}

	return &Server{
		host:    host,
		port:    port,
		handler: handler,
	}
}

// Configures the routes and starts the server
func (s *Server) Run() {
	var address = s.host + ":" + s.port

	//Routes
	userHandler := user.NewUserHandler()
	userHandler.RegisterRoutes(s.handler)

	// Start
	http.ListenAndServe(address, s.handler)
}
