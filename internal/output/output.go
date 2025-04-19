/*
Package output provides mechanisms for outputting processed articles.

The package includes implementations for different output destinations,
such as local file output and Slack notification integration.
*/
package output

import "github.com/u2d-man/polyfeed/internal/core"

// OutputService defines the interface for outputting articles
type OutputService interface {
	// Output sends the provided articles to the output destination
	// Returns an error if the operation fails
	Output(articles []core.Article) error
}
