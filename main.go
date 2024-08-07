package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthz", handleHealthz)
	mux.HandleFunc("GET /v1/err", handleErr)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Server is running on port %s", port)
	log.Fatal(srv.ListenAndServe())

}

func handleHealthz(w http.ResponseWriter, r *http.Request) {
	type responseHealthz struct {
		Status string `json:"status"`
	}

	respondWithJSON(w, http.StatusOK, responseHealthz{Status: "ok"})
}

func handleErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "error message")
}
