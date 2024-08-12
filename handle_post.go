package main

import (
	"net/http"
	"strconv"

	"github.com/m790101/blog-aggregator/internal/database"
)

func (cfg *apiConfig) getPostsByUser(w http.ResponseWriter, r *http.Request, user database.User) {

	limit := r.URL.Query().Get("limit")
	if len(limit) < 1 {
		limit = "10"
	}
	limitNum, err := strconv.Atoi(limit)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "limit number is error")
	}

	data, err := cfg.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  int32(limitNum),
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	respondWithJSON(w, http.StatusOK, data)
}
