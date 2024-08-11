package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/m790101/blog-aggregator/internal/database"
)

func (cfg *apiConfig) handleAddFollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type paramAddFollow struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	req := paramAddFollow{}
	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	id := uuid.New()

	res := database.AddFeedFollowParams{
		ID:        id,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    req.FeedId,
	}

	feed, err := cfg.DB.AddFeedFollow(r.Context(), res)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	respondWithJSON(w, http.StatusCreated, feed)
}

func (cfg *apiConfig) handleRemoveFollowFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	path := r.PathValue("feedFollowID")
	feedId := uuid.MustParse(path)

	cfg.DB.RemoveFeedFollow(r.Context(), feedId)

	respondWithJSON(w, http.StatusOK, "")
}

func (cfg *apiConfig) handleGetFollowFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := cfg.DB.GetFeedFollow(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, feeds)
}
