package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mmcdole/gofeed"
	"github.com/u2d-man/polyfeed/internal/core"
	"github.com/u2d-man/polyfeed/internal/fetcher"
	"github.com/u2d-man/polyfeed/internal/output"
)

func main() {
	rssFetcher := fetcher.StaticRSSFetcher{URLs: []string{""}}

	urls, err := rssFetcher.GetRssURLs()
	if err != nil {
		log.Fatalf("Failed to get RSS URLs: %v", err)
	}

	parser := gofeed.NewParser()
	articles, err := core.FetchArticles(parser, urls)
	if err != nil {
		log.Fatalf("Failed to fetch articles: %v", err)
	}

	if len(articles) == 0 {
		fmt.Println("No articles found.")
		return
	}

	s := output.SlackOutput{WebhookURL: os.Getenv("WEBHOOK_URL")}
	if err := s.Send(articles); err != nil {
		log.Fatalf("Failed output to slack: %v", err)
	}

	if err := core.SaveToFile(articles, core.OutputFile); err != nil {
		log.Fatalf("Failed to save articles: %v", err)
	}

	fmt.Printf("Articles saved to %s\n", core.OutputFile)
}
