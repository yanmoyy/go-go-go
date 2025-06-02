package handler

// import (
// 	"encoding/json"
// 	"net/http"

// 	"github.com/yanmoyy/go-go-go/internal/api/model"
// )

// func (h *Handler) GetRooms(w http.ResponseWriter, r *http.Request) {
// 	type response struct {
// 		Rooms []model.Room `json:"rooms"`
// 	}

// 	rooms, err := h.db.GetRooms(r.Context())
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, "Couldn't get rooms", err)
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, response{
// 		Rooms: rooms,
// 	})
// }

// func (h *Handler) CreateRoom(w http.ResponseWriter, r *http.Request) {
// 	type parameters struct {
// 		Name string `json:"name"`
// 	}
// 	type response struct {
// 		Room model.Room `json:"room"`
// 	}

// 	decoder := json.NewDecoder(r.Body)
// 	params := parameters{}
// 	err := decoder.Decode(&params)
// 	if err != nil {
// 		respondWithParameterError(w, err)
// 		return
// 	}

// 	room, err := h.db.CreateRoom(r.Context(), params.Name)
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, "Couldn't create room", err)
// 		return
// 	}

// 	respondWithJSON(w, http.StatusCreated, response{
// 		Room: room,
// 	})
// }

// func (h *Handler) GetBoard(w http.ResponseWriter, r *http.Request) {
// 	type response struct {
// 		Board model.Game `json:"board"`
// 	}

// 	gameID := r.URL.Query().Get("game_id")
// 	if gameID == "" {
// 		respondWithError(w, http.StatusBadRequest, "Game ID is required", nil)
// 		return
// 	}

// 	game, err := h.db.GetGame(r.Context(), gameID)
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, "Couldn't get game", err)
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, response{
// 		Board: game,
// 	})
// }

// func (h *Handler) MakeMove(w http.ResponseWriter, r *http.Request) {
// 	type parameters struct {
// 		GameID string `json:"game_id"`
// 		X      int    `json:"x"`
// 		Y      int    `json:"y"`
// 	}
// 	type response struct {
// 		Game model.Game `json:"game"`
// 	}

// 	decoder := json.NewDecoder(r.Body)
// 	params := parameters{}
// 	err := decoder.Decode(&params)
// 	if err != nil {
// 		respondWithParameterError(w, err)
// 		return
// 	}

// 	// TODO: Implement 알까기 game logic
// 	game, err := h.db.MakeMove(r.Context(), params.GameID, params.X, params.Y)
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, "Couldn't make move", err)
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, response{
// 		Game: game,
// 	})
// }

// func (h *Handler) Surrender(w http.ResponseWriter, r *http.Request) {
// 	type parameters struct {
// 		GameID string `json:"game_id"`
// 	}

// 	decoder := json.NewDecoder(r.Body)
// 	params := parameters{}
// 	err := decoder.Decode(&params)
// 	if err != nil {
// 		respondWithParameterError(w, err)
// 		return
// 	}

// 	err = h.db.SurrenderGame(r.Context(), params.GameID)
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, "Couldn't surrender game", err)
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, map[string]string{
// 		"message": "Game surrendered successfully",
// 	})
// }

// func (h *Handler) LeaveGame(w http.ResponseWriter, r *http.Request) {
// 	type parameters struct {
// 		GameID string `json:"game_id"`
// 	}

// 	decoder := json.NewDecoder(r.Body)
// 	params := parameters{}
// 	err := decoder.Decode(&params)
// 	if err != nil {
// 		respondWithParameterError(w, err)
// 		return
// 	}

// 	err = h.db.LeaveGame(r.Context(), params.GameID)
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, "Couldn't leave game", err)
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, map[string]string{
// 		"message": "Left game successfully",
// 	})
// }
