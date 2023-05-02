package graph

import "fmt"

type visitedType map[string]bool

func hasPath(graph map[string][]string, source, dest string, visited visitedType) bool {
	if source == dest {
		return true
	}

	if _, ok := visited[source]; ok {
		return false
	}

	visited[source] = true

	for _, neighbour := range graph[source] {
		if hasPath(graph, neighbour, dest, visited) {
			return true
		}
	}

	return false
}

func DoHasPathUndirected() {
	graph := MakeAdjacencyList()
	fmt.Println("Has Path in Undirected Graph DFS")
	fmt.Printf("b/w i & k: %t\n", hasPath(graph, "i", "k", visitedType{}))

}
