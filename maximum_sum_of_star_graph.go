package main

import "sort"

func maxStarSum(vals []int, edges [][]int, k int) int {
	//	sum := 0
	al := make(map[int][]int, 0)
	for _, edge := range edges {
		al[edge[0]] = append(al[edge[0]], edge[1])
		al[edge[1]] = append(al[edge[1]], edge[0])
		sort.Ints(al[edge[0]])
		sort.Ints(al[edge[1]])
	}
	return -1
}

func main() {
	maxStarSum([]int{1, 2, 3, 4, 10, -10, -20}, [][]int{{}}, 2)
}
