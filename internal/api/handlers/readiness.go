package handlers

import (
	"net/http"
)

func (h *handler) Readiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
