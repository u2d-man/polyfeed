# PolyFeed

[![Go Test](https://github.com/u2d-man/polyfeed/actions/workflows/go-test.yml/badge.svg)](https://github.com/u2d-man/polyfeed/actions/workflows/go-test.yml)

PolyFeed is a CLI tool that fetches RSS articles, summarizes them using OpenAI, and outputs the results to local files or Slack.

## Features

- Fetches articles from multiple RSS feeds
- Filters articles published within the last 24 hours
- Summarizes article content using OpenAI GPT models
- Outputs results to JSON file and optionally to Slack

## Usage

### Local Execution

```bash
# Set your OpenAI API key as an environment variable
export OPENAI_API_KEY=your_openai_api_key

# Optional: Set Slack webhook URL for Slack output
export WEBHOOK_URL=your_slack_webhook_url

# Run the program with the path to your RSS feed list file
polyfeed rss_feeds.txt
```

### Automated Execution with GitHub Actions

PolyFeed can be automatically executed using GitHub Actions on a schedule. To set this up:

1. Fork this repository
2. Set up the following GitHub repository secrets:
   - `OPENAI_API_KEY`: Your OpenAI API key
   - `SLACK_WEBHOOK_URL`: (Optional) Your Slack webhook URL
   - `RSS_FEEDS`: A list of RSS feed URLs, one per line

The workflow will run daily at 9:00 UTC (18:00 JST) and process the RSS feeds. Results are stored as GitHub artifacts and can be sent to Slack.

## RSS Feed List Format

Create a text file with one RSS feed URL per line:

```
https://example.com/feed.xml
https://another-site.com/rss
# Lines starting with # are treated as comments
```

## Installation

```bash
go install github.com/u2d-man/polyfeed@latest
```

## Output

Processed articles are saved to `fetch_rss.json` in the current directory. If a Slack webhook URL is provided, the results will also be posted to the associated Slack channel.

## Configuration

- OpenAI model: `gpt-4o-mini` (default)
- Time window: 24 hours (articles published within this window are processed)

## Requirements

- Go 1.23 or higher
- An OpenAI API key
- (Optional) A Slack webhook URL