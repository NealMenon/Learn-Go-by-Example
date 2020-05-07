package main

// Sum returns sum of array elems
// func Sum(numbers [5]int) int {
// 	var sum int
// 	// for i := 0; i < 5; i++ {
// 	// 	sum += numbers[i]
// 	// }
// 	sum = 0
// 	for _, val := range numbers {
// 		sum += val
// 	}

// 	return sum
// }

// Sum for slice
func Sum(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

/* SumAll takes a varying number of slices,
 *  returning a new slice containing the totals for each
 *  slice passed in.
 */

// SumAll func
func SumAll(numsToSum ...[]int) []int {
	// vals := make([]int, len(numsToSum))
	vals := []int{}
	for _, nums := range numsToSum {
		// vals[i] = Sum(nums)
		vals = append(vals, Sum(nums))

	}
	return vals
}

// SumAllTails sums excluding head
func SumAllTails(numsToSum ...[]int) []int {
	var vals []int
	for _, nums := range numsToSum {
		if len(nums) == 0 {
			vals = append(vals, 0)
		} else {
			tail := nums[1:]
			vals = append(vals, Sum(tail))
		}
	}
	return vals
}

func main() {
	// fmt.Println(Sum([]int{1, 2, 3, 4, 5}))
	// fmt.Println("Hi")
}
