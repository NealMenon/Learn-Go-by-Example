package main

import "fmt"

// Counter is struct
type Counter struct {
	value int
}

// Inc increments counter value
func (c *Counter) Inc() {
	c.value++
}

// Value increments counter value
func (c *Counter) Value() int {
	return c.value
}

func main() {
	fmt.Println("Build")
}
