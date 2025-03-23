package core

import (
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
	"golang.org/x/text/unicode/norm"
)

func FetchArticles(urls []string) ([]Article, error) {
	var allArticles []Article
	parser := gofeed.NewParser()
	cutoff := time.Now().Add(-TimeWindow * time.Hour).Format(TimeLayout)

	for _, url := range urls {
		feed, err := parser.ParseURL(url)
		if err != nil {
			return nil, fmt.Errorf("failed to parse feed from %s: %w", url, err)
		}

		for _, item := range feed.Items {
			published, err := ParseAndFormatTime(item.Published)
			if err != nil || published <= cutoff {
				continue
			}

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
