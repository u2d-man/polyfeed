package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/u2d-man/polyfeed/internal/core"
)

type SlackOutput struct {
	WebhookURL string
	Client     *http.Client
}

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

	_, err = s.Client.Post(s.WebhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	return err
}
