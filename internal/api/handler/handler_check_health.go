package handler

import (
	"net/http"
)

func (h *Handler) CheckHealth(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
