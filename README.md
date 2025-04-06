# PolyFeed

> ⚠️ This project is under active development. Features and usage may change frequently.

**PolyFeed** is a CLI tool written in Go that fetches RSS articles, summarizes them using the OpenAI API, and outputs the results to local files or Slack.

## 🚀 Features

- Fetch articles from any RSS feed
- Summarize content using OpenAI's GPT API
- Output results to:
  - Local JSON or text files
  - Slack channels via Incoming Webhooks
- Supports static input files for development/testing
- Includes unit tests for core functionality

## 🛠 Usage

```bash
go run ./cmd/main.go
```

> ⚠️ This project is under active development. Configuration via flags or environment variables will be added in a future release.

## 📁 Project Structure

```go
cmd/ - Application entry point
internal/core/ - Core logic: data models, fetchers, OpenAI integration
internal/fetcher/ - RSS and static file fetchers
internal/output/ - Output handlers (e.g. file writer, Slack)
tests/ - Unit tests
```

## 📦 Requirements

Go 1.23+

OpenAI API key

(Optional) Slack Incoming Webhook URL

## 🔧 Configuration

Environment variable support is planned in a future version. The following variables are expected:

| Variable            | Description                           |
| ------------------- | ------------------------------------- |
| `OPENAI_API_KEY`    | Your OpenAI API key                   |
| `SLACK_WEBHOOK_URL` | Slack Incoming Webhook URL (optional) |
| `RSS_FEED_URLS`     | Comma-separated list of RSS feed URLs |
| `OUTPUT_DIR`        | Output directory path (optional)      |

## 📄 Example Output

Summarized Article (JSON):

```json
{
  "title": "New AI Tool Released",
  "link": "https://example.com/article",
  "summary": "OpenAI has released a new tool for summarizing text using GPT-4..."
}
```

Slack Message (Markdown):

```arduino
📰 _New AI Tool Released_
<https://example.com/article>
> OpenAI has released a new tool for summarizing text using GPT-4...
```

## 🧪 Running Tests

```
go test ./...
```

## 📄 License

MIT License. See LICENSE for details.
