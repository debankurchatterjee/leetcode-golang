package main

import "fmt"

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	queue := []int{}
	graph := make([][]int, n)
	degree := make([]int, n)

	for _, e := range edges {
		na, nb := e[0], e[1]
		graph[na] = append(graph[na], nb)
		graph[nb] = append(graph[nb], na)
		degree[na]++
		degree[nb]++
	}

	for i := range degree {
		if degree[i] == 1 {
			queue = append(queue, i)
		}
	}

	// topological sort
	for n > 2 {
		size := len(queue)
		n -= size
		for size > 0 {
			node := queue[0]
			queue = queue[1:]
			size--
			for _, v := range graph[node] {
				degree[v]--
				if degree[v] == 1 {
					queue = append(queue, v)
				}
			}

		}
	}

	return queue
}

func main() {
	fmt.Println(findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}}))
}
