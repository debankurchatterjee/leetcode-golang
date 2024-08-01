package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	courseDependencyMatrix := make(map[int][]int, 0)
	courseDependencyDegree := make([]int, numCourses)
	result := make([]int, 0)
	for _, preqs := range prerequisites {
		u := preqs[1]
		v := preqs[0]
		courseDependencyMatrix[u] = append(courseDependencyMatrix[u], v)
		// In-degree of all the nodes
		courseDependencyDegree[v]++
	}
	q := make([]int, 0)
	// enqueue all elements with indegree zero
	for i, item := range courseDependencyDegree {
		if item == 0 {
			q = append(q, i)
		}
	}
	for len(q) != 0 {
		n := q[0]
		result = append(result, n)
		q = q[1:]
		for _, nbrs := range courseDependencyMatrix[n] {
			courseDependencyDegree[nbrs]--
			if courseDependencyDegree[nbrs] == 0 {
				q = append(q, nbrs)
			}
		}
	}
	if numCourses == len(result) {
		return result
	}
	return []int{}
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	courseDependencyMatrix := make(map[int][]int, 0)
	courseDependencyDegree := make([]int, numCourses)
	result := make([]int, 0)
	for _, preqs := range prerequisites {
		u := preqs[1]
		v := preqs[0]
		courseDependencyMatrix[u] = append(courseDependencyMatrix[u], v)
		// In-degree of all the nodes
		courseDependencyDegree[v]++
	}
	q := make([]int, 0)
	// enqueue all elements with indegree zero
	for i, item := range courseDependencyDegree {
		if item == 0 {
			q = append(q, i)
		}
	}
	for len(q) != 0 {
		n := q[0]
		result = append(result, n)
		q = q[1:]
		for _, nbrs := range courseDependencyMatrix[n] {
			courseDependencyDegree[nbrs]--
			if courseDependencyDegree[nbrs] == 0 {
				q = append(q, nbrs)
			}
		}
	}
	if numCourses == len(result) {
		return true
	}
	return false
}

func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}
	al := make(map[int][]int, 0)
	visited := make(map[int]bool, 0)
	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		al[u] = append(al[u], v)
		al[v] = append(al[v], u)
	}
	q := make([]int, 0)
	q = append(q, 0)
	visited[0] = true
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		for _, nbr := range al[node] {
			if !visited[nbr] {
				q = append(q, nbr)
				visited[nbr] = true
			}
		}
	}
	if len(visited) == n {
		return true
	}
	return false
}

func countComponents(n int, edges [][]int) int {
	seen := make(map[int]bool, 0)
	al := make(map[int][]int, 0)
	for _, edge := range edges {
		al[edge[0]] = append(al[edge[0]], edge[1])
		al[edge[1]] = append(al[edge[1]], edge[0])
	}
	totalCycleCount := 0
	for i := 0; i < n; i++ {
		if !seen[i] {
			totalCycleCount++
			bfs(i, al, seen)
		}
	}
	return totalCycleCount
}

func bfs(node int, al map[int][]int, seen map[int]bool) {
	seen[node] = true
	q := make([]int, 0)
	q = append(q, node)
	for len(q) != 0 {
		currentNode := q[0]
		q = q[1:]
		for _, nbr := range al[currentNode] {
			if seen[nbr] {
				continue
			}
			seen[nbr] = true
			q = append(q, nbr)
		}
	}
}

func main() {
	//	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}, {0, 3}}))

	fmt.Println(countComponents(3, [][]int{{0, 1}, {1, 2}}))
}
