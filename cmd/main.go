package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mmcdole/gofeed"
	"github.com/u2d-man/polyfeed/internal/core"
	"github.com/u2d-man/polyfeed/internal/fetcher"
	"github.com/u2d-man/polyfeed/internal/output"
)

func main() {
	if (len(os.Args)) < 2 {
		fmt.Println("Usage: polyfeed <rss_file.txt>")
		os.Exit(1)
	}
	rssFilePath := os.Args[1]

	rssFetcher := fetcher.FileRSSFetcher{Path: rssFilePath}
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

	if err := core.SaveToFile(articles, core.OutputFile); err != nil {
		log.Fatalf("Failed to save articles: %v", err)
	}

	s := output.SlackOutput{WebhookURL: os.Getenv("WEBHOOK_URL"), Client: http.DefaultClient}
	if err := s.Send(articles); err != nil {
		log.Fatalf("Failed output to slack: %v", err)
	}
}
