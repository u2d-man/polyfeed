package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/u2d-man/polyfeed/internal/core"
	"github.com/u2d-man/polyfeed/internal/output"
)

func TestSlackSend_Test_Success(t *testing.T) {
	var receivedBody []byte

	// mock server for testing
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("failed to read request body: %v", err)
		}
		receivedBody = body
		w.WriteHeader(http.StatusOK)
	}))
	defer mockServer.Close()

	slack := output.SlackOutput{
		WebhookURL: mockServer.URL,
		Client:     mockServer.Client(),
	}

	articles := []core.Article{
		{
			Title:     "This is article title.",
			Link:      "https://example.com/article",
			Analyzed:  "This is article analyzed.",
			Published: "2025-03-23 12:34:56",
		},
	}

	err := slack.Send(articles)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}

	// Was a correctly formatted message sent?
	var payload map[string]string
	err = json.Unmarshal(receivedBody, &payload)
	if err != nil {
		t.Fatalf("failed to unmarshal payload: %v", err)
	}

	text, ok := payload["text"]
	if !ok {
		t.Errorf("payload missing 'text' field")
	}

	if !strings.Contains(text, "This is article title.") || !strings.Contains(text, "https://example.com/article") {
		t.Errorf("unexpected payload text: %s", text)
	}
}

func TestSlackOutput_Send_EmptyArticles(t *testing.T) {
	slack := output.SlackOutput{
		WebhookURL: "https://example.com/should/not/be/called",
		Client:     http.DefaultClient,
	}

	// If an empty array is passed, expect to send nothing and return nil.
	err := slack.Send([]core.Article{})
	if err != nil {
		t.Errorf("expected no error when sending empty articles, got: %v", err)
	}
}
