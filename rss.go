package main

import (
	"encoding/xml"
)

type RssFeed struct {
	Channel struct {
		Text          string    `xml:",chardata"`
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
	Text        string   `xml:",chardata"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Guid        string   `xml:"guid"`
	Description string   `xml:"description"`
}

// func urlToFeed(url string) (RssFeed, error) {

// }
