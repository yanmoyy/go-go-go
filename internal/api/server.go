package api

import (
	"net/http"
	"time"

	"github.com/yanmoyy/go-go-go/internal/api/handler"
	"github.com/yanmoyy/go-go-go/internal/config"
	"github.com/yanmoyy/go-go-go/internal/database"
)

const (
	// Auth endpoints
	endpointLogin  = "/api/auth/login"
	endpointSignup = "/api/auth/signup"
	endpointLogout = "/api/auth/logout"

	// Game endpoints
	endpointRooms         = "/api/rooms"
	endpointCreateRoom    = "/api/rooms/create"
	endpointGameBoard     = "/api/game/board"
	endpointGameMove      = "/api/game/move"
	endpointGameSurrender = "/api/game/surrender"
	endpointGameLeave     = "/api/game/leave"

	// Health check
	endpointHealthz = "/healthz"
)

type Server struct {
	cfg    *config.Config
	db     *database.Queries
	router *router
}

func NewServer(cfg *config.Config, db *database.Queries) *Server {
	return &Server{
		cfg:    cfg,
		db:     db,
		router: newRouter(),
	}
}

func (s *Server) Start() error {
	s.setupRoutes()

	srv := &http.Server{
		Addr:              ":" + s.cfg.Port,
		ReadHeaderTimeout: time.Second * 5,
		Handler:           s.router.mux,
	}

	return srv.ListenAndServe()
}

func (s *Server) setupRoutes() {
	h := handler.NewHandler(s.db)

	// Serve static files
	s.router.ServeStatic("/", "app")
	// Health check
	s.router.Get(endpointHealthz, h.CheckHealth)

	// Auth routes
	s.router.Post(endpointLogin, h.Login)
	s.router.Post(endpointSignup, h.Signup)
	s.router.Post(endpointLogout, h.Logout)
}
