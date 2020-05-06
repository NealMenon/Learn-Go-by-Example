package main

import "fmt"

// Sum returns sum of array elems
func Sum(numbers [5]int) int {
	var sum int
	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }
	sum = 0
	for _, val := range numbers {
		sum += val
	}

	return sum
}

func main() {
	fmt.Println(Sum([5]int{1, 2, 3, 4, 5}))
}
