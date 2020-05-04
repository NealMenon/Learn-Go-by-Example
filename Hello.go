package main

import (
	"fmt"
)

func helloer() string {
	return "Hello you\n"
}

func main() {
	fmt.Printf(helloer())
	var i int
	i = 424
	fmt.Println(i)
}
