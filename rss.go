package main

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"time"
)

type RssFeed struct {
	Channel struct {
		Text          string    `xml:"text"`
		Title         string    `xml:"title"`
		Link          string    `xml:"link"`
		Description   string    `xml:"description"`
		Generator     string    `xml:"generator"`
		Language      string    `xml:"language"`
		LastBuildDate string    `xml:"lastBuildDate"`
		Item          []XmlItem `xml:"item"`
	} `xml:"channel"`
}

type XmlItem struct {
	XMLName     xml.Name `xml:"item"`
	Text        string   `xml:"text"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Guid        string   `xml:"guid"`
	Description string   `xml:"description"`
}

func urlToFeed(url string) (RssFeed, error) {

	Client := http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("error creating request: %v", err)
		return RssFeed{}, err
	}
	res, err := Client.Do(req)
	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error reading response body: %v", err)
		return RssFeed{}, err
	}

	feed := RssFeed{}

	err = xml.Unmarshal(data, &feed)
	if err != nil {
		return RssFeed{}, err
	}

	return feed, nil

}
