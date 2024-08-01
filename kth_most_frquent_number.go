package main

func topKFrequent(nums []int, k int) []int {
	store := make(map[int]int)
	for _, num := range nums {
		store[num]++
	}
	bucket := make([][]int, len(nums)+1)
	// insert all the elements in the bucket
	for k, v := range store {
		bucket[v] = append(bucket[v], k)
	}
	result := make([]int, k)
	counter := 0
	for i := len(bucket) - 1; i >= 0 && counter < k; i-- {
		if bucket[i] != nil {
			for _, item := range bucket[i] {
				result[counter] = item
				counter++
			}
		}
	}
	return result
}

func main() {
	topKFrequent([]int{4, 1, -1, 2, -1, 2, 3}, 2)
}
