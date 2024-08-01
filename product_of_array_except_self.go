package main

import "fmt"

// Time O(N) and space O(N)
func productExceptSelf(nums []int) []int {
	prefixArr := make([]int, len(nums))
	suffixArr := make([]int, len(nums))
	ans := make([]int, 0)
	// find the prefix sum
	prefixArr[0] = 1
	suffixArr[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {
		prefixArr[i] = prefixArr[i-1] * nums[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		suffixArr[i] = suffixArr[i+1] * nums[i+1]
	}
	for i := 0; i < len(prefixArr); i++ {
		ans = append(ans, prefixArr[i]*suffixArr[i])
	}
	return ans
}

func productExceptSelfBetter(nums []int) []int {
	ans := make([]int, len(nums))
	// find the prefix sum
	ans[0] = 1
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	suffRight := 1
	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] = ans[i] * suffRight
		suffRight = suffRight * nums[i]
	}
	return ans
}

func main() {
	fmt.Println(productExceptSelfBetter([]int{4, 5, 1, 8, 2, 10, 6}))
}
