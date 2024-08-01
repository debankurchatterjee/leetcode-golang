package main

import (
	"fmt"
	"slices"
)

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	ds := make([]int, 0)
	f_pre(ds, &result, nums, 0)
	return result
}

func f_pre(ds []int, result *[][]int, nums []int, idx int) {
	if len(ds) == len(nums) {
		*result = append(*result, append([]int{}, ds...))
		return
	}
	for _, element := range nums {
		if !slices.Contains(ds, element) {
			ds = append(ds, element)
			f_pre(ds, result, nums, idx+1)
			ds = ds[:len(ds)-1]
		}
	}
}

func combi(nums []int) [][]int {
	result := make([][]int, 0)
	ds := make([]int, 0)
	f_combi(ds, &result, nums, 0)
	return result
}

func f_combi(ds []int, result *[][]int, nums []int, idx int) {
	if idx == len(nums) {
		*result = append(*result, append([]int{}, ds...))
		return
	}
	ds = append(ds, nums[idx])
	f_combi(ds, result, nums, idx+1)
	ds = ds[:len(ds)-1]
	f_combi(ds, result, nums, idx+1)
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(combi([]int{1, 2, 3}))
}
