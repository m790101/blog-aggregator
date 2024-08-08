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

	type CreateReq struct {
		ID        int32  `json:"id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		Name      string `json:"name"`
	}

	id := uuid.New().String()

	createReq := database.CreateUserParams{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      trimmedName,
	}

	user, err := cfg.DB.CreateUser(r.Context(), createReq)

	respondWithJSON(w, http.StatusCreated, user)
}
