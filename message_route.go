package main

func messageRoute(n int, edges [][]int) int {
	graph := NewGraph(n)
	for i := 0; i < len(edges); i++ {
		for j := 0; j < len(edges[i]); j++ {
			graph.addEdge(i, j, true)
		}
	}
	return graph.bfsShortestPath(1, n)
}
