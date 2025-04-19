/*
Package fetcher provides interfaces and implementations for retrieving RSS feed URLs.

The package defines the RSSFetcher interface and provides different implementation
strategies such as fetching URLs from files or environment variables.
*/
package fetcher

// RSSFetcher defines the interface for retrieving RSS feed URLs
type RSSFetcher interface {
	// GetRssURLs retrieves a list of RSS feed URLs from a source
	// Returns an error if the URLs cannot be retrieved
	GetRssURLs() ([]string, error)
}
