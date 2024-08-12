package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/m790101/blog-aggregator/internal/database"
)

func startScrapping(db *database.Queries, concurrency int, timeInterval time.Duration) {
	log.Printf("Start scrapping with concurrency %d and time interval %v", concurrency, timeInterval)
	ticker := time.NewTicker(timeInterval)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedToFetch(context.Background(), int32(concurrency))

		if err != nil {
			log.Printf("Error getting feeds to fetch: %v", err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)

		}
		wg.Wait()
	}

}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedFetch(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error marking feed as fetched: %v", err)
	}
	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("Error fetching feed: %v", err)
		return
	}
	for _, item := range rssFeed.Channel.Item {
		log.Printf("feed title %s, ", item.Title)
	}

	log.Printf("Feed fetched: %s, numbers %d found", feed.Url, len(rssFeed.Channel.Item))
}
