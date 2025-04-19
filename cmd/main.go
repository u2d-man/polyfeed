// Package main provides the entry point for the PolyFeed application.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/u2d-man/polyfeed/internal/core"
	"github.com/u2d-man/polyfeed/internal/fetcher"
	"github.com/u2d-man/polyfeed/internal/output"
)

// PolyFeed is a CLI tool that fetches RSS articles, summarizes them using OpenAI,
// and outputs the results to local files or Slack.
//
// Usage:
//
//	polyfeed <rss_file.txt>
//
// Environment Variables:
//
//	OPENAI_API_KEY - Required: Your OpenAI API key
//	WEBHOOK_URL - Optional: Slack webhook URL for output
//
// The rss_file.txt should contain one RSS feed URL per line.
// Articles published within the last 24 hours will be fetched,
// summarized, and saved to fetch_rss.json in the current directory.
// If WEBHOOK_URL is set, the results will also be sent to Slack.
func main() {
	if (len(os.Args)) < 2 {
		fmt.Println("usage: polyfeed <rss_file.txt>")
		os.Exit(1)
	}
	rssFilePath := os.Args[1]

	if _, err := os.Stat(rssFilePath); os.IsNotExist(err) {
		log.Fatalf("file does not exists: %s", rssFilePath)
	}

	rssFetcher := fetcher.FileRSSFetcher{Path: rssFilePath}
	urls, err := rssFetcher.GetRssURLs()
	if err != nil {
		log.Fatalf("failed to get RSS URLs: %v", err)
	}

	parser := gofeed.NewParser()
	articles, err := core.FetchArticles(parser, urls)
	if err != nil {
		log.Fatalf("failed to fetch articles: %v", err)
	}

	if len(articles) == 0 {
		fmt.Println("no articles found.")
		return
	}

	if err := core.SaveToFile(articles, core.OutputFile); err != nil {
		log.Fatalf("failed to save articles: %v", err)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	webhookURL := os.Getenv("WEBHOOK_URL")
	if webhookURL != "" {
		s := output.SlackOutput{WebhookURL: webhookURL, Client: client}
		if err := s.Send(articles); err != nil {
			log.Fatalf("failed output to slack: %v", err)
		}
	}
}
