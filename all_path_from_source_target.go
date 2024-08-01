package main

import "fmt"

var results = make([][]int, 0)

func allPathsSourceTarget(graph [][]int) [][]int {
	res := [][]int{}
	var dfs func([]int, int)
	dfs = func(path []int, node int) {
		path = append(path, node)
		if node == len(graph)-1 {
			completedPath := make([]int, 0)
			completedPath = append(completedPath, path...)
			res = append(res, path)
			return
		}
		for _, neighbor := range graph[node] {
			dfs(path, neighbor)
		}
	}
	dfs([]int{}, 0)
	return res
}

func main() {
	fmt.Println(allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
}
