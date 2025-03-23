package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SummarizeContent(text string) (string, error) {
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
