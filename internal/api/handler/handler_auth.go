package handler

import (
	"encoding/json"
	"net/http"

	"github.com/yanmoyy/go-go-go/internal/api/model"
	"github.com/yanmoyy/go-go-go/internal/database"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	type response struct {
		Token string     `json:"token"`
		User  model.User `json:"user"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithParameterError(w, err)
		return
	}

	user, err := h.db.GetUserByUsername(r.Context(), params.Username)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid credentials", err)
		return
	}

	// // TODO: Add proper password hashing and verification
	// if user.HashedPassword != params.Password {
	// 	respondWithError(w, http.StatusUnauthorized, "Invalid credentials", nil)
	// 	return
	// }

	// TODO: Generate JWT token
	token := "dummy-token"

	respondWithJSON(w, http.StatusOK, response{
		Token: token,
		User:  model.DBUserToUser(user),
	})
}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	type response struct {
		User model.User `json:"user"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithParameterError(w, err)
		return
	}

	// TODO: Add password hashing
	user, err := h.db.CreateUser(r.Context(), database.CreateUserParams{
		params.Username, params.Password,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, response{
		User: model.DBUserToUser(user),
	})
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement token invalidation
	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Logged out successfully",
	})
}
