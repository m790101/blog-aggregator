package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func getRoutes() *http.ServeMux {
	godotenv.Load()
	dbURL := os.Getenv("DBURL")
	mux := http.NewServeMux()

	_, err := runDB(dbURL)

	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	cfg := &apiConfig{}

	mux.HandleFunc("POST /api/v1/users", cfg.handleCreateUser)
	mux.HandleFunc("GET /api/v1/users", cfg.middlewareAuth(cfg.handleGetUser))
	mux.HandleFunc("POST /api/v1/feeds", cfg.middlewareAuth(cfg.handleCreateFeeds))
	mux.HandleFunc("GET /api/v1/feeds", cfg.handleGetFeeds)

	mux.HandleFunc("POST /api/v1/feed_follows", cfg.middlewareAuth(cfg.handleAddFollowFeed))
	mux.HandleFunc("DELETE /api/v1/feed_follows/{feedFollowID}", cfg.middlewareAuth(cfg.handleRemoveFollowFeeds))
	mux.HandleFunc("GET /api/v1/feed_follows", cfg.middlewareAuth(cfg.handleGetFollowFeeds))
	mux.HandleFunc("GET /api/v1/posts", cfg.middlewareAuth(cfg.getPostsByUser))

	return mux
}
