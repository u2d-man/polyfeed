package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/mmcdole/gofeed"
	"golang.org/x/text/unicode/norm"

	"github.com/u2d-man/polyfeed/internal/fetcher"
)

type Article struct {
	Title    string `json:"title"`
	Link     string `json:"link"`
	Analyzed string `json:"analyzed"`
}

func main() {
	rssFetcher := fetcher.StaticRSSFetcher{URLs: []string{}}
	url, err := rssFetcher.GetRssURLs()
	articles, err := fetchRSS(url)
	if err != nil {
		log.Fatal("Error fetching RSS:", err)
	}

	if len(articles) > 0 {
		err = saveArticles(articles, "fetch_rss.json")
		if err != nil {
			log.Fatal("Error saving articles:", err)
		}
		fmt.Println("Articles saved to fetch_rss.json")
	} else {
		fmt.Println("No English articles found.")
	}
}

func fetchRSS(url []string) ([]Article, error) {
	var articles []Article

	for _, url := range url {
		fp := gofeed.NewParser()
		feed, err := fp.ParseURL(url)
		if err != nil {
			return nil, err
		}
		for _, item := range feed.Items {
			text := norm.NFC.String(item.Description)
			analyzed, err := requestOpenAI(text)
			if err != nil {
				return nil, err
			}
			articles = append(articles, Article{
				Title:    item.Title,
				Link:     item.Link,
				Analyzed: analyzed,
			})
		}
	}

	return articles, nil
}

func saveArticles(articles []Article, filename string) error {
	data, err := json.MarshalIndent(articles, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func getOpenAIAPIKey() (string, error) {
	key := os.Getenv("OPENAI_API_KEY")

	if key == "" {
		return key, fmt.Errorf("OPENAI_API_KEY environment variable not set")
	}

	return key, nil
}

func requestOpenAI(text string) (string, error) {
	if text == "" {
		fmt.Println("No text to open")
		return "", nil
	}

	prompt := fmt.Sprintf(`以下の文章の内容を分析し、要点を簡潔にまとめてください。
- 主要なトピックは何か？
- 記事の目的や意図は何か？
- 読者にとって重要なポイントは何か？

%s`, text)

	reqBody, _ := json.Marshal(map[string]interface{}{
		"model": "gpt-4o-mini",
		"messages": []map[string]string{
			{"role": "user", "content": fmt.Sprint(prompt)},
		},
	})

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	apiKey, err := getOpenAIAPIKey()
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)

	body, _ := io.ReadAll(resp.Body)
	var res map[string]interface{}
	json.Unmarshal(body, &res)

	if choices, ok := res["choices"].([]interface{}); ok && len(choices) > 0 {
		if choice, ok := choices[0].(map[string]interface{}); ok {
			if msg, ok := choice["message"].(map[string]interface{}); ok {
				if content, ok := msg["content"].(string); ok {
					return content, nil
				}
			}
		}
	}
	return "", fmt.Errorf("failed to get translation")
}
