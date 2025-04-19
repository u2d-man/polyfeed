package core

import (
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
	"golang.org/x/text/unicode/norm"
)

type FeedParser interface {
	ParseURL(feedURL string) (*gofeed.Feed, error)
}

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

			analyzed, err := SummarizeContent(norm.NFC.String(item.Description))
			if err != nil {
				return nil, err
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
