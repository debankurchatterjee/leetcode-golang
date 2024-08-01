package main

func validPath(n int, edges [][]int, source, destination int) bool {
	// Store graph data into a HashMap
	adjacencyList := make(map[int][]int, n)
	for _, edge := range edges {
		nodeA := edge[0]
		nodeB := edge[1]
		adjacencyList[nodeA] = append(adjacencyList[nodeA], nodeB)
		adjacencyList[nodeB] = append(adjacencyList[nodeB], nodeA)
	}
	visited := make([]bool, n, n)
	queue := NewQueue(n)
	queue.enqueue(source)
	visited[source] = true
	for !queue.isEmpty() {
		currentNode := queue.dequeue()
		if currentNode == destination {
			return true
		}
		for _, edge := range adjacencyList[currentNode.(int)] {
			if !visited[edge] {
				queue.enqueue(edge)
				visited[edge] = true
			}
		}
	}
	return false
}
