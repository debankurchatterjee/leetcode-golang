package main

import (
	"fmt"
)

// combine is a recursive function to generate and print all combinations of length k from the array.
func combine(arr []int, n int, k int, index int, current []int) {
	// When the combination is complete, print it.
	if len(current) == k {
		fmt.Println(current)
		return
	}

	// Return if no more elements are there to put in current.
	if index == n {
		return
	}

	// Current is included, put next at next location.
	current = append(current, arr[index])

	// Current is included, replace it with next (or) Current is excluded (remove last added item) and replace it with next.
	combine(arr, n, k, index+1, current)
	current = current[:len(current)-1]
	combine(arr, n, k, index+1, current)
}

func main() {
	arr := []int{1, 2, 3, 4}
	n := len(arr)
	//	k := len(arr) // Length of the combinations

	combine(arr, n, 3, 0, []int{})
}
