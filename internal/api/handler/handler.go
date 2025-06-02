package handler

import (
	"github.com/yanmoyy/go-go-go/internal/database"
)

type Handler struct {
	db *database.Queries
}

func NewHandler(db *database.Queries) *Handler {
	return &Handler{db: db}
}
