package main

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	visited[0] = true
	q := []int{0}
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		for _, room := range rooms[cur] {
			if !visited[room] {
				visited[room] = true
				q = append(q, room)
			}
		}
	}
	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}

func main() {
	canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}})
}
