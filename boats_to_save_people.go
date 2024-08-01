package main

import "sort"

// sort the array and check if the heaviest person can share the boat with the lightest person
func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	left := 0
	right := len(people) - 1
	ans := 0
	for left <= right {
		ans++
		if people[left]+people[right] <= limit {
			left++
		}
		right--
	}
	return ans
}
