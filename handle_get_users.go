package main

import (
	"net/http"

	"github.com/m790101/blog-aggregator/internal/database"
)

func (cfg *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, user)
}
