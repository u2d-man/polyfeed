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
	cutoff := time.Now().Add(-TimeWindow * time.Hour).Format(TimeLayout)
	fmt.Println(cutoff)

	for _, url := range urls {
		fmt.Printf("Start fetch. %s \n", url)

		feed, err := parser.ParseURL(url)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", url, err)
		}

		for _, item := range feed.Items {
			published, err := ParseAndFormatTime(item.Published)
			if err != nil || published <= cutoff {
				continue
			}
			fmt.Println(published)

			analyzed, err := SummarizeContent(norm.NFC.String(item.Description))
			if err != nil {
				return nil, err
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
