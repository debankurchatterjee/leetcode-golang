package main

import (
	"fmt"
	"sort"
)

func main() {
	ans := make([][]int, 0)
	nums := []int{3, 1, 3, 5, 1, 1}
	sort.Ints(nums)
	combinationSumIIOp(0, 8, nums, make([]int, 0), &ans)
	fmt.Println(ans)
}

func combinationSumII(index, target int, nums []int, ds []int, ans *[][]int) {
	if target == 0 {
		*ans = append(*ans, append([]int{}, ds...))
		return
	}
	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > target {
			break
		}
		ds = append(ds, nums[i])
		combinationSumII(i+1, target-nums[i], nums, ds, ans)
		ds = ds[:len(ds)-1]
	}
}

func combinationSumIIOp(index, target int, nums []int, ds []int, ans *[][]int) {
	if target == 0 {
		*ans = append(*ans, append([]int{}, ds...))
		return
	}
	// include the current element
	if index == len(nums) {
		return
	}
	if nums[index] <= target {
		ds = append(ds, nums[index])
		combinationSumIIOp(index+1, target-nums[index], nums, ds, ans)
		ds = ds[:len(ds)-1]
	}
	for index+1 < len(nums) && nums[index] == nums[index+1] {
		index++
	}
	combinationSumIIOp(index+1, target, nums, ds, ans)
}
