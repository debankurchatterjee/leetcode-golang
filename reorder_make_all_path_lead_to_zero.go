package main

import "fmt"

func minReorder(n int, connections [][]int) int {
	adjacencyList := make(map[int][][]int, n)
	for _, edges := range connections {
		adjacencyList[edges[0]] = append(adjacencyList[edges[0]], []int{edges[1], 1})
		adjacencyList[edges[1]] = append(adjacencyList[edges[1]], []int{edges[0], 0})
	}
	ans := 0
	visited := make([]bool, n)
	q := make([]int, n)
	q[0] = 0
	visited[0] = true
	for len(q) != 0 {
		currentNode := q[0]
		q = q[1:]
		for _, nbrs := range adjacencyList[currentNode] {
			if !visited[nbrs[0]] {
				visited[nbrs[0]] = true
				ans = +nbrs[1]
				q = append(q, nbrs[0])
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(minReorder(5, [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}}))
}
