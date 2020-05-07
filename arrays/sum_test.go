package main

import (
	"reflect"
	"testing"
)

// go test -cover tells us coverage of test code
func TestSum(t *testing.T) {

	checkSum := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v given", got, want)
		}
	}

	t.Run("collection of n numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := []int{Sum(numbers)}
		want := []int{15}
		checkSum(t, got, want)
	})

	t.Run("multiple lists passed", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		checkSum(t, got, want)
	})

	t.Run("testsumalltails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSum(t, got, want)
	})

	t.Run("Safety check on sumalltails", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSum(t, got, want)
	})

}
