package tests

import (
	"testing"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/u2d-man/polyfeed/internal/core"
)

type mockParser struct{}

func (m *mockParser) ParseURL(url string) (*gofeed.Feed, error) {
	return &gofeed.Feed{
		Items: []*gofeed.Item{
			{
				Title:       "Mock Title",
				Link:        "https://example.com",
				Description: "This is test article.",
				Published:   time.Now().Format(time.RFC1123),
			},
		},
	}, nil
}

// Mock SummarizeContent to avoid real API call
func init() {
	core.SummarizeContent = func(text string) (string, error) {
		return "要約されたテキスト", nil
	}
}

func TestFetchArticles_WithMockParser(t *testing.T) {
	parser := &mockParser{}
	urls := []string{"https://example.com/rss"}

	articles, err := core.FetchArticles(parser, urls)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(articles) != 1 {
		t.Fatalf("expected 1 article, got %d", len(articles))
	}

	article := articles[0]
	if article.Title != "Mock Title" {
		t.Errorf("unexpected title: %s", article.Title)
	}
	if article.Analyzed != "要約されたテキスト" {
		t.Errorf("unexpected analyzed result: %s", article.Analyzed)
	}
}
