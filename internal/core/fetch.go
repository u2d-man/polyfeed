package core

import (
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
	"golang.org/x/text/unicode/norm"
)

// FeedParser defines the interface for parsing RSS feeds.
// It abstracts the details of feed parsing implementation, allowing for
// different parser implementations or mocking for tests.
type FeedParser interface {
	// ParseURL fetches and parses an RSS feed from the provided URL
	// Returns the parsed feed or an error if fetching or parsing fails
	ParseURL(feedURL string) (*gofeed.Feed, error)
}

// FetchArticles retrieves and processes articles from the provided RSS feed URLs.
// It performs the following operations for each URL:
// 1. Fetches and parses the RSS feed
// 2. Filters articles based on publication date
// 3. Summarizes article content using OpenAI
// 4. Validates required article fields
//
// Parameters:
//   - parser: The RSS feed parser implementation
//   - urls: A list of RSS feed URLs to fetch articles from
//
// Returns:
//   - A slice of processed Article structures
//   - An error if feed fetching, parsing, or processing fails
//
// Articles published before the cutoff time (current time minus TimeWindow hours)
// are filtered out. Articles without titles, links, or descriptions are also skipped.
func FetchArticles(parser FeedParser, urls []string) ([]Article, error) {
	var allArticles []Article

	now := time.Now()
	cutoffTime := now.Add(-time.Duration(TimeWindow) * time.Hour)
	cutoffTimeStr := cutoffTime.Format(TimeLayout)

	fmt.Printf("fetching articles published after: %s\n", cutoffTimeStr)

	for _, url := range urls {
		fmt.Printf("Start fetch. %s \n", url)

		feed, err := parser.ParseURL(url)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", url, err)
		}

		for _, item := range feed.Items {
			if item.Published == "" {
				fmt.Printf("pulished is blank. %s\n", item.Title)

				continue
			}

			published, err := ParseAndFormatTime(item.Published)
			if err != nil {
				fmt.Printf("warning: Failed to parse publication time for %s: %v\n", item.Title, err)

				continue
			}

			pubTime, err := time.Parse(TimeLayout, published)
			if err != nil {
				fmt.Printf("warning: Failed to parse formatted time %s: %v\n", published, err)

				continue
			}

			if pubTime.Before(cutoffTime) {
				fmt.Printf("skipping older article from %s: %s\n", published, item.Title)

				continue
			}

			if item.Description == "" {
				fmt.Printf("skipping empty description %s\n", item.Title)

				continue
			}

			analyzed, err := SummarizeContent(norm.NFC.String(item.Description))
			if err != nil {
				return nil, fmt.Errorf("failed to summarize '%s': %w", item.Title, err)
			}

			if item.Title == "" || item.Link == "" {
				fmt.Printf("warning: Article missing required fields: title=%q. link=%q\n", item.Title, item.Link)

				continue
			}

			allArticles = append(allArticles, Article{
				Title:     item.Title,
				Link:      item.Link,
				Analyzed:  analyzed,
				Published: published,
			})
		}
	}
	return allArticles, nil
}
