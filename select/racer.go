package main

import (
	"fmt"
	"net/http"
	"time"
)

// Racer asdf
func Racer(a, b string) string {
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		return a
	}

	return b
}

func main() {
	fmt.Println("printer")
}
