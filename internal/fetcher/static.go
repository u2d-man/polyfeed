package fetcher

type StaticRSSFetcher struct {
	URLs []string
}

func (s StaticRSSFetcher) GetRssURLs() ([]string, error) {
	return s.URLs, nil
}
