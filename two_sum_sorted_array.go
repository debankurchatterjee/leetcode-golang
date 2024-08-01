package main

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		sum := numbers[right] + numbers[left]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}

func main() {
	twoSum([]int{-1, 0}, -1)
}
