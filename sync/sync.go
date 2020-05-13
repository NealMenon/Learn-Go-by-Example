package main

import (
	"fmt"
	"sync"
)

// Counter is struct
type Counter struct {
	mu    sync.Mutex
	value int
}

// NewCounter returns pointer to counter
func NewCounter() *Counter {
	return &Counter{}
}

// Inc increments counter value
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value increments counter value
func (c *Counter) Value() int {
	return c.value
}

func main() {
	fmt.Println("Build")
}
