package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/m790101/blog-aggregator/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	dbURL := os.Getenv("DBURL")
	mux := http.NewServeMux()

	db, err := runDB(dbURL)

	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	dbQueries := database.New(db)

	cfg := &apiConfig{
		DB: dbQueries,
	}
	go startScrapping(dbQueries, 5, time.Minute)

	mux.HandleFunc("POST /api/v1/users", cfg.handleCreateUser)
	mux.HandleFunc("GET /api/v1/users", cfg.middlewareAuth(cfg.handleGetUser))
	mux.HandleFunc("POST /api/v1/feeds", cfg.middlewareAuth(cfg.handleCreateFeeds))
	mux.HandleFunc("GET /api/v1/feeds", cfg.handleGetFeeds)

	mux.HandleFunc("POST /api/v1/feed_follows", cfg.middlewareAuth(cfg.handleAddFollowFeed))
	mux.HandleFunc("DELETE /api/v1/feed_follows/{feedFollowID}", cfg.middlewareAuth(cfg.handleRemoveFollowFeeds))
	mux.HandleFunc("GET /api/v1/feed_follows", cfg.middlewareAuth(cfg.handleGetFollowFeeds))
	mux.HandleFunc("GET /api/v1/posts", cfg.middlewareAuth(cfg.getPostsByUser))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Server is running on port %s", port)
	log.Fatal(srv.ListenAndServe())

}

func runDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	return db, nil
}
