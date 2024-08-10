package main

import (
	"net/http"

	"github.com/m790101/blog-aggregator/internal/auth"
)

func (cfg *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetAuthToken(r.Header)

	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := cfg.DB.GetUser(r.Context(), token)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}
