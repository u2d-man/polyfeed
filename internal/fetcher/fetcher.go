package fetcher

type RSSFetcher interface {
	GetRssURLs() ([]string, error)
}
