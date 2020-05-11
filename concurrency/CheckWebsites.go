package concurrency

// WebsiteChecker func
type WebsiteChecker func(string) bool

// CheckWebsites checks validity of website
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}
