package concurrency

import "fmt"

// WebsiteChecker func
type WebsiteChecker func(string) bool

// result is struct for each url
type result struct {
	string
	bool
}

// CheckWebsites checks validity of website
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func(url string) {
			results[url] = wc(url)
		}(url)

	}

	return results
}

func main() {
	fmt.Println("Test main")
}
