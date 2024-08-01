package main

import "fmt"

type Queue struct {
	elements []interface{}
	Size     int
}

func NewQueue(size int) *Queue {
	return &Queue{
		elements: make([]interface{}, 0, size),
		Size:     size,
	}
}

func (q *Queue) enqueue(item interface{}) {
	q.elements = append(q.elements, item)
}

func (q *Queue) dequeue() interface{} {
	item := q.elements[0]
	q.elements = q.elements[1:]
	return item
}

func (q *Queue) isEmpty() bool {
	return len(q.elements) == 0
}

type stack struct {
	items []int
}

func (s *stack) push(element int) {
	s.items = append(s.items, element)
}

func (s *stack) pop() int {
	n := len(s.items) - 1
	result := s.items[n]
	s.items = s.items[:n]
	return result
}

type Graph struct {
	V    int
	list map[int][]int
}

func NewGraph(v int) *Graph {
	return &Graph{
		V:    v,
		list: make(map[int][]int, v),
	}
}

func (g *Graph) addEdge(i, j int, undir bool) {
	currentData := g.list[i]
	currentData = append(currentData, j)
	g.list[i] = currentData
	if undir {
		currentDataJ := g.list[j]
		currentDataJ = append(currentDataJ, i)
		g.list[j] = currentDataJ
	}
}

func (g *Graph) bfs(source int) {
	visited := make([]bool, g.V)
	queue := NewQueue(g.V)
	queue.enqueue(source)
	//	fmt.Println("BFS", source)
	for !queue.isEmpty() {
		item := queue.dequeue()
		listOfNbrs := g.list[item.(int)]
		fmt.Print(item)
		for i := 0; i < len(listOfNbrs); i++ {
			if !visited[listOfNbrs[i]] {
				visited[listOfNbrs[i]] = true
				queue.enqueue(listOfNbrs[i])
				// fmt.Println("BFS", listOfNbrs[i])
			}
		}
	}
}

func (g *Graph) dfs(source int) {

}

func (g *Graph) connectedComponents() {

}

func (g *Graph) bfsShortestPath(source, destination int) int {
	visited := make([]bool, g.V)
	queue := NewQueue(g.V)
	queue.enqueue(source)
	parent := make([]int, g.V)
	parent[source] = source
	isDestinationReached := false
	for !queue.isEmpty() {
		currentNode := queue.dequeue().(int)
		if currentNode == destination {
			isDestinationReached = true
		}
		listOfNbrs := g.list[currentNode]
		for i := 0; i < len(listOfNbrs); i++ {
			if !visited[listOfNbrs[i]] {
				parent[listOfNbrs[i]] = currentNode
				visited[listOfNbrs[i]] = true
				queue.enqueue(listOfNbrs[i])
			}
		}
	}
	if isDestinationReached {
		temp := destination
		dist := 1
		for temp != source {
			fmt.Print(temp, "<--")
			temp = parent[temp]
			dist++
		}
		fmt.Print(source)
		fmt.Println("\n", dist)
		return dist
	}
	return -1
}

func (g *Graph) PrintAdjacencyList() {
	for k, v := range g.list {
		fmt.Printf("%d-->%v\n", k, v)
	}
}

func main() {
	graph := NewGraph(7)
	graph.addEdge(0, 1, true)
	graph.addEdge(0, 2, true)
	graph.addEdge(1, 3, true)
	graph.addEdge(1, 4, true)
	graph.addEdge(2, 5, true)
	graph.addEdge(2, 6, true)

	//	graph.dfs(0)
	//graph.PrintAdjacencyList()
	graph.bfs(0)
	// graph.bfsShortestPath(1, 6)
}
