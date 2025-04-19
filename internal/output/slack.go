package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/u2d-man/polyfeed/internal/core"
)

// SlackOutput implements the OutputService interface for Slack.
// It sends summarized articles to a Slack channel via webhook.
type SlackOutput struct {
	// WebhookURL is the Slack Incoming Webhook URL
	WebhookURL string

	// Client is the HTTP client used for making the webhook request
	Client *http.Client
}

// Send outputs the provided articles to a Slack channel via webhook.
// Each article is formatted with its title, link, publication date, and summary.
//
// Parameters:
//   - articles: The list of articles to send to Slack
//
// Returns:
//   - nil if the articles were successfully sent
//   - An error if there was a problem sending the message
//
// If the articles list is empty, the method returns nil without sending a message.
func (s SlackOutput) Send(articles []core.Article) error {
	if len(articles) == 0 {
		return nil
	}

	var message string
	for i, a := range articles {
		entry := fmt.Sprintf(
			"%d. %s\n%s\n投稿日: %s\n\n%s\n\n------------------------------\n\n",
			i+1,
			a.Title,
			a.Link,
			a.Published,
			a.Analyzed,
		)
		message += entry
	}

	payload := map[string]string{"text": message}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := s.Client.Post(s.WebhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read Slack API response: %w", err)
		}

		return fmt.Errorf("slack API returned status %d: %s", resp.StatusCode, respBody)
	}

	return nil
}
