package main

import "fmt"

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	results := make([]int, 0)
	incomingEdges := make([]bool, n)
	for i := 0; i < n; i++ {
		incomingEdges[i] = false
	}
	for _, edge := range edges {
		incomingEdges[edge[1]] = true
	}
	for j, incomingEdge := range incomingEdges {
		if !incomingEdge {
			results = append(results, j)
		}
	}
	return results
}

func main() {
	fmt.Println(findSmallestSetOfVertices(5, [][]int{{0, 1}, {2, 1}, {3, 1}, {1, 4}, {2, 4}}))
}
