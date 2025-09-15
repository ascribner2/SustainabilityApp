package main

import (
	"database/sql"
	"net/http"

	"github.com/ascribner/sustainabilityapp/internal/handlers/auth"
	"github.com/ascribner/sustainabilityapp/internal/handlers/item"
	"github.com/ascribner/sustainabilityapp/internal/handlers/user"
	"github.com/ascribner/sustainabilityapp/internal/repos"
	"github.com/ascribner/sustainabilityapp/internal/services"
)

type Server struct {
	host    string
	port    string
	handler *http.ServeMux
	db      *sql.DB
}

// Creates a new instance of the server struct
func NewServer(host string, port string, handler *http.ServeMux, db *sql.DB) *Server {
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
		db:      db,
	}
}

// Configures the routes and starts the server
func (s *Server) Run() error {
	var address = s.host + ":" + s.port

	//Routes
	userRepo := repos.NewUserRepo(s.db)
	userService := services.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)
	userHandler.RegisterRoutes(s.handler)

	authRepo := repos.NewAuthRepo(s.db)
	authService := services.NewAuthService(authRepo)
	authHandler := auth.NewAuthHandler(authService)
	authHandler.RegisterRoutes(s.handler)

	itemRepo := repos.NewItemRepo(s.db)
	itemService := services.NewItemService(itemRepo)
	itemHandler := item.NewItemHandler(itemService)
	itemHandler.RegisterRoutes(s.handler)

	// Start
	err := http.ListenAndServe(address, s.handler)
	if err != nil {
		return err
	}
	return nil
}
