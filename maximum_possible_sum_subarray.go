package main

import "fmt"

// Given an array a of length n, find the maximum possible sum from a subarray of length at most k.
func max_possible_sum(a []int, k int) {
	prefixSum := make([]int, len(a))
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
		prefixSum[i] = sum
	}
	fmt.Println(prefixSum)
}

func main() {
	max_possible_sum([]int{1, 2, 3, 4}, 5)
}
