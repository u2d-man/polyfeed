package core

import "time"

const (
	// OutputFile is the default filename for storing fetched articles
	OutputFile = "fetch_rss.json"

	// OpenAIURL is the endpoint for the OpenAI Chat Completions API
	OpenAIURL = "https://api.openai.com/v1/chat/completions"

	// Model specifies which OpenAI model to use for article summarization
	Model = "gpt-4o-mini"

	// TimeLayout defines the standard time format for the application (Go time format)
	TimeLayout = "2006-01-02 15:04:05"

	// InputFormat is the expected time format for RFC1123 dates from RSS feeds
	InputFormat = "Mon, 02 Jan 2006 15:04:05 MST"

	// InputFormatISO is the fallback time format (RFC3339) for RSS feeds
	InputFormatISO = time.RFC3339

	// TimeWindow defines how many hours back to consider articles as "recent"
	// Articles published before this window will be skipped
	TimeWindow = 24

	// EnvAPIKey is the name of the environment variable containing the OpenAI API key
	EnvAPIKey = "OPENAI_API_KEY"
)
