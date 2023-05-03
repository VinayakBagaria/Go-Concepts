package graph

import "fmt"

func undirectedCycleCheck(adjList map[string][]string, visited map[string]bool, current string, parent string) bool {
	visited[current] = true

	for _, neighbour := range adjList[current] {
		if _, ok := visited[neighbour]; ok {
			if neighbour != parent {
				return true
			}
		} else if undirectedCycleCheck(adjList, visited, neighbour, current) {
			return true
		}
	}

	return false
}

func DoHasCycleUndirected() {
	edges := [][]string{
		{"0", "1"},
		{"0", "4"},
		{"1", "2"},
		{"1", "4"},
		{"2", "3"},
		{"4", "5"},
	}

	adjList := MakeAdjacencyList(edges)
	visited := map[string]bool{}

	fmt.Printf("Undirected Cycle: %t\n", undirectedCycleCheck(adjList, visited, "0", ""))
}
