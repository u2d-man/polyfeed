package main

import (
	"fmt"
	"log"

	"github.com/u2d-man/polyfeed/internal/core"
	"github.com/u2d-man/polyfeed/internal/fetcher"
)

func main() {
	rssFetcher := fetcher.StaticRSSFetcher{URLs: []string{}}

	urls, err := rssFetcher.GetRssURLs()
	if err != nil {
		log.Fatalf("Failed to get RSS URLs: %v", err)
	}

	articles, err := core.FetchArticles(urls)
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

	fmt.Printf("Articles saved to %s\n", core.OutputFile)
}
