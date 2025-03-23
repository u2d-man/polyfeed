package tests

import (
	"testing"

	"github.com/u2d-man/polyfeed/internal/fetcher"
)

func TestStaticFetcher(t *testing.T) {
	// Is the interface implemented? use type RSSFetcher not StaticRSSFetcher.
	var f fetcher.RSSFetcher = fetcher.StaticRSSFetcher{
		URLs: []string{"https://example.com/rss/sample.xml"},
	}

	urls, err := f.GetRssURLs()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(urls) != 1 || urls[0] != "https://example.com/rss/sample.xml" {
		t.Errorf("Expected [https://example.com/rss/sample.xml], got %v", urls)
	}
}
