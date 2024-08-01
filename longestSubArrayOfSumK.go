package main

import "fmt"

// this is the optimal solution if the array contains +ve,-ve and zeros
func maximumSubarraySumPositiveElements(nums []int, k int) int64 {
	store := make(map[int]int)
	prefixSum := 0
	lenOfSub := 0
	totalSubArrays := 0
	for i := 0; i < len(nums); i++ {
		prefixSum += nums[i]
		if prefixSum == k {
			lenOfSub = max(lenOfSub, i+1)
			totalSubArrays++
		}
		if _, ok := store[prefixSum-k]; ok {
			lenOfSub = max(lenOfSub, i-store[prefixSum-k])
			totalSubArrays++
		}
		// if the sum is calculated previously we shouldn't update the map so that we can get the left most sum
		if _, ok := store[prefixSum]; !ok {
			store[prefixSum] = i
		}
	}
	fmt.Println(lenOfSub)
	return 0
}

// if the array contains +ve and zeros
func twoPointerSolution(nums []int, k int) {
	left := 0
	right := 0
	sum := nums[right-left+1]
	maxlen := 0
	for right < len(nums) {
		sum += nums[right]
		for left <= right && sum > k {
			sum = sum - nums[left]
			left++
		}
		if sum == k {
			maxlen = max(maxlen, right-left+1)
		}
		right++
	}
	fmt.Println(maxlen)
}

func main() {
	// maximumSubarraySumPositiveElements([]int{2, 0, 0, 3}, 3)
	twoPointerSolution([]int{2, 0, 0, 3}, 3)
}
