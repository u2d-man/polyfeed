# PolyFeed

> ⚠️ This project is under active development. Features and usage may change frequently.

**PolyFeed** is a CLI tool written in Go that fetches RSS articles, summarizes them using the OpenAI API, and outputs the results to local files or Slack.

## 🚀 Features

- Fetch articles from RSS feeds using URLs listed in a text file
- Summarize content using OpenAI's GPT-4o-mini API
- Output results to:
  - Local JSON files
  - Slack channels via Incoming Webhooks
- Supports static input files for development/testing
- Includes unit tests for core functionality

## 🛠 Usage

```bash
go run ./cmd/main.go <rss_file.txt>
```

Where `<rss_file.txt>` is a text file containing RSS feed URLs, one per line.

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

Go 1.23.5

OpenAI API key

(Optional) Slack Incoming Webhook URL

## 🔧 Configuration

The following environment variables are currently supported:

| Variable          | Description                           |
| ----------------- | ------------------------------------- |
| `OPENAI_API_KEY`  | Your OpenAI API key (required)        |
| `WEBHOOK_URL`     | Slack Incoming Webhook URL (optional) |

Future versions plan to add support for:
| Variable       | Description                             |
| -------------- | --------------------------------------- |
| `RSS_FEED_URLS`| Comma-separated list of RSS feed URLs   |
| `OUTPUT_DIR`   | Custom output directory path            |

## 📄 Example Output

Summarized Article (JSON):

```json
{
  "title": "New AI Tool Released",
  "link": "https://example.com/article",
  "summary": "OpenAI has released a new tool for summarizing text using GPT-4..."
}
```

Slack Message (Plain Text):

```
1. New AI Tool Released
https://example.com/article
投稿日: 2023-04-19 15:04:05

OpenAI has released a new tool for summarizing text using GPT-4...

------------------------------
```

## 🧪 Running Tests

```
go test ./...
```

## 📄 License

MIT License. See LICENSE for details.
