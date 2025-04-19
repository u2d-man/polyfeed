package fetcher

import (
	"bufio"
	"os"
	"strings"
)

// FileRSSFetcher implements the RSSFetcher interface by reading RSS feed URLs from a file.
// Each line in the file is expected to contain a single RSS feed URL.
// Lines starting with # are treated as comments and skipped.
type FileRSSFetcher struct {
	// Path is the filesystem path to the RSS URLs file
	Path string
}

// GetRssURLs reads RSS feed URLs from the file specified in the Path field.
// It trims whitespace from each line and returns all non-empty, non-comment lines as URLs.
//
// Returns:
//   - A slice of strings containing the RSS feed URLs
//   - An error if the file cannot be read or processed
func (f FileRSSFetcher) GetRssURLs() ([]string, error) {
	file, err := os.Open(f.Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// skip comment line.
		if line != "" && !strings.HasPrefix(line, "#") {
			urls = append(urls, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}
