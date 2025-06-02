package handlers

import "github.com/yanmoyy/go-go-go/internal/database"

type handler struct {
	db *database.Queries
}

func NewHandler(db *database.Queries) *handler {
	return &handler{db: db}
}
