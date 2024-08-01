package main

import "fmt"

type DfsStack struct {
	items []int
}

func (s *DfsStack) push(element int) {
	s.items = append(s.items, element)
}

func (s *DfsStack) pop() int {
	n := len(s.items) - 1
	result := s.items[n]
	s.items = s.items[:n]
	return result
}

func (s *DfsStack) peek() int {
	if len(s.items) > 0 {
		return s.items[len(s.items)-1]
	}
	return -1
}

func DfsOnComponents(edges [][]int, n int) {
	visited := make(map[int]bool, n)
	adjList := make(map[int][]int, n)
	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		adjList[u] = append(adjList[u], v)
		adjList[v] = append(adjList[v], u)
	}
	numComponents := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			numComponents++
			dfsHelper(i, visited, adjList)
		}
	}
	fmt.Println(numComponents)
}

func dfsHelper(u int, visited map[int]bool, adjList map[int][]int) {
	dfsStack := &DfsStack{}
	dfsStack.push(u)
	visited[u] = true
	for len(dfsStack.items) > 0 {
		item := dfsStack.pop()
		if !visited[item] {
			visited[item] = true
		}
		for _, nbrs := range adjList[item] {
			if !visited[nbrs] {
				dfsStack.push(nbrs)
			}
		}
	}
}

func dfsForCycleCheck(edges [][]int, n int) {
	//	visited := make(map[int]bool, n)
	adjList := make(map[int][]int, n)
	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		adjList[u] = append(adjList[u], v)
		adjList[v] = append(adjList[v], u)
	}
	//dfsCycleCheckHelper(0, visited, adjList)
	//	fmt.Println(bfsCycleDetection1(visited, 0, adjList, n))
	cycleDetectionUsingDegree(adjList)
}

func dfsCycleCheckHelper(u int, visited map[int]bool, adjList map[int][]int) bool {
	dfsStack := &DfsStack{}
	dfsStack.push(u)
	parent := make(map[int]int, 0)
	parent[u] = u
	for len(dfsStack.items) > 0 {
		item := dfsStack.pop()
		if !visited[item] {
			visited[item] = true
		}
		for _, nbrs := range adjList[item] {
			if !visited[nbrs] {
				dfsStack.push(nbrs)
				parent[nbrs] = item
			} else if visited[nbrs] && parent[nbrs] != item {
				fmt.Println("nbrs", nbrs)
				fmt.Println("item", item)
				fmt.Println("parent", parent[nbrs])
			}
		}
	}
	fmt.Println(parent)
	return false
}

func bfsCycleDetection1(visited map[int]bool, source int, al map[int][]int, n int) bool {
	parent := make([]int, n)
	parent[source] = -1
	visited[source] = true
	q := make([]int, 0)
	q = append(q, source)
	for len(q) > 0 {
		currentVertex := q[0]
		q = q[1:]
		for _, nbrs := range al[currentVertex] {
			if !visited[nbrs] {
				visited[nbrs] = true
				parent[nbrs] = currentVertex
				q = append(q, nbrs)
			} else if visited[nbrs] && nbrs != parent[currentVertex] {
				fmt.Println(parent)
				return true
			}
		}
	}
	fmt.Println(parent)
	return false
}

func cycleDetectionUsingDegree(al map[int][]int) {
	degree := make(map[int]int, 0)
	visited := make(map[int]bool, 0)
	q := make([]int, 0)
	// Generating the degree map form AL
	for k, v := range al {
		degree[k] = len(v)
	}
	// Adding those nodes whose degree is 1 to the queue
	for {
		for k, v := range degree {
			if v == 1 && !visited[k] {
				q = append(q, k)
			}
		}
		if len(q) == 0 {
			break
		}
		for len(q) != 0 {
			currentVertex := q[0]
			q = q[1:]
			visited[currentVertex] = true
			for _, nbrs := range al[currentVertex] {
				degree[nbrs] = degree[nbrs] - 1
			}
		}
	}
	for k, _ := range al {
		if visited[k] == false {
			fmt.Println(k)
		}
	}
}

func main() {
	//	DfsOnComponents([][]int{{0, 1}, {1, 2}, {3, 4}, {5, 6}, {7, 8}, {8, 9}}, 10)
	dfsForCycleCheck([][]int{{1, 2}, {1, 0}, {2, 0}, {0, 3}, {3, 4}}, 5)
}
