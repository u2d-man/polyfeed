// Package examples provides usage examples for the PolyFeed application.
package examples

// This file provides example usage of PolyFeed's key components.
// These examples are meant for documentation and are not executed as part of the application.

// ExampleBasicUsage demonstrates the basic usage flow of PolyFeed.
func ExampleBasicUsage() {
	// Example: Basic usage of PolyFeed
	//
	// Create a text file with RSS URLs:
	//   $ echo "https://example.com/rss.xml" > feeds.txt
	//
	// Set the OpenAI API key:
	//   $ export OPENAI_API_KEY="your-api-key"
	//
	// Run PolyFeed:
	//   $ polyfeed feeds.txt
	//
	// The summarized articles will be saved to fetch_rss.json in the current directory.
}

// ExampleSlackIntegration demonstrates how to configure PolyFeed to send results to Slack.
func ExampleSlackIntegration() {
	// Example: Sending results to Slack
	//
	// First, set up a Slack Incoming Webhook URL:
	// 1. Go to your Slack workspace
	// 2. Add the "Incoming WebHooks" app
	// 3. Create a new configuration and copy the webhook URL
	//
	// Then set the environment variables:
	//   $ export OPENAI_API_KEY="your-api-key"
	//   $ export WEBHOOK_URL="your-slack-webhook-url"
	//
	// Run PolyFeed:
	//   $ polyfeed feeds.txt
	//
	// The results will be both saved to file and sent to the Slack channel.
}
