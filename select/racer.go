package main

import (
	"fmt"
	"net/http"
	"time"
)

// Racer asdf
func Racer(a, b string) string {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}
	return b
}

func measureResponseTime(u string) time.Duration {
	start := time.Now()
	http.Get(u)
	return time.Since(start)
}
func main() {
	fmt.Println("printer")
}
