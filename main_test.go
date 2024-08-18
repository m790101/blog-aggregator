package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestRun(t *testing.T) {
	godotenv.Load()
	dbURL := os.Getenv("DBURL")

	_, err := runDB(dbURL)
	if err != nil {
		t.Error("fail test")
	}
}

func TestGetUsers(t *testing.T) {

	mux := getRoutes()

	ts := httptest.NewServer(mux)

	res, err := ts.Client().Get(ts.URL + "/api/v1/users")
	if err != nil {
		t.Error(err)
	} else {
		if res.StatusCode != http.StatusUnauthorized {
			log.Printf("status code is %d", res.StatusCode)
			t.Error("wrong status")
		}
	}

	defer ts.Close()

}

func TestGetFeeds(t *testing.T) {
	mux := getRoutes()

	ts := httptest.NewServer(mux)

	res, err := ts.Client().Get(ts.URL + "/api/v1/feeds")
	if err != nil {
		t.Error(err)
	} else {
		if res.StatusCode != http.StatusOK {
			log.Printf("status code is %d", res.StatusCode)
			t.Error("wrong status")
		}
	}

	defer ts.Close()

}
