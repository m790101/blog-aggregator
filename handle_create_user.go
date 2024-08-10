package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/m790101/blog-aggregator/internal/database"
)

func (cfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {

	type Params struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	req := Params{}

	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	trimmedName := strings.TrimSpace(req.Name)
	if trimmedName == "" {
		respondWithError(w, http.StatusBadRequest, "name cannot be empty")
		return
	}

	id := uuid.New().String()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	createReq := database.CreateUserParams{
		ID:        id,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      trimmedName,
	}
	user, err := cfg.DB.CreateUser(r.Context(), createReq)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}
