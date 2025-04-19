package core

import "time"

const (
	OutputFile     = "fetch_rss.json"
	OpenAIURL      = "https://api.openai.com/v1/chat/completions"
	Model          = "gpt-4o-mini"
	TimeLayout     = "2006-01-02 15:04:05"
	InputFormat    = "Mon, 02 Jan 2006 15:04:05 MST"
	InputFormatISO = time.RFC3339
	TimeWindow     = 24 // 当日の記事を取得するために24時間に変更
	EnvAPIKey      = "OPENAI_API_KEY"
)
