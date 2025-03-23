package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/u2d-man/polyfeed/internal/fetcher"
	"golang.org/x/text/unicode/norm"
)

const (
	outputFile  = "fetch_rss.json"
	openAIURL   = "https://api.openai.com/v1/chat/completions"
	model       = "gpt-4o-mini"
	timeLayout  = "2006-01-02 15:04:05"
	inputFormat = time.RFC1123
	timeWindow  = 24 * time.Hour
	envAPIKey   = "OPENAI_API_KEY"
)

type Article struct {
	Title     string `json:"title"`
	Link      string `json:"link"`
	Analyzed  string `json:"analyzed"`
	Published string `json:"published"`
}

func main() {
	rssFetcher := fetcher.StaticRSSFetcher{URLs: []string{"https://zenn.dev/feed"}}

	urls, err := rssFetcher.GetRssURLs()
	if err != nil {
		log.Fatalf("Failed to get RSS URLs: %v", err)
	}

	articles, err := fetchArticles(urls)
	if err != nil {
		log.Fatalf("Failed to fetch articles: %v", err)
	}

	if len(articles) == 0 {
		fmt.Println("No articles found.")
		return
	}

	if err := saveToFile(articles, outputFile); err != nil {
		log.Fatalf("Failed to save articles: %v", err)
	}

	fmt.Printf("Articles saved to %s\n", outputFile)
}

func fetchArticles(urls []string) ([]Article, error) {
	var allArticles []Article
	parser := gofeed.NewParser()
	cutoff := time.Now().Add(-timeWindow).Format(timeLayout)

	for _, url := range urls {
		feed, err := parser.ParseURL(url)
		if err != nil {
			return nil, fmt.Errorf("failed to parse feed from %s: %w", url, err)
		}

		for _, item := range feed.Items {
			published, err := parseAndFormatTime(item.Published)
			if err != nil || published <= cutoff {
				continue
			}

			// request OpenAI API
			analyzed, err := summarizeContent(norm.NFC.String(item.Description))
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

func parseAndFormatTime(raw string) (string, error) {
	t, err := time.Parse(inputFormat, raw)
	if err != nil {
		return "", err
	}
	return t.Format(timeLayout), nil
}

func saveToFile(data any, filename string) error {
	content, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, content, 0644)
}

func getAPIKey() (string, error) {
	key := os.Getenv(envAPIKey)
	if key == "" {
		return "", fmt.Errorf("%s environment variable not set", envAPIKey)
	}
	return key, nil
}

func summarizeContent(text string) (string, error) {
	if text == "" {
		fmt.Println("No text to summarize")
		return "", nil
	}

	prompt := fmt.Sprintf(`以下の文章の内容を分析し、要点を簡潔にまとめてください。
- 主要なトピックは何か？
- 記事の目的や意図は何か？
- 読者にとって重要なポイントは何か？

%s`, text)

	payload := map[string]interface{}{
		"model": model,
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	apiKey, err := getAPIKey()
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", openAIURL, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)

	respBody, _ := io.ReadAll(resp.Body)
	var res struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.Unmarshal(respBody, &res); err != nil {
		return "", err
	}

	if len(res.Choices) == 0 {
		return "", fmt.Errorf("no choices returned from OpenAI")
	}

	return res.Choices[0].Message.Content, nil
}
