package api

import (
	"net/http"
	"time"

	"github.com/yanmoyy/go-go-go/internal/api/handlers"
	"github.com/yanmoyy/go-go-go/internal/config"
	"github.com/yanmoyy/go-go-go/internal/database"
)

const (
	endPointUsers = "/api/users"
)

type Server struct {
	cfg *config.Config
	db  *database.Queries
}

func NewServer(cfg *config.Config, db *database.Queries) *Server {
	return &Server{cfg: cfg, db: db}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	s.setupRoutes(mux)

	srv := &http.Server{
		Addr:              ":" + s.cfg.Port,
		ReadHeaderTimeout: time.Second * 5,
		Handler:           mux,
	}

	return srv.ListenAndServe()
}

func (s *Server) setupRoutes(mux *http.ServeMux) {
	handler := handlers.NewHandler(s.db)
	mux.HandleFunc("POST "+endPointUsers, handler.CreateUser)
}
