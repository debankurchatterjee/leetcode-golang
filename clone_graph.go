package main

type Node struct {
	Val       int
	Neighbors []*Node
}

// Iterative BFS Solution
func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}
	visited := make(map[*Node]*Node)
	queue := make([]*Node, 0)
	queue = append(queue, node)
	visited[node] = &Node{
		Val:       node.Val,
		Neighbors: node.Neighbors,
	}
	for len(queue) != 0 {
		item := queue[0]
		queue = queue[1:]
		for _, nbr := range item.Neighbors {
			if visited[nbr] == nil {
				visited[nbr] = &Node{
					Val:       nbr.Val,
					Neighbors: make([]*Node, 0),
				}
				queue = append(queue, nbr)
			}
			visited[nbr].Neighbors = append(visited[nbr].Neighbors, nbr.Neighbors...)
		}
	}
	return visited[node]
}
