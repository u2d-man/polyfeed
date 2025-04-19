package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Defined in an anonymous function to be replaced at test time.
var SummarizeContent = func(text string) (string, error) {
	if text == "" {
		fmt.Println("No text to summarize")
		return "", nil
	}

	prompt := fmt.Sprintf(`以下の文章の内容を分析し、要点を簡潔にまとめてください。
- 主要なトピック
- 記事の目的や意図
- 読者にとって重要なポイント
- ChatGPT がこの記事を読んだ感想・注意点

以下注意点に注意してください。
- markdown 記法は使用し
ないでください。
%s`, text)

	client := &http.Client{Timeout: 30 * time.Second}

	payload := map[string]interface{}{
		"model": Model,
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	apiKey, err := GetAPIKey()
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", OpenAIURL, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)

		return "", fmt.Errorf("openAI API returned non-OK status: %d, body: %s", resp.StatusCode, respBody)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read OpenAI response: %w", err)
	}
	var res struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.Unmarshal(respBody, &res); err != nil {
		return "", fmt.Errorf("failed to unmarshal OpenAI response: %w", err)
	}

	if len(res.Choices) == 0 {
		return "", fmt.Errorf("no choices returned from OpenAI")
	}

	return res.Choices[0].Message.Content, nil
}
