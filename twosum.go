package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numcDict := make(map[int]int, 0)
	for i, num := range nums {
		diff := target - num
		val, ok := numcDict[diff]
		if ok {
			return []int{i, val}
		}
		numcDict[num] = i
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))

	test := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(test)-1; i++ {
		test[i] = test[i+1]
	}
	fmt.Println(test)
}
