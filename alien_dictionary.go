package main

import (
	"strings"
)

func alienOrder2(words []string) string {
	adj := make(map[byte][]byte)
	indegree := make(map[byte]byte)
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			indegree[words[i][j]] = 0
		}
	}
	for i := 0; i < len(words)-1; i++ {
		s1 := words[i]
		s2 := words[i+1]
		if len(s1) > len(s2) && strings.HasPrefix(s1, s2) {
			return ""
		}
		size := min(len(s1), len(s2))
		for ptr := 0; ptr < size; ptr++ {
			if s1[ptr] != s2[ptr] {
				adj[s1[ptr]] = append(adj[s1[ptr]], s2[ptr])
				indegree[s2[ptr]] = indegree[s2[ptr]] + 1
				break
			}
		}
	}
	q := make([]byte, 0)
	var topoResults []byte
	for k, v := range indegree {
		if v == 0 {
			q = append(q, k)
		}
	}
	for len(q) != 0 {
		item := q[0]
		q = q[1:]
		topoResults = append(topoResults, item)
		for _, nbrs := range adj[item] {
			indegree[nbrs] = indegree[nbrs] - 1
			if indegree[nbrs] == 0 {
				q = append(q, nbrs)
			}
		}
	}
	if len(topoResults) < len(indegree) {
		return ""
	}
	return string(topoResults)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	alienOrder2([]string{"z", "x", "a", "zb", "zx"})
}
