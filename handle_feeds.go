package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/m790101/blog-aggregator/internal/database"
)

func (cfg *apiConfig) handleCreateFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	type paramFeed struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	req := paramFeed{}
	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	id := uuid.New()

	res := database.CreateFeedParams{
		ID:        id,
		Name:      req.Name,
		Url:       req.Url,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), res)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	idFollow := uuid.New()

	cfg.DB.AddFeedFollow(r.Context(), database.AddFeedFollowParams{
		ID:        idFollow,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	respondWithJSON(w, http.StatusCreated, feed)
}

func (cfg *apiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetFeed(r.Context())
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, feeds)
}
