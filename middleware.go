package main

import (
	"net/http"

	"github.com/m790101/blog-aggregator/internal/auth"
	"github.com/m790101/blog-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
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
		handler(w, r, user)
	}
}
